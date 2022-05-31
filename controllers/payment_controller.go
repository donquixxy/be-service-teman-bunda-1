package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/request"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type PaymentControllerInterface interface {
	PaymentStatus(c echo.Context) error
	PaymentCreditCardSuccess(c echo.Context) error
	PaymentCreditCardCancel(c echo.Context) error
}

type PaymentControllerImplementation struct {
	ConfigurationWebserver  config.Webserver
	Logger                  *logrus.Logger
	PaymentServiceInterface services.PaymentServiceInterface
}

func NewPaymentController(configurationWebserver config.Webserver,
	logger *logrus.Logger,
	paymentServiceInterface services.PaymentServiceInterface) PaymentControllerInterface {
	return &PaymentControllerImplementation{
		ConfigurationWebserver:  configurationWebserver,
		Logger:                  logger,
		PaymentServiceInterface: paymentServiceInterface,
	}
}

func (controller *PaymentControllerImplementation) PaymentCreditCardSuccess(c echo.Context) error {
	return c.File("./template/credit_card_success_payment.html")
}

func (controller *PaymentControllerImplementation) PaymentCreditCardCancel(c echo.Context) error {
	return c.File("./template/credit_card_cancel_payment.html")
}

func (controller *PaymentControllerImplementation) PaymentStatus(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	request := request.ReadFromPaymentStatusRequestBody(c, requestId, controller.Logger)
	paymentResponse := controller.PaymentServiceInterface.PaymentStatus(requestId, request)
	response := response.Response{Code: 200, Mssg: "success", Data: paymentResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}
