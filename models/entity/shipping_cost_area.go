package entity

type ShippingCostArea struct {
	Id           int     `gorm:"primaryKey;column:id;"`
	ZonaName     string  `gorm:"column:zona_name;"`
	ShippingCost float64 `gorm:"column:shipping_cost;"`
}

func (ShippingCostArea) TableName() string {
	return "shipping_cost_area"
}
