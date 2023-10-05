package entity

import "time"

type AppVersion struct {
	ID        string    `json:"id" gorm:"column:id"`
	OS        string    `json:"os" gorm:"column:os"`
	State     string    `json:"state" gorm:"column:state"`
	Ver       string    `json:"ver" gorm:"column:ver"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (*AppVersion) TableName() string {
	return "app_version"
}
