package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type KabupatenControllerInterface interface {
	FindAllKabupatenByIdProvinsi(c echo.Context) error
}

type KabupatenControllerImplementation struct {
	ConfigWebserver           config.Webserver
	Logger                    *logrus.Logger
	KabupatenServiceInterface services.KabupatenServiceInterface
}

func NewKabupatenController(configWebserver config.Webserver, kabupatenServiceInterface services.KabupatenServiceInterface) KabupatenControllerInterface {
	return &KabupatenControllerImplementation{
		ConfigWebserver:           configWebserver,
		KabupatenServiceInterface: kabupatenServiceInterface,
	}
}

func (controller *KabupatenControllerImplementation) FindAllKabupatenByIdProvinsi(c echo.Context) error {
	requestId := ""
	id, _ := strconv.Atoi(c.QueryParam("idprop"))
	kabupatenResponses := controller.KabupatenServiceInterface.FindAllKabupatenByIdProvinsi(requestId, id)
	responses := response.Response{Code: 200, Mssg: "success", Data: kabupatenResponses, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}
