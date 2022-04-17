package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type BalancePointControllerInterface interface {
	FindBalancePointWithTxByIdUser(c echo.Context) error
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
	IdUser := c.Param("id")
	balancePointWithTxResponse := controller.BalancePointServiceInterface.FindBalancePointWithTxByIdUser(requestId, IdUser)
	response := response.Response{Code: "200", Mssg: "success", Data: balancePointWithTxResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}
