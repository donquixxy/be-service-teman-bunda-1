package entity

import "time"

type PaymentLog struct {
	Id             string    `gorm:"primaryKey;column:id;"`
	IdOrder        string    `gorm:"column:id_order;"`
	NumberOrder    string    `gorm:"column:number_order;"`
	TypeLog        string    `gorm:"column:type_log;"`
	PaymentMethod  string    `gorm:"column:payment_method;"`
	PaymentChannel string    `gorm:"column:payment_channel;"`
	Log            string    `gorm:"column:log;"`
	CreatedAt      time.Time `gorm:"column:created_at;"`
}

func (PaymentLog) TableName() string {
	return "payment_log"
}
