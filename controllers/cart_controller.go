package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/middleware"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type CartControllerInterface interface {
	FindCartByIdUser(c echo.Context) error
}

type CartControllerImplementation struct {
	ConfigWebserver      config.Webserver
	Logger               *logrus.Logger
	CartServiceInterface services.CartServiceInterface
}

func NewCartController(configWebserver config.Webserver, cartServiceInterface services.CartServiceInterface) CartControllerInterface {
	return &CartControllerImplementation{
		ConfigWebserver:      configWebserver,
		CartServiceInterface: cartServiceInterface,
	}
}

func (controller *CartControllerImplementation) FindCartByIdUser(c echo.Context) error {
	requestId := ""
	IdUser := middleware.TokenClaimsIdUser(c)
	cartResponses := controller.CartServiceInterface.FindCartByIdUser(requestId, IdUser)
	responses := response.Response{Code: "200", Mssg: "success", Data: cartResponses, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}
