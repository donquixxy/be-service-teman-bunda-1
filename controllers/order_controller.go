package controllers

import (
	"fmt"
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
	UpdateStatusOrder(c echo.Context) error
	SendRequestToIpaymu(c echo.Context) error
	FindOrderByUser(c echo.Context) error
	FindOrderById(c echo.Context) error
	CancelOrderById(c echo.Context) error
	CompleteOrderById(c echo.Context) error
	OrderCheckPayment(c echo.Context) error
	// SendTelegram(c echo.Context) error
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

// func (controller *OrderControllerImplementation) SendTelegram(c echo.Context) error {
// 	controller.OrderServiceInterface.SendTelegram()
// 	response := response.Response{Code: 200, Mssg: "success", Data: "Nice", Error: []string{}}
// 	return c.JSON(http.StatusOK, response)
// }

func (controller *OrderControllerImplementation) FindOrderByUser(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idUser := middleware.TokenClaimsIdUser(c)
	orderStatus := c.QueryParam("order_status")
	orderResponse := controller.OrderServiceInterface.FindOrderByUser(requestId, idUser, orderStatus)
	response := response.Response{Code: 200, Mssg: "success", Data: orderResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *OrderControllerImplementation) FindOrderById(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idOrder := c.QueryParam("id_order")
	orderResponse := controller.OrderServiceInterface.FindOrderById(requestId, idOrder)
	response := response.Response{Code: 200, Mssg: "success", Data: orderResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *OrderControllerImplementation) OrderCheckPayment(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idOrder := c.QueryParam("id_order")
	orderResponse := controller.OrderServiceInterface.OrderCheckPayment(requestId, idOrder)
	response := response.Response{Code: 200, Mssg: "success", Data: orderResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *OrderControllerImplementation) CreateOrder(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idUser := middleware.TokenClaimsIdUser(c)
	request := request.ReadFromCreateOrderRequestBody(c, requestId, controller.Logger)
	orderResponse := controller.OrderServiceInterface.CreateOrder(requestId, idUser, request)
	response := response.Response{Code: 201, Mssg: "order created", Data: orderResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *OrderControllerImplementation) UpdateStatusOrder(c echo.Context) error {
	fmt.Println("Log Ada Request Ke Sini")
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	request := request.ReadFromCallBackIpaymuRequest(c, requestId, controller.Logger)
	orderResponse := controller.OrderServiceInterface.UpdateStatusOrder(requestId, request)
	response := response.Response{Code: 200, Mssg: "succes update status order", Data: orderResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *OrderControllerImplementation) CancelOrderById(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idOrder := c.QueryParam("id_order")
	err := controller.OrderServiceInterface.CancelOrderById(requestId, idOrder)
	response := response.Response{Code: 201, Mssg: "succes cancel order", Data: err, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *OrderControllerImplementation) CompleteOrderById(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idOrder := c.QueryParam("id_order")
	err := controller.OrderServiceInterface.CompleteOrderById(requestId, idOrder)
	response := response.Response{Code: 201, Mssg: "succes compelte order", Data: err, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *OrderControllerImplementation) SendRequestToIpaymu(c echo.Context) error {
	response := response.Response{Code: 201, Mssg: "order created", Data: nil, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}
