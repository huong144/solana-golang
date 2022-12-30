package schema

import "gorm.io/gorm"

type CheckBlockFlag struct {
	gorm.Model
	BlockNumber int64 `gorm:"column:block_number;index:idx_check_block_flag"`
}

func (CheckBlockFlag) TableName() string {
	return "check-block-flag"
}
