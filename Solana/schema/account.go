package schema

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	//ID        uint   `gorm:"column:id;primaryKey;autoIncrement:true"`
	Address   string `gorm:"column:address;index:idx_address,unique"`
	Type      string `gorm:"column:type"`
	IsChecked bool   `gorm:"column:is_checked;default:false"`
}

func (Account) TableName() string {
	return "account"
}
