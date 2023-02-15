package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
)

type MainControllerInterface interface {
	Main(c echo.Context) error
}

type MainControllerImplementation struct {
	ConfigurationWebserver config.Webserver
}

func NewMainController(configurationWebserver config.Webserver) MainControllerInterface {
	return &MainControllerImplementation{
		ConfigurationWebserver: configurationWebserver,
	}
}

func (controller *MainControllerImplementation) Main(c echo.Context) error {
	respon := response.Response{Code: 200, Mssg: "success", Data: "Welcome To AETHER!", Error: []string{}}
	return c.JSON(http.StatusOK, respon)
}
