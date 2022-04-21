package entity

type Product struct {
	Id              string          `gorm:"primaryKey;column:id;"`
	ProductName     string          `gorm:"column:product_name;"`
	Price           float64         `gorm:"column:price;"`
	Description     string          `gorm:"column:description;"`
	PictureUrl      string          `gorm:"column:picture_url;"`
	Thumbnail       string          `gorm:"column:thumbnail;"`
	Stock           int             `gorm:"column:stock;"`
	IdCategory      int             `gorm:"column:id_category;"`
	ProductCategory ProductCategory `gorm:"foreignKey:IdCategory"`
	ProductDiscount ProductDiscount `gorm:"foreignKey:IdProduct"`
}

func (Product) TableName() string {
	return "products"
}
