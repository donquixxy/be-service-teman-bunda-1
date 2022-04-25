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
	FindBalancePointWithTxByIdUser(c echo.Context) error
	FindBalancePointByIdUser(c echo.Context) error
	BalancePointUseCheck(c echo.Context) error
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

func (controller *BalancePointControllerImplementation) FindBalancePointWithTxByIdUser(c echo.Context) error {
	requestId := ""
	IdUser := middleware.TokenClaimsIdUser(c)
	balancePointWithTxResponse := controller.BalancePointServiceInterface.FindBalancePointWithTxByIdUser(requestId, IdUser)
	response := response.Response{Code: 200, Mssg: "success", Data: balancePointWithTxResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *BalancePointControllerImplementation) FindBalancePointByIdUser(c echo.Context) error {
	requestId := ""
	IdUser := middleware.TokenClaimsIdUser(c)
	balancePointResponse := controller.BalancePointServiceInterface.FindBalancePointByIdUser(requestId, IdUser)
	response := response.Response{Code: 200, Mssg: "success", Data: balancePointResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *BalancePointControllerImplementation) BalancePointUseCheck(c echo.Context) error {
	requestId := ""
	idUser := middleware.TokenClaimsIdUser(c)
	amount, _ := strconv.ParseFloat(c.QueryParam("amount"), 64)
	balancePointUseCheckResponse := controller.BalancePointServiceInterface.BalancePointUseCheck(requestId, idUser, amount)
	response := response.Response{Code: 200, Mssg: "success", Data: balancePointUseCheckResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}
