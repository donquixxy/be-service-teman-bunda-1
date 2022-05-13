package services

import (
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type SettingServiceInterface interface {
	FindSettingShippingCost(requestId string) (settingShippingCost response.FindSettingShippingCost)
}

type SettingServiceImplementation struct {
	ConfigWebserver            config.Webserver
	DB                         *gorm.DB
	Logger                     *logrus.Logger
	SettingRepositoryInterface mysql.SettingRepositoryInterface
}

func NewSettingService(configWebserver config.Webserver, DB *gorm.DB, logger *logrus.Logger, settingRepositoryInterface mysql.SettingRepositoryInterface) SettingServiceInterface {
	return &SettingServiceImplementation{
		ConfigWebserver:            configWebserver,
		DB:                         DB,
		Logger:                     logger,
		SettingRepositoryInterface: settingRepositoryInterface,
	}
}

func (service *SettingServiceImplementation) FindSettingShippingCost(requestId string) (shippingCostResponse response.FindSettingShippingCost) {
	shippingCost, err := service.SettingRepositoryInterface.FindSettingShippingCost(service.DB)
	exceptions.PanicIfError(err, requestId, service.Logger)
	shippingCostResponse = response.ToFindSettingShippingCost(shippingCost)
	return shippingCostResponse
}
