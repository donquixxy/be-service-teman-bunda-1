package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type BannerRepositoryInterface interface {
	FindAllBanner(DB *gorm.DB) ([]entity.Banner, error)
}

type BannerRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewBannerRepository(configDatabase *config.Database) BannerRepositoryInterface {
	return &BannerRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *BannerRepositoryImplementation) FindAllBanner(DB *gorm.DB) ([]entity.Banner, error) {
	var banners []entity.Banner
	results := DB.Find(&banners)
	return banners, results.Error
}
