package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type SettingControllerInterface interface {
	FindSettingShippingCost(c echo.Context) error
	FindSettingVerApp(c echo.Context) error
}

type SettingControllerImplementation struct {
	ConfigWebserver         config.Webserver
	Logger                  *logrus.Logger
	SettingServiceInterface services.SettingServiceInterface
}

func NewSettingController(configWebserver config.Webserver, settingServiceInterface services.SettingServiceInterface) SettingControllerInterface {
	return &SettingControllerImplementation{
		ConfigWebserver:         configWebserver,
		SettingServiceInterface: settingServiceInterface,
	}
}

func (controller *SettingControllerImplementation) FindSettingShippingCost(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	shippingCost := controller.SettingServiceInterface.FindSettingShippingCost(requestId)
	responses := response.Response{Code: 200, Mssg: "success", Data: shippingCost, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}

func (controller *SettingControllerImplementation) FindSettingVerApp(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	verApp := controller.SettingServiceInterface.FindSettingVerApp(requestId)
	responses := response.Response{Code: 200, Mssg: "success", Data: verApp, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}
