package services

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type BalancePointServiceInterface interface {
	FindBalancePointWithTxByIdUser(requestId string, IdUser string) (balancePointWithTxResponses response.FindBalancePointWithTxByIdUser)
	FindBalancePointByIdUser(requestId string, IdUser string) (balancePointResponses response.FindBalancePointByIdUser)
	BalancePointUseCheck(requestId string, IdUser string) (balancePointUseCheckResponse response.BalancePointUseCheck)
}

type BalancePointServiceImplementation struct {
	ConfigWebserver                 config.Webserver
	DB                              *gorm.DB
	Logger                          *logrus.Logger
	BalancePointRepositoryInterface mysql.BalancePointRepositoryInterface
	SettingsRepositoryInterface     mysql.SettingsRepositoryInterface
}

func NewBalancePointService(configWebserver config.Webserver,
	DB *gorm.DB,
	logger *logrus.Logger,
	balancePointRepositoryInterface mysql.BalancePointRepositoryInterface,
	settingsRepositoryInterface mysql.SettingsRepositoryInterface) BalancePointServiceInterface {
	return &BalancePointServiceImplementation{
		ConfigWebserver:                 configWebserver,
		DB:                              DB,
		Logger:                          logger,
		BalancePointRepositoryInterface: balancePointRepositoryInterface,
		SettingsRepositoryInterface:     settingsRepositoryInterface,
	}
}

func (service *BalancePointServiceImplementation) FindBalancePointWithTxByIdUser(requestId string, IdUser string) (balancePointWithTxResponse response.FindBalancePointWithTxByIdUser) {
	balancePointWithTx, _ := service.BalancePointRepositoryInterface.FindBalancePointWithTxByIdUser(service.DB, IdUser)
	if balancePointWithTx.IdUser == "" {
		err := errors.New("user not found")
		exceptions.PanicIfRecordNotFound(err, requestId, []string{"Not Found"}, service.Logger)
	}
	balancePointWithTxResponse = response.ToFindBalancePointWithTxByIdUser(balancePointWithTx)
	return balancePointWithTxResponse
}

func (service *BalancePointServiceImplementation) FindBalancePointByIdUser(requestId string, IdUser string) (balancePointResponse response.FindBalancePointByIdUser) {
	balancePoint, _ := service.BalancePointRepositoryInterface.FindBalancePointByIdUser(service.DB, IdUser)
	if balancePoint.IdUser == "" {
		err := errors.New("user not found")
		exceptions.PanicIfRecordNotFound(err, requestId, []string{"Not Found"}, service.Logger)
	}
	balancePointResponse = response.ToFindBalancePointByIdUser(balancePoint)
	return balancePointResponse
}

func (service *BalancePointServiceImplementation) BalancePointUseCheck(requestId string, IdUser string) (balancePointUseCheckResponse response.BalancePointUseCheck) {
	//check balance point
	balancePoint, _ := service.BalancePointRepositoryInterface.BalancePointUseCheck(service.DB, IdUser)
	if balancePoint.IdUser == "" {
		err := errors.New("user not found")
		exceptions.PanicIfRecordNotFound(err, requestId, []string{"Not Found"}, service.Logger)
	}

	//get limit order to use point value
	settings, _ := service.SettingsRepositoryInterface.FindSettingsByName(service.DB, "limit_order")
	if settings.SettingsName == "" {
		err := errors.New("settings not found")
		exceptions.PanicIfRecordNotFound(err, requestId, []string{"Not Found"}, service.Logger)
	}

	balancePointUseCheckResponse = response.ToBalancePointUseCheck(balancePoint)
	return balancePointUseCheckResponse
}
