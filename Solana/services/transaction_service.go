package services

import (
	"encoding/json"
	"flag"
	"gorm.io/gorm"
	"log"
	"solana-crawl-service/Solana/schema"
	"sync"
	"time"
)

type Params []string

type Result struct {
	Id                int            `json:"id"`
	Transactions      []Transactions `json:"transactions"`
	ParentSlot        int            `json:"parentSlot"`
	BlockTime         int64          `json:"blockTime"`
	BlockSlot         int64          `json:"blockSlot"`
	BlockHash         string         `json:"blockHash"`
	BlockHeight       int32          `json:"blockHeight"`
	PreviousBlockhash string         `json:"previousBlockhash"`
}

type Transactions struct {
	Transaction Transaction `json:"transaction"`
	Meta        Meta        `json:"meta"`
}

type Meta struct {
	Err               interface{}        `json:"err"`
	Fee               uint64             `json:"fee"`
	InnerInstructions []InnerInstruction `json:"innerInstructions"`
	PostBalances      []uint64           `json:"postBalances"`
	PostTokenBalances []interface{}      `json:"postTokenBalances"`
	PreBalances       []uint64           `json:"preBalances"`
	PreTokenBalances  []interface{}      `json:"preTokenBalances"`
	Status            interface{}        `json:"status"`
	LogMessages       []string           `json:"logMessages"`
}

type Transaction struct {
	Message   Message  `json:"message"`
	Signature []string `json:"signatures"`
}

type Message struct {
	AccountKeys         []AccountKey  `json:"accountKeys"`
	AddressTableLookups interface{}   `json:"addressTableLookups"`
	Instructions        []Instruction `json:"instructions"`
	RecentBlockhash     string        `json:"recentBlockhash"`
}

type AccountKey struct {
	PubKey    string `json:"pubKey"`
	Signer    bool   `json:"signer"`
	Source    string `json:"source"`
	Writeable bool   `json:"writeable"`
}

type InnerInstruction struct {
	Index        int           `json:"index"`
	Instructions []Instruction `json:"instructions"`
}

type Instruction struct {
	Accounts  []string `json:"accounts"`
	Data      string   `json:"data"`
	ProgramId string   `json:"programId"`
	Program   string   `json:"program"`
	Parsed    Parsed   `json:"parsed"`
}

type Parsed struct {
	Info interface{} `json:"info"`
	Type string      `json:"type"`
}

type BlockData struct {
	Blockhash string `json:"blockhash"`
	BlockTime int64  `json:"blockTime"`
	BlockSlot int64  `json:"blockSlot"`
}

func GetDecodeListData(db *gorm.DB, blockSlot uint64, data []uint8) bool {
	rsp := Result{}
	json.Unmarshal(data, &rsp)
	blockDT := BlockData{
		Blockhash: rsp.BlockHash,
		BlockTime: rsp.BlockTime,
		BlockSlot: int64(blockSlot),
	}
	//db, err := Database.ConnectDB()
	//sqlDB, err := db.DB()
	//defer sqlDB.Close()

	//wg.Add(len(rsp.Transactions))
	//wg.Add(100)
	//txArrs := []schema.Transaction{}

	// Config go routines
	maxNbConcurrentGoroutines := flag.Int("maxNbConcurrentGoroutines", 100, "the number of goroutines that are allowed to run concurrently")
	nbJobs := flag.Int("nbJobs", len(rsp.Transactions), "the number of jobs that we need to do")
	flag.Parse()

	concurrentGoroutines := make(chan struct{}, *maxNbConcurrentGoroutines)
	var wg sync.WaitGroup
	for i := 0; i < *nbJobs; i++ {
		wg.Add(1)
		go func() {
			concurrentGoroutines <- struct{}{}
			log.Println("doing", i)
			insertTxDecoded(rsp.Transactions[i], blockDT, db)
			log.Println("finished", i)
			<-concurrentGoroutines
		}()

	}
	wg.Wait()

	return true
}

func insertTxDecoded(data Transactions, block BlockData, db *gorm.DB) {
	errTx, _ := json.Marshal(data.Meta.Err)
	sig, _ := json.Marshal(data.Transaction.Signature)
	logs, _ := json.Marshal(data.Meta.LogMessages)
	// Balance changed
	postBalance := data.Meta.PostBalances
	preBalance := data.Meta.PreBalances
	var blcArr []schema.BalanceChange
	var sigTxRelArr []schema.SignatureAddressRel
	for index, item := range data.Transaction.Message.AccountKeys {
		blc := schema.BalanceChange{
			AccountAddress: item.PubKey,
			PreBalance:     int64(preBalance[index]),
			PostBalance:    int64(postBalance[index]),
			IsSigner:       item.Signer,
			IsWriteable:    item.Writeable,
		}
		sigTxRel := schema.SignatureAddressRel{
			Signature:      string(sig),
			Address:        item.PubKey,
			BlockNumber:    block.BlockSlot,
			BlockTimestamp: block.BlockTime,
		}
		blcArr = append(blcArr, blc)
		sigTxRelArr = append(sigTxRelArr, sigTxRel)
	}

	//Instruction + Inner Instruction
	var arrIns = make([]schema.Instruction, len(data.Transaction.Message.Instructions))
	for i, item := range data.Transaction.Message.Instructions {
		var innerInsArr []schema.InnerInstruction
		if len(data.Meta.InnerInstructions) > 0 {
			for _, e := range data.Meta.InnerInstructions {
				//innerInsArr := make([]schema.InnerInstruction, len(e.Instructions))
				var count = 0
				if e.Index == i {
					for _, ins := range e.Instructions {
						parse, _ := json.Marshal(ins.Parsed.Info)
						account, _ := json.Marshal(ins.Accounts)
						innerItem := schema.InnerInstruction{
							Parse:      string(parse),
							Type:       ins.Parsed.Type,
							Program:    ins.Program,
							ProgramId:  ins.ProgramId,
							Account:    string(account),
							DataBase58: ins.Data,
						}
						innerInsArr = append(innerInsArr, innerItem)
					}
					count++
				}
			}
		} else {
			innerInsArr = nil
		}
		parse, _ := json.Marshal(item.Parsed.Info)
		account, _ := json.Marshal(item.Accounts)
		ins := schema.Instruction{
			Index:            i,
			Parse:            string(parse),
			Type:             item.Parsed.Type,
			Program:          item.Program,
			ProgramID:        item.ProgramId,
			Accounts:         string(account),
			Data58:           item.Data,
			InnerInstruction: innerInsArr,
		}
		arrIns[i] = ins
	}
	txModel := schema.Transaction{
		BlockSlot:      int(block.BlockSlot),
		BlockTimestamp: uint64(block.BlockTime),
		Signature:      string(sig),
		BalanceChange:  blcArr,
		Instruction:    arrIns,
		Logs:           string(logs),
		Errs:           string(errTx),
	}

	//return txModel

	// add to queue

	func() {
		insertedSigTxRel(db, sigTxRelArr)
		insertedTx(db, txModel)
		time.Sleep(1 * time.Millisecond)
	}()

}

func insertedTx(db *gorm.DB, data schema.Transaction) error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func insertedSigTxRel(db *gorm.DB, data []schema.SignatureAddressRel) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
