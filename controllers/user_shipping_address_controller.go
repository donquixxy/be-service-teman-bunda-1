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

type UserShippingAddressControllerInterface interface {
	FindUserShippingAddress(c echo.Context) error
	CreateUserShippingAddress(c echo.Context) error
	DeleteUserShippingAddress(c echo.Context) error
}

type UserShippingAddressControllerImplementation struct {
	ConfigWebserver                     config.Webserver
	Logger                              *logrus.Logger
	UserShippingAddressServiceInterface services.UserShippingAddressServiceInterface
}

func NewUserShippingAddressController(configWebserver config.Webserver, userShippingAddressServiceInterface services.UserShippingAddressServiceInterface) UserShippingAddressControllerInterface {
	return &UserShippingAddressControllerImplementation{
		ConfigWebserver:                     configWebserver,
		UserShippingAddressServiceInterface: userShippingAddressServiceInterface,
	}
}

func (controller *UserShippingAddressControllerImplementation) DeleteUserShippingAddress(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idUserAddress := c.QueryParam("id_user_address")
	err := controller.UserShippingAddressServiceInterface.DeleteUserShippingAddress(requestId, idUserAddress)
	if err == nil {
		responses := response.Response{Code: 200, Mssg: "success", Data: "", Error: []string{}}
		return c.JSON(http.StatusOK, responses)
	} else {
		return nil
	}
}

func (controller *UserShippingAddressControllerImplementation) CreateUserShippingAddress(c echo.Context) error {
	// fmt.Println("Masuk cretae user address")
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idUser := middleware.TokenClaimsIdUser(c)
	request := request.ReadFromCreateUserShippingAddressRequestBody(c, requestId, controller.Logger)
	err := controller.UserShippingAddressServiceInterface.CreateUserShippingAddress(requestId, idUser, request)
	if err == nil {
		responses := response.Response{Code: 200, Mssg: "success", Data: "", Error: []string{}}
		return c.JSON(http.StatusOK, responses)
	} else {
		return nil
	}
}

func (controller *UserShippingAddressControllerImplementation) FindUserShippingAddress(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idUser := middleware.TokenClaimsIdUser(c)
	userAddressResponses := controller.UserShippingAddressServiceInterface.FindUserShippingAddressByIdUser(requestId, idUser)
	responses := response.Response{Code: 200, Mssg: "success", Data: userAddressResponses, Error: []string{}}
	return c.JSON(http.StatusOK, responses)
}
