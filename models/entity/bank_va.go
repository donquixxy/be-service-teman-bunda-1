package entity

type BankVa struct {
	Id                 string  `gorm:"primaryKey;column:id;"`
	BankName           string  `gorm:"column:bank_name;"`
	BankLogo           string  `gorm:"column:bank_logo;"`
	BankCode           string  `gorm:"column:bank_code;"`
	AdminFee           float64 `gorm:"column:admin_fee;"`
	AdminFeePercentage float64 `gorm:"column:admin_fee_percentage;"`
}

func (BankVa) TableName() string {
	return "banks_virtual_account"
}
