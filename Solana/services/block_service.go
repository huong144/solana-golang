package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"solana-crawl-service/Database"
	"solana-crawl-service/Solana/models"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

func GetSaveBlock(block models.Block) *mongo.InsertOneResult {
	blockRepository, ctx, client := Database.ConnectDatabase("Block")
	c, err := blockRepository.InsertOne(ctx, block)
	if err != nil {
		fmt.Printf("okela", err, err)
	}
	defer client.Disconnect(ctx)
	return c
}

func GetSyncBlock() {
	var lastSync models.SyncBlock
	var lastSlotSync int32
	var startBlock uint64
	var endBlock uint64
	latestSlot, err := GetSlot()
	if err != nil {
		log.Fatalln("error", err)
	}
	filter := bson.D{{"type", "solana"}, {"status", "new-fetch"}}
	opts := options.FindOne()
	syncedBlockRepository, ctx, client := Database.ConnectDatabase("SyncBlock")
	syncedBlockRepository.FindOne(context.TODO(), filter, opts).Decode(&lastSync)
	if lastSync.BlockNumber != 0 {
		// lastSlotSync = int32(lastSync.BlockNumber)
		// updated
		syncedBlockRepository.UpdateOne(context.TODO(),
			bson.D{{"type", "solana"}, {"status", "new-fetch"}},
			bson.D{{"$set", bson.D{{"block_number", latestSlot}}}},
			options.Update().SetUpsert(true))
		startBlock = uint64(lastSync.BlockNumber + 1)
		endBlock = latestSlot
	} else {
		lastSlotSync = int32(latestSlot)
		syncedBlockRepository.InsertOne(ctx, models.SyncBlock{
			BlockNumber: int64(lastSlotSync),
			Type:        "solana",
			Status:      "new-fetch",
		})
		startBlock = latestSlot
		endBlock = latestSlot
	}
	// Update into SyncBlock
	defer client.Disconnect(ctx)
	getBlock(startBlock, endBlock)

}

func GetSyncOldBlock() {
	var lastOldSync models.SyncBlock
	var lastSlotSync int32
	var startBlock uint64
	var endBlock uint64
	filter := bson.D{{"type", "solana"}, {"status", "old-fetch"}}
	filter2 := bson.D{{"type", "solana"}, {"status", "new-fetch"}}
	opts := options.FindOne()
	syncedBlockRepository, ctx, client := Database.ConnectDatabase("SyncBlock")
	syncedBlockRepository.FindOne(context.TODO(), filter, opts).Decode(&lastOldSync)
	if lastOldSync.BlockNumber != 0 {
		lastSlotSync = int32(lastOldSync.BlockNumber)
		startBlock = uint64(lastOldSync.BlockNumber) - uint64(100)
		endBlock = uint64(lastOldSync.BlockNumber) - 1
		// updated
		syncedBlockRepository.UpdateOne(context.TODO(),
			bson.D{{"type", "solana"}, {"status", "old-fetch"}},
			bson.D{{"$set", bson.D{{"block_number", startBlock}}}},
			options.Update().SetUpsert(true))

	} else {
		syncedBlockRepository.FindOne(context.TODO(), filter2, opts).Decode(&lastOldSync)
		if lastOldSync.BlockNumber != 0 {
			startBlock = uint64(lastOldSync.BlockNumber) - 100
			endBlock = uint64(lastOldSync.BlockNumber) - 1
			lastSlotSync = int32(startBlock)

		} else {
			latestSlot, err := GetSlot()
			if err != nil {
				log.Fatalln("error", err)
			}
			startBlock = uint64(latestSlot) - 100
			endBlock = uint64(lastSlotSync)
			lastSlotSync = int32(startBlock)

		}

	}
	// Update into SyncBlock
	defer client.Disconnect(ctx)
	getOldBlock(startBlock, endBlock)
	syncedBlockRepository.InsertOne(ctx, models.SyncBlock{
		BlockNumber: int64(lastSlotSync),
		Type:        "solana",
		Status:      "old-fetch",
	})

}

func getOldBlock(startBlock uint64, endBlock uint64) {
	fmt.Println("Fetching old data from ", startBlock, "to", endBlock)
	for i := startBlock; i <= endBlock; i++ {
		handleProcessBlockData(i, nil)
	}
}

func TestGetBlock(block uint64) {
	handleProcessBlockData(block, nil)
}

func getBlock(startBlock uint64, endBlock uint64) {
	fmt.Println("Start fetching data from ", startBlock, " to ", endBlock)
	// Connect DB
	db, err := Database.ConnectDB()
	if err != nil {
		log.Println("Connect DB failed")
	}
	//var wg sync.WaitGroup
	//wg.Add(1)
	for i := startBlock; i <= endBlock; i++ {
		fmt.Println(i)
		handleProcessBlockData(i, db)
		//defer wg.Done()
		//time.Sleep(2 * time.Second)
	}
	//wg.Wait()
}

func handleProcessBlockData(block uint64, db *gorm.DB) {
	//var wg sync.WaitGroup
	c := client.NewClient(rpc_endpoint)
	//res, error := c.GetBlock(context.TODO(), block)
	res, _ := c.RpcClient.GetBlockWithConfig(
		context.TODO(),
		block,
		rpc.GetBlockConfig{
			Encoding:           rpc.GetBlockConfigEncodingJsonParsed,
			TransactionDetails: rpc.GetBlockConfigTransactionDetailsFull,
			Commitment:         "finalized",
		},
	)
	dataEncode, _ := json.Marshal(res.Result)
	//stringJson := string(dataEncode)
	//wg.Add(2)
	func() {
		//GetDecodeListData(db, block, dataEncode)
		getProcessData(dataEncode, block)
		//wg.Done()
	}()
	//wg.Wait()
	//log.Println(dataEncode, stringJson)

	//go getProcessData(res, block)
	//HandleTransactionData(res, block)

}

func getProcessData(data []uint8, block uint64) {
	GetSaveBlock(models.Block{
		Slot: block,
		//ParentSlot:        data.ParentSLot,
		DataRaw: string(data),
	})
}

func GetSlot() (uint64, error) {
	c := client.NewClient(rpc_endpoint)

	res, error := c.GetSlot(context.TODO())
	if error != nil {
		log.Println("failed to get slot Solana, err: %v", error)
	}
	return res, error
}
