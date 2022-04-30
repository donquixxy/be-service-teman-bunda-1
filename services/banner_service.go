package services

import (
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type BannerServiceInterface interface {
	FindAllBanner(requestId string) (bannerResponses []response.FindAllBannerResponse)
}

type BannerServiceImplementation struct {
	ConfigWebserver           config.Webserver
	DB                        *gorm.DB
	Logger                    *logrus.Logger
	BannerRepositoryInterface mysql.BannerRepositoryInterface
}

func NewBannerService(configWebserver config.Webserver, DB *gorm.DB, logger *logrus.Logger, bannerRepositoryInterface mysql.BannerRepositoryInterface) BannerServiceInterface {
	return &BannerServiceImplementation{
		ConfigWebserver:           configWebserver,
		DB:                        DB,
		Logger:                    logger,
		BannerRepositoryInterface: bannerRepositoryInterface,
	}
}

func (service *BannerServiceImplementation) FindAllBanner(requestId string) (bannerResponses []response.FindAllBannerResponse) {
	banners, err := service.BannerRepositoryInterface.FindAllBanner(service.DB)
	exceptions.PanicIfError(err, requestId, service.Logger)
	bannerResponses = response.ToFindAllBannerResponse(banners)
	return bannerResponses
}
