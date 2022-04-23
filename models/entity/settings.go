package entity

type Settings struct {
	Id            int     `gorm:"primary_key;column:id"`
	SettingsName  string  `gorm:"column:settings_name"`
	SettingsTitle string  `gorm:"column:settings_title"`
	Value         float64 `gorm:"column:value"`
}

func (Settings) TableName() string {
	return "settings"
}
