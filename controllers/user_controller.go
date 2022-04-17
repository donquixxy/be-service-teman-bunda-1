package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/request"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type UserControllerInterface interface {
	CreateUser(c echo.Context) error
	FindUserByReferal(c echo.Context) error
	FindUserById(c echo.Context) error
}

type UserControllerImplementation struct {
	ConfigurationWebserver config.Webserver
	Logger                 *logrus.Logger
	UserServiceInterface   services.UserServiceInterface
}

func NewUserController(configurationWebserver config.Webserver,
	logger *logrus.Logger,
	userServiceInterface services.UserServiceInterface) UserControllerInterface {
	return &UserControllerImplementation{
		ConfigurationWebserver: configurationWebserver,
		Logger:                 logger,
		UserServiceInterface:   userServiceInterface,
	}
}

func (controller *UserControllerImplementation) CreateUser(c echo.Context) error {
	requestId := ""
	request := request.ReadFromCreateUserRequestBody(c, requestId, controller.Logger)
	userResponse := controller.UserServiceInterface.CreateUser(requestId, request)
	response := response.Response{Code: "201", Mssg: "success", Data: userResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *UserControllerImplementation) FindUserByReferal(c echo.Context) error {
	requestId := ""
	referal := c.Param("referal")
	userResponse := controller.UserServiceInterface.FindUserByReferal(requestId, referal)
	response := response.Response{Code: "200", Mssg: "success", Data: userResponse.ReferalCode, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *UserControllerImplementation) FindUserById(c echo.Context) error {
	requestId := ""
	id := c.Param("id")
	userResponse := controller.UserServiceInterface.FindUserById(requestId, id)
	response := response.Response{Code: "200", Mssg: "success", Data: userResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}
