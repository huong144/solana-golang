package models

type Block struct {
	Slot uint64 `json:"slot" bson:"slot"`
	//Blockhash         string `json:"block_hash" bson:"block_hash"`
	//BlockTime         int64  `json:"block_time" bson:"block_time"`
	//BlockHeight       int64  `json:"block_height" bson:"block_height"`
	//PreviousBlockhash string `json:"previous_block_hash" bson:"previous_block_hash""`
	//ParentSlot        uint64 `json:"parent_slot" bson:"parent_slot"`
	//Transaction       []Transaction `json:"transaction"`
	DataRaw string `json:"data_raw" bson:"data_raw"`
}

type SyncBlock struct {
	BlockNumber int64  `json:"block_number" bson:"block_number"`
	Type        string `json:"type" json:"type"`
	Status      string `json:"status" bson:"status"`
}
