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

type KecamatanControllerInterface interface {
	FindAllKecamatanByIdKabupaten(c echo.Context) error
}

type KecamatanControllerImplementation struct {
	ConfigWebserver           config.Webserver
	Logger                    *logrus.Logger
	KecamatanServiceInterface services.KecamatanServiceInterface
}

func NewKecamatanController(configWebserver config.Webserver, kecamatanServiceInterface services.KecamatanServiceInterface) KecamatanControllerInterface {
	return &KecamatanControllerImplementation{
		ConfigWebserver:           configWebserver,
		KecamatanServiceInterface: kecamatanServiceInterface,
	}
}

func (controller *KecamatanControllerImplementation) FindAllKecamatanByIdKabupaten(c echo.Context) error {
	requestId := ""
	id, _ := strconv.Atoi(c.Param("id"))
	kecamatanResponses := controller.KecamatanServiceInterface.FindAllKecamatanByIdKabupaten(requestId, id)
	responses := response.Response{Code: 200, Mssg: "success", Data: kecamatanResponses, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}
