package entity

type ProductBrand struct {
	Id           string `gorm:"primaryKey;column:id;"`
	BrandName    string `gorm:"column:brand_name;"`
	BrandLogoUrl string `gorm:"column:brand_logo_url;"`
}

func (ProductBrand) TableName() string {
	return "products_brand"
}
