package entity

import "time"

type UserLevelMember struct {
	Id              string    `gorm:"primaryKey;column:id;"`
	LevelName       string    `gorm:"column:level_name;"`
	BonusPercentage float64   `gorm:"column:bonus_percentage;"`
	CreatedDate     time.Time `gorm:"column:created_at;"`
}

func (UserLevelMember) TableName() string {
	return "users_level_member"
}
