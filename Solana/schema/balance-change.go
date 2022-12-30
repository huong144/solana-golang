package schema

import "gorm.io/gorm"

type BalanceChange struct {
	gorm.Model
	//ID             uint   `gorm:"column:id;primaryKey;autoIncrement:true"`
	AccountAddress string `gorm:"column:account_address";`
	TransactionId  uint   `gorm:"column:transaction_id"`
	PreBalance     int64  `gorm:"column:pre_balance"`
	PostBalance    int64  `gorm:"column:post_balance"`
	IsSigner       bool   `gorm:"column:is_signer"`
	IsWriteable    bool   `gorm:"column:is_writeable"`
}

func (BalanceChange) TableName() string {
	return "balance_change"
}
