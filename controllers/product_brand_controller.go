package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type ProductBrandControllerInterface interface {
	FindAllProductBrand(c echo.Context) error
}

type ProductBrandControllerImplementation struct {
	ConfigWebserver              config.Webserver
	Logger                       *logrus.Logger
	ProductBrandServiceInterface services.ProductBrandServiceInterface
}

func NewProductBrandController(configWebserver config.Webserver, productBrandServiceInterface services.ProductBrandServiceInterface) ProductBrandControllerInterface {
	return &ProductBrandControllerImplementation{
		ConfigWebserver:              configWebserver,
		ProductBrandServiceInterface: productBrandServiceInterface,
	}
}

func (controller *ProductBrandControllerImplementation) FindAllProductBrand(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	productBrandResponses := controller.ProductBrandServiceInterface.FindAllProductBrand(requestId)
	responses := response.Response{Code: 200, Mssg: "success", Data: productBrandResponses, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}
