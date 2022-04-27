package entity

type BankVa struct {
	Id       string `gorm:"primaryKey;column:id;"`
	BankName string `gorm:"column:bank_name;"`
	BankLogo string `gorm:"column:bank_logo;"`
	BankCode string `gorm:"column:bank_code;"`
}

func (BankVa) TableName() string {
	return "banks_virtual_account"
}
