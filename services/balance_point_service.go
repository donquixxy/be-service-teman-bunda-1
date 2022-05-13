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
	FindBalancePointByIdUser(requestId string, IdUser string) (balancePointResponses response.FindBalancePointByIdUser)
	BalancePointCheckAmount(requestId string, IdUser string, amount float64) string
	BalancePointCheckOrderTx(requestId string, IdUser string, totalBill float64) string
}

type BalancePointServiceImplementation struct {
	ConfigWebserver                 config.Webserver
	DB                              *gorm.DB
	Logger                          *logrus.Logger
	BalancePointRepositoryInterface mysql.BalancePointRepositoryInterface
	SettingsRepositoryInterface     mysql.SettingRepositoryInterface
	OrderRepositoryInterface        mysql.OrderRepositoryInterface
}

func NewBalancePointService(configWebserver config.Webserver,
	DB *gorm.DB,
	logger *logrus.Logger,
	balancePointRepositoryInterface mysql.BalancePointRepositoryInterface,
	settingsRepositoryInterface mysql.SettingRepositoryInterface,
	orderRepositoryInterface mysql.OrderRepositoryInterface) BalancePointServiceInterface {
	return &BalancePointServiceImplementation{
		ConfigWebserver:                 configWebserver,
		DB:                              DB,
		Logger:                          logger,
		BalancePointRepositoryInterface: balancePointRepositoryInterface,
		SettingsRepositoryInterface:     settingsRepositoryInterface,
		OrderRepositoryInterface:        orderRepositoryInterface,
	}
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

func (service *BalancePointServiceImplementation) BalancePointCheckOrderTx(requestId string, idUser string, totalBill float64) string {

	// get limit order to use point value
	settings, _ := service.SettingsRepositoryInterface.FindSettingsByName(service.DB, "limit_order")
	if settings.SettingsName == "" {
		exceptions.PanicIfRecordNotFound(errors.New("settings not found"), requestId, []string{"Not Found"}, service.Logger)
	}

	// cek apakah jumlah order bulan ini sudah sesuai dengan limit order untuk menggunakan point
	// Get data order bulan ini
	orders, _ := service.OrderRepositoryInterface.FindOrderByDate(service.DB, idUser)
	var totalOrderAcumulate float64
	for _, order := range orders {
		totalOrderAcumulate = totalOrderAcumulate + order.PaymentByCash
	}
	totalOrderAcumulate = totalOrderAcumulate + totalBill

	if totalOrderAcumulate < settings.Value {
		exceptions.PanicIfBadRequest(errors.New("cant use point"), requestId, []string{"akumulasi total belanja kurang"}, service.Logger)
	}

	return "ok"
}

func (service *BalancePointServiceImplementation) BalancePointCheckAmount(requestId string, idUser string, amount float64) string {
	//check balance point
	balancePoint, _ := service.BalancePointRepositoryInterface.BalancePointUseCheck(service.DB, idUser)
	if balancePoint.IdUser == "" {
		exceptions.PanicIfRecordNotFound(errors.New("user not found"), requestId, []string{"Not Found"}, service.Logger)
	}

	if amount > balancePoint.BalancePoints {
		exceptions.PanicIfBadRequest(errors.New("point not enough"), requestId, []string{"point not enough"}, service.Logger)
	}

	return "ok"
}
