package entity

type BankTransfer struct {
	Id        string `gorm:"primaryKey;column:id;"`
	BankName  string `gorm:"column:bank_name;"`
	BankLogo  string `gorm:"column:bank_logo;"`
	BankCode  string `gorm:"column:bank_code;"`
	NoAccount string `gorm:"column:no_account;"`
	BankAn    string `gorm:"column:bank_an;"`
}

func (BankTransfer) TableName() string {
	return "banks_manual_transfer"
}
