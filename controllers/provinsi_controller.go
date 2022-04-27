package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type ProvinsiControllerInterface interface {
	FindAllProvinsi(c echo.Context) error
}

type ProvinsiControllerImplementation struct {
	ConfigWebserver          config.Webserver
	Logger                   *logrus.Logger
	ProvinsiServiceInterface services.ProvinsiServiceInterface
}

func NewProvinsiController(configWebserver config.Webserver, provinsiServiceInterface services.ProvinsiServiceInterface) ProvinsiControllerInterface {
	return &ProvinsiControllerImplementation{
		ConfigWebserver:          configWebserver,
		ProvinsiServiceInterface: provinsiServiceInterface,
	}
}

func (controller *ProvinsiControllerImplementation) FindAllProvinsi(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	provinsiResponses := controller.ProvinsiServiceInterface.FindAllProvinsi(requestId)
	responses := response.Response{Code: 200, Mssg: "success", Data: provinsiResponses, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}
