package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/middleware"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/request"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type CartControllerInterface interface {
	FindCartByIdUser(c echo.Context) error
	AddProductToCart(c echo.Context) error
	CartPlusQtyProduct(c echo.Context) error
	CartMinQtyProduct(c echo.Context) error
	UpdateQtyProductInCart(c echo.Context) error
}

type CartControllerImplementation struct {
	ConfigWebserver      config.Webserver
	Logger               *logrus.Logger
	CartServiceInterface services.CartServiceInterface
}

func NewCartController(configWebserver config.Webserver, cartServiceInterface services.CartServiceInterface) CartControllerInterface {
	return &CartControllerImplementation{
		ConfigWebserver:      configWebserver,
		CartServiceInterface: cartServiceInterface,
	}
}

func (controller *CartControllerImplementation) FindCartByIdUser(c echo.Context) error {
	requestId := ""
	IdUser := middleware.TokenClaimsIdUser(c)
	IdKelurahan := middleware.TokenClaimsIdKelurahan(c)
	cartResponses := controller.CartServiceInterface.FindCartByIdUser(requestId, IdUser, IdKelurahan)
	responses := response.Response{Code: 200, Mssg: "success", Data: cartResponses, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}

func (controller *CartControllerImplementation) AddProductToCart(c echo.Context) error {
	requestId := ""
	IdUser := middleware.TokenClaimsIdUser(c)
	request := request.ReadFromAddProductToCartRequestBody(c, requestId, controller.Logger)
	cartResponse := controller.CartServiceInterface.AddProductToCart(requestId, IdUser, request)
	response := response.Response{Code: 201, Mssg: "success add product to cart", Data: cartResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *CartControllerImplementation) CartPlusQtyProduct(c echo.Context) error {
	requestId := ""
	request := request.ReadFromUpdateProductInCartRequestBody(c, requestId, controller.Logger)
	cartResponseResult := controller.CartServiceInterface.CartPlusQtyProduct(requestId, request)
	response := response.Response{Code: 201, Mssg: "success add qty product", Data: cartResponseResult, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *CartControllerImplementation) CartMinQtyProduct(c echo.Context) error {
	requestId := ""
	request := request.ReadFromUpdateProductInCartRequestBody(c, requestId, controller.Logger)
	cartResponseResult := controller.CartServiceInterface.CartMinQtyProduct(requestId, request)
	response := response.Response{Code: 201, Mssg: "success reduce qty product", Data: cartResponseResult, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *CartControllerImplementation) UpdateQtyProductInCart(c echo.Context) error {
	requestId := ""
	request := request.ReadFromUpdateProductInCartRequestBody(c, requestId, controller.Logger)
	cartResponseResult := controller.CartServiceInterface.UpdateQtyProductInCart(requestId, request)
	response := response.Response{Code: 201, Mssg: "success update qty product", Data: cartResponseResult, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}
