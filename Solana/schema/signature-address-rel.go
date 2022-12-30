package schema

import "gorm.io/gorm"

type SignatureAddressRel struct {
	gorm.Model
	Signature      string `gorm:"column:signature;index:idx_sig_tx,unique"`
	Address        string `gorm:"column:address;index:idx_sig_tx,unique"`
	BlockNumber    int64  `gorm:"column:block_number;index:idx_sig_tx,unique"`
	BlockTimestamp int64  `gorm:"column:block_timestamp;index:idx_sig_tx,unique"`
}

func (SignatureAddressRel) TableName() string {
	return "signature-address-rel"
}
