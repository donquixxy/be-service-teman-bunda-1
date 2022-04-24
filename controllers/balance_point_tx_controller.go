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

type BalancePointTxControllerInterface interface {
	FindBalancePointTxByIdBalancePoint(c echo.Context) error
}

type BalancePointTxControllerImplementation struct {
	ConfigWebserver                config.Webserver
	Logger                         *logrus.Logger
	BalancePointTxServiceInterface services.BalancePointTxServiceInterface
}

func NewBalancePointTxController(configWebserver config.Webserver,
	logger *logrus.Logger,
	balancePointTxServiceInterface services.BalancePointTxServiceInterface) BalancePointTxControllerInterface {
	return &BalancePointTxControllerImplementation{
		ConfigWebserver:                configWebserver,
		Logger:                         logger,
		BalancePointTxServiceInterface: balancePointTxServiceInterface,
	}
}

func (controller *BalancePointTxControllerImplementation) FindBalancePointTxByIdBalancePoint(c echo.Context) error {
	requestId := ""
	idUser := middleware.TokenClaimsIdUser(c)
	date := c.QueryParam("date")
	balancePointWithTxResponse := controller.BalancePointTxServiceInterface.FindBalancePointWithTxByIdBalancePoint(requestId, date, idUser)
	response := response.Response{Code: 200, Mssg: "success", Data: balancePointWithTxResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}
