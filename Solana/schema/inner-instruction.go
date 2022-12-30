package schema

import (
	"gorm.io/gorm"
)

type InnerInstruction struct {
	gorm.Model
	//ID            uint   `gorm:"column:id;primaryKey;autoIncrement:true"`
	InstructionId uint   `gorm:"column:instruction_id"`
	Parse         string `gorm:"column:parse;type:text"`
	Type          string `gorm:"column:type"`
	Program       string `gorm:"column:program"`
	ProgramId     string `gorm:"column:program_id"`
	Account       string `gorm:"column:accounts;type:text"`
	DataBase58    string `gorm:"column:data_base58"`
}

func (InnerInstruction) TableName() string {
	return "inner_instruction"
}
