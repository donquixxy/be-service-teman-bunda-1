package entity

type ServiceZonaArea struct {
	Id           int     `gorm:"primaryKey;column:id;"`
	ZonaName     string  `gorm:"column:zona_name;"`
	ShippingCost float64 `gorm:"column:shipping_cost;"`
}

func (ServiceZonaArea) TableName() string {
	return "service_zona_area"
}
