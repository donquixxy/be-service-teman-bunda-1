package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/middleware"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type BalancePointControllerInterface interface {
	FindBalancePointByIdUser(c echo.Context) error
	BalancePointCheckAmount(c echo.Context) error
	BalancePointCheckOrderTx(c echo.Context) error
}

type BalancePointControllerImplementation struct {
	ConfigWebserver              config.Webserver
	Logger                       *logrus.Logger
	BalancePointServiceInterface services.BalancePointServiceInterface
}

func NewBalancePointController(configWebserver config.Webserver,
	logger *logrus.Logger,
	balancePointServiceInterface services.BalancePointServiceInterface) BalancePointControllerInterface {
	return &BalancePointControllerImplementation{
		ConfigWebserver:              configWebserver,
		Logger:                       logger,
		BalancePointServiceInterface: balancePointServiceInterface,
	}
}

func (controller *BalancePointControllerImplementation) FindBalancePointByIdUser(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	IdUser := middleware.TokenClaimsIdUser(c)
	balancePointResponse := controller.BalancePointServiceInterface.FindBalancePointByIdUser(requestId, IdUser)
	response := response.Response{Code: 200, Mssg: "success", Data: balancePointResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *BalancePointControllerImplementation) BalancePointCheckAmount(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idUser := middleware.TokenClaimsIdUser(c)
	amount, _ := strconv.ParseFloat(c.QueryParam("amount"), 64)
	balancePointCheckResponse := controller.BalancePointServiceInterface.BalancePointCheckAmount(requestId, idUser, amount)
	response := response.Response{Code: 200, Mssg: "success", Data: balancePointCheckResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *BalancePointControllerImplementation) BalancePointCheckOrderTx(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idUser := middleware.TokenClaimsIdUser(c)
	balancePointCheckResponse := controller.BalancePointServiceInterface.BalancePointCheckOrderTx(requestId, idUser)
	response := response.Response{Code: 200, Mssg: "success", Data: balancePointCheckResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}
