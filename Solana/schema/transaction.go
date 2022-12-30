package schema

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	//ID             uint            `gorm:"column:id;primaryKey;autoIncrement:true"`
	BlockSlot      int             `gorm:"column:block_slot"`
	BlockTimestamp uint64          `gorm:"column:block_timestamp"`
	Signature      string          `gorm:"column:signature;index:idx_signature_tx,unique"`
	BalanceChange  []BalanceChange `gorm:"foreignKey:TransactionId"`
	Instruction    []Instruction   `gorm:"foreignKey:TransactionId"`
	Logs           string          `gorm:"column:logs;type:text"`
	Errs           string          `gorm:"column:err"`
}

func (Transaction) TableName() string {
	return "transaction"
}
