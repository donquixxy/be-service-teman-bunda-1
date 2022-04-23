package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type SettingsRepositoryInterface interface {
	FindSettingsByName(DB *gorm.DB, settingName string) (entity.Settings, error)
}

type SettingsRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewSettingsRepository(configDatabase *config.Database) SettingsRepositoryInterface {
	return &SettingsRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *SettingsRepositoryImplementation) FindSettingsByName(DB *gorm.DB, settingName string) (entity.Settings, error) {
	var settings entity.Settings
	results := DB.Where("settings_name = ?", settingName).First(&settings)
	return settings, results.Error
}
