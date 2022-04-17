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
}

type BalancePointServiceImplementation struct {
	ConfigWebserver                 config.Webserver
	DB                              *gorm.DB
	Logger                          *logrus.Logger
	BalancePointRepositoryInterface mysql.BalancePointRepositoryInterface
}

func NewBalancePointService(configWebserver config.Webserver,
	DB *gorm.DB,
	logger *logrus.Logger,
	balancePointRepositoryInterface mysql.BalancePointRepositoryInterface) BalancePointServiceInterface {
	return &BalancePointServiceImplementation{
		ConfigWebserver:                 configWebserver,
		DB:                              DB,
		Logger:                          logger,
		BalancePointRepositoryInterface: balancePointRepositoryInterface,
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
