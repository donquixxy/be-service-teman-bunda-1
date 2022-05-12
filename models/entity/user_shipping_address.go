package entity

type UserShippingAddress struct {
	Id        string  `gorm:"primaryKey;column:id;"`
	IdUser    string  `gorm:"column:id_user;"`
	Status    int     `gorm:"column:status;"`
	Address   string  `gorm:"column:address;"`
	Latitude  float64 `gorm:"column:latitude;"`
	Longitude float64 `gorm:"column:longitude;"`
	Radius    float64 `gorm:"column:radius;"`
}

func (UserShippingAddress) TableName() string {
	return "users_shipping_address"
}
