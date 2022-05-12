package entity

type PaymentChannel struct {
	Id                 int           `gorm:"column:id;"`
	IdPaymentMethod    int           `gorm:"column:id_payment_method;"`
	PaymentMethod      PaymentMethod `gorm:"foreignKey:IdPaymentMethod"`
	Name               string        `gorm:"column:name;"`
	Alias              string        `gorm:"column:alias;"`
	Code               string        `gorm:"column:code;"`
	Logo               string        `gorm:"column:logo;"`
	BankFee            float64       `gorm:"column:bank_fee;"`
	AdminFee           float64       `gorm:"column:admin_fee;"`
	BankFeePercentage  float64       `gorm:"column:bank_fee_percentage;"`
	AdminFeePercentage float64       `gorm:"column:admin_fee_percentage;"`
	IsActive           string        `gorm:"column:is_active;"`
}

func (PaymentChannel) TableName() string {
	return "payment_channel"
}
