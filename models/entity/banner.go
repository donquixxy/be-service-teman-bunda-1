package entity

type Banner struct {
	Id          string `gorm:"primaryKey;column:id;"`
	BannerTitle string `gorm:"column:banner_title;"`
	BannerUrl   string `gorm:"column:banner_url;"`
}

func (Banner) TableName() string {
	return "banner"
}
