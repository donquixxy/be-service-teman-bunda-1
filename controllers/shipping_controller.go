package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/middleware"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
	"gorm.io/gorm"
)

type ShippingControllerInterface interface {
	GetShippingCostByIdKelurahan(c echo.Context) error
}

type ShippingControllerImplementation struct {
	ConfigurationWebserver   config.Webserver
	DB                       *gorm.DB
	Validate                 *validator.Validate
	ShippingServiceInterface services.ShippingServiceInterface
}

func NewShippingController(configurationWebserver config.Webserver,
	shippingServiceInterface services.ShippingServiceInterface) ShippingControllerInterface {
	return &ShippingControllerImplementation{
		ConfigurationWebserver:   configurationWebserver,
		ShippingServiceInterface: shippingServiceInterface,
	}
}

func (controller *ShippingControllerImplementation) GetShippingCostByIdKelurahan(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idKelurahan := middleware.TokenClaimsIdKelurahan(c)
	fmt.Println("id kelurahan", idKelurahan)
	userResponse := controller.ShippingServiceInterface.GetShippingCostByIdKelurahan(requestId, idKelurahan)
	response := response.Response{Code: 200, Mssg: "success", Data: userResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}
