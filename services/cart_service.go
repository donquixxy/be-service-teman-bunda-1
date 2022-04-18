package services

import (
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type CartServiceInterface interface {
	FindCartByIdUser(requestId string, IdUser string) (cartResponses response.FindCartByIdUser)
}

type CartServiceImplementation struct {
	ConfigWebserver         config.Webserver
	DB                      *gorm.DB
	Logger                  *logrus.Logger
	CartRepositoryInterface mysql.CartRepositoryInterface
}

func NewCartService(configWebserver config.Webserver, DB *gorm.DB, logger *logrus.Logger, cartRepositoryInterface mysql.CartRepositoryInterface) CartServiceInterface {
	return &CartServiceImplementation{
		ConfigWebserver:         configWebserver,
		DB:                      DB,
		Logger:                  logger,
		CartRepositoryInterface: cartRepositoryInterface,
	}
}

func (service *CartServiceImplementation) FindCartByIdUser(requestId string, IdUser string) (cartResponses response.FindCartByIdUser) {
	carts, err := service.CartRepositoryInterface.FindCartByIdUser(service.DB, IdUser)
	exceptions.PanicIfError(err, requestId, service.Logger)
	cartResponses = response.ToFindCartByIdUser(carts)
	return cartResponses
}
