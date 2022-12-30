package schema

import "gorm.io/gorm"

type Instruction struct {
	gorm.Model
	//ID               uint               `gorm:"column:id;primaryKey;autoIncrement:true"`
	Index            int                `gorm:"column:index;index:idx_instruction_tx,unique"`
	Parse            string             `gorm:"column:parse"`
	Type             string             `gorm:"column:type"`
	Program          string             `gorm:"column:program"`
	ProgramID        string             `gorm:"column:program_id"`
	Accounts         string             `gorm:"column:accounts"`
	Data58           string             `gorm:"column:data_58"`
	InnerInstruction []InnerInstruction `gorm:"foreignKey:InstructionId"`
	TransactionId    int                `gorm:"column:transaction_id;index:idx_instruction_tx,unique"`
}

func (Instruction) TableName() string {
	return "instruction"
}
