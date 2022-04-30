package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type BannerControllerInterface interface {
	FindAllBanner(c echo.Context) error
}

type BannerControllerImplementation struct {
	ConfigWebserver        config.Webserver
	Logger                 *logrus.Logger
	BannerServiceInterface services.BannerServiceInterface
}

func NewBannerController(configWebserver config.Webserver, bannerServiceInterface services.BannerServiceInterface) BannerControllerInterface {
	return &BannerControllerImplementation{
		ConfigWebserver:        configWebserver,
		BannerServiceInterface: bannerServiceInterface,
	}
}

func (controller *BannerControllerImplementation) FindAllBanner(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	bannerResponses := controller.BannerServiceInterface.FindAllBanner(requestId)
	responses := response.Response{Code: 200, Mssg: "success", Data: bannerResponses, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}
