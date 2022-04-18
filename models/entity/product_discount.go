package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductDiscount struct {
	Id         string          `gorm:"primaryKey;column:id;"`
	IdProduct  string          `gorm:"column:id_product;"`
	Percentage decimal.Decimal `gorm:"column:percentage;"`
	Nominal    decimal.Decimal `gorm:"column:nominal;"`
	StartDate  time.Time       `gorm:"column:start_date;"`
	EndDate    time.Time       `gorm:"column:end_date;"`
}

func (ProductDiscount) TableName() string {
	return "products_discount"
}
