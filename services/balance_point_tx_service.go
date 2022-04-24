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

type BalancePointTxServiceInterface interface {
	FindBalancePointWithTxByIdBalancePoint(requestId string, date string, idUser string) (balancePointTxWithTxResponses []response.FindBalancePointTxByIdBalancePoint)
}

type BalancePointTxServiceImplementation struct {
	ConfigWebserver                   config.Webserver
	DB                                *gorm.DB
	Logger                            *logrus.Logger
	BalancePointTxRepositoryInterface mysql.BalancePointTxRepositoryInterface
	BalancePointRepositoryInterface   mysql.BalancePointRepositoryInterface
}

func NewBalancePointTxService(configWebserver config.Webserver,
	DB *gorm.DB,
	logger *logrus.Logger,
	balancePointTxRepositoryInterface mysql.BalancePointTxRepositoryInterface,
	balancePointRepositoryInterface mysql.BalancePointRepositoryInterface) BalancePointTxServiceInterface {
	return &BalancePointTxServiceImplementation{
		ConfigWebserver:                   configWebserver,
		DB:                                DB,
		Logger:                            logger,
		BalancePointTxRepositoryInterface: balancePointTxRepositoryInterface,
		BalancePointRepositoryInterface:   balancePointRepositoryInterface,
	}
}

func (service *BalancePointTxServiceImplementation) FindBalancePointWithTxByIdBalancePoint(requestId string, date string, idUser string) (balancePointTxResponses []response.FindBalancePointTxByIdBalancePoint) {
	// Get User balance point
	balancePoint, _ := service.BalancePointRepositoryInterface.FindBalancePointByIdUser(service.DB, idUser)
	if balancePoint.Id == "" {
		exceptions.PanicIfRecordNotFound(errors.New("data not found"), requestId, []string{"Data Not Found"}, service.Logger)
	}

	balancePointTx, _ := service.BalancePointTxRepositoryInterface.FindBalancePointTxByIdBalancePoint(service.DB, date, balancePoint.Id)
	if len(balancePointTx) == 0 {
		exceptions.PanicIfRecordNotFound(errors.New("data not found"), requestId, []string{"Data Not Found"}, service.Logger)
	}
	balancePointTxResponses = response.ToFindBalancePointTxByIdBalancePoint(balancePointTx)
	return balancePointTxResponses
}
