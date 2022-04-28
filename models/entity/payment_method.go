package entity

type PaymentMethod struct {
	PaymentMethod string `gorm:"column:payment_method;"`
}

func (PaymentMethod) TableName() string {
	return "payment_method"
}
