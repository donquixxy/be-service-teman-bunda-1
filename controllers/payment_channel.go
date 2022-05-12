package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type PaymentChannelControllerInterface interface {
	FindListPaymentChannel(c echo.Context) error
	FindListPaymentChannelv2(c echo.Context) error
}

type PaymentChannelControllerImplementation struct {
	ConfigWebserver                config.Webserver
	Logger                         *logrus.Logger
	PaymentChannelServiceInterface services.PaymentChannelServiceInterface
}

func NewPaymentChannelController(configWebserver config.Webserver, logger *logrus.Logger, paymentChannelServiceInterface services.PaymentChannelServiceInterface) PaymentChannelControllerInterface {
	return &PaymentChannelControllerImplementation{
		ConfigWebserver:                configWebserver,
		Logger:                         logger,
		PaymentChannelServiceInterface: paymentChannelServiceInterface,
	}
}

func (controller *PaymentChannelControllerImplementation) FindListPaymentChannel(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	paymentChannelResponses := controller.PaymentChannelServiceInterface.FindListPaymentChannel(requestId)
	responses := response.Response{Code: 200, Mssg: "success", Data: paymentChannelResponses, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}

func (controller *PaymentChannelControllerImplementation) FindListPaymentChannelv2(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	paymentChannelResponses := controller.PaymentChannelServiceInterface.FindListPaymentChannelv2(requestId)
	responses := response.Response{Code: 200, Mssg: "success", Data: paymentChannelResponses, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}
