package entity

type ProductCategory struct {
	Id           string `gorm:"primaryKey;column:id;"`
	CategoryName string `gorm:"column:category_name;"`
}

func (ProductCategory) TableName() string {
	return "products_category"
}
