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

type KelurahanControllerInterface interface {
	FindAllKelurahanByIdKecamatan(c echo.Context) error
}

type KelurahanControllerImplementation struct {
	ConfigWebserver           config.Webserver
	Logger                    *logrus.Logger
	KelurahanServiceInterface services.KelurahanServiceInterface
}

func NewKelurahanController(configWebserver config.Webserver, kelurahanServiceInterface services.KelurahanServiceInterface) KelurahanControllerInterface {
	return &KelurahanControllerImplementation{
		ConfigWebserver:           configWebserver,
		KelurahanServiceInterface: kelurahanServiceInterface,
	}
}

func (controller *KelurahanControllerImplementation) FindAllKelurahanByIdKecamatan(c echo.Context) error {
	requestId := ""
	id, _ := strconv.Atoi(c.Param("id"))
	kelurahanResponses := controller.KelurahanServiceInterface.FindAllKelurahanByIdKecamatan(requestId, id)
	responses := response.Response{Code: 200, Mssg: "success", Data: kelurahanResponses, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}
