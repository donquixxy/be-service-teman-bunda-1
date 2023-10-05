package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type AppVersionRepository interface {
	FindSettingByID(db *gorm.DB, id string) (*entity.AppVersion, error)
	FindVersionApp(db *gorm.DB, os string) []*entity.AppVersion
	FindByValue(db *gorm.DB, os int) []*entity.AppVersion
}

type appVersionRepository struct {
}

func NewAppVersionRepository() AppVersionRepository {
	return &appVersionRepository{}
}

func (*appVersionRepository) FindByValue(db *gorm.DB, os int) []*entity.AppVersion {
	var result []*entity.AppVersion

	db.Where("value = ?", os).Find(&result)

	return result
}

func (*appVersionRepository) FindSettingByID(db *gorm.DB, id string) (*entity.AppVersion, error) {
	var result *entity.AppVersion

	res := db.Where("id = ?", id).First(&result)

	return result, res.Error
}

func (*appVersionRepository) FindVersionApp(db *gorm.DB, os string) []*entity.AppVersion {
	var result []*entity.AppVersion

	db.Where("os = ?", os).Find(&result)

	return result
}
