package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/middleware"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/request"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type OrderControllerInterface interface {
	CreateOrder(c echo.Context) error
}

type OrderControllerImplementation struct {
	ConfigurationWebserver config.Webserver
	Logger                 *logrus.Logger
	OrderServiceInterface  services.OrderServiceInterface
}

func NewOrderController(configurationWebserver config.Webserver,
	logger *logrus.Logger,
	orderServiceInterface services.OrderServiceInterface) OrderControllerInterface {
	return &OrderControllerImplementation{
		ConfigurationWebserver: configurationWebserver,
		Logger:                 logger,
		OrderServiceInterface:  orderServiceInterface,
	}
}

func (controller *OrderControllerImplementation) CreateOrder(c echo.Context) error {
	requestId := ""
	idUser := middleware.TokenClaimsIdUser(c)
	request := request.ReadFromCreateOrderRequestBody(c, requestId, controller.Logger)
	orderResponse := controller.OrderServiceInterface.CreateOrder(requestId, idUser, request)
	response := response.Response{Code: 201, Mssg: "order created", Data: orderResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}
