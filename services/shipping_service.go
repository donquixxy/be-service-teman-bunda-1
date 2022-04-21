package services

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type ShippingServiceInterface interface {
	GetShippingCostByIdKelurahan(requestId string, idKlurahan int) (shippingCostAreaResponse response.GetShippingCostByIdKelurahanResponse)
}

type ShippingServiceImplementation struct {
	ConfigWebserver             config.Webserver
	DB                          *gorm.DB
	Validate                    *validator.Validate
	Logger                      *logrus.Logger
	ShippingRepositoryInterface mysql.ShippingRepositoryInterface
}

func NewShippingService(
	configWebserver config.Webserver,
	DB *gorm.DB,
	validate *validator.Validate,
	logger *logrus.Logger,
	shippingRepositoryInterface mysql.ShippingRepositoryInterface) ShippingServiceInterface {
	return &ShippingServiceImplementation{
		ConfigWebserver:             configWebserver,
		DB:                          DB,
		Validate:                    validate,
		Logger:                      logger,
		ShippingRepositoryInterface: shippingRepositoryInterface}
}

func (service *ShippingServiceImplementation) GetShippingCostByIdKelurahan(requestId string, idKelurahan int) (shippingCostResponse response.GetShippingCostByIdKelurahanResponse) {
	shippingCost, err := service.ShippingRepositoryInterface.GetShippingCostByIdKelurahan(service.DB, idKelurahan)
	fmt.Println(shippingCost)
	exceptions.PanicIfRecordNotFound(err, requestId, []string{"Data not found"}, service.Logger)
	shippingCostResponse = response.ToGetShippingCostByIdKelurahanResponse(shippingCost)
	return shippingCostResponse
}
