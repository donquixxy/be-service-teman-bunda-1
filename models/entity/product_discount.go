package entity

import "time"

type ProductDiscount struct {
	Id         string    `gorm:"primaryKey;column:id;"`
	IdProduct  string    `gorm:"column:id_product;"`
	Percentage string    `gorm:"column:percentage;"`
	Nominal    string    `gorm:"column:nominal;"`
	StartDate  time.Time `gorm:"column:start_date;"`
	EndDate    time.Time `gorm:"column:end_date;"`
}

func (ProductDiscount) TableName() string {
	return "products_discount"
}
