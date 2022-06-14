package controllers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/middleware"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/request"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/services"
)

type UserControllerInterface interface {
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	FindUserByReferal(c echo.Context) error
	FindUserById(c echo.Context) error
	UpdateStatusActiveUser(c echo.Context) error
	PasswordCodeRequest(c echo.Context) error
	PasswordResetCodeVerify(c echo.Context) error
	UpdateUserPassword(c echo.Context) error
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
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	request := request.ReadFromCreateUserRequestBody(c, requestId, controller.Logger)
	userResponse := controller.UserServiceInterface.CreateUser(requestId, request)
	response := response.Response{Code: 201, Mssg: "user created", Data: userResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *UserControllerImplementation) UpdateUser(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idUser := middleware.TokenClaimsIdUser(c)
	request := request.ReadFromUpdateUserRequestBody(c, requestId, controller.Logger)
	userResponse := controller.UserServiceInterface.UpdateUser(requestId, idUser, request)
	response := response.Response{Code: 201, Mssg: "user updated", Data: userResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *UserControllerImplementation) UpdateUserPassword(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	request := request.ReadFromUpdateUserPasswordRequestBody(c, requestId, controller.Logger)
	userResponse := controller.UserServiceInterface.UpdateUserPassword(requestId, request)
	response := response.Response{Code: 201, Mssg: "user password updated", Data: userResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *UserControllerImplementation) PasswordResetCodeVerify(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	request := request.ReadFromPasswordResetCodeVerifyBody(c, requestId, controller.Logger)
	controller.UserServiceInterface.PasswordResetCodeVerify(requestId, request)
	response := response.Response{Code: 200, Mssg: "verification success"}
	return c.JSON(http.StatusOK, response)
}

func (controller *UserControllerImplementation) PasswordCodeRequest(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	request := request.ReadFromPasswordCodeRequestBody(c, requestId, controller.Logger)
	controller.UserServiceInterface.PasswordCodeRequest(requestId, request)
	response := response.Response{Code: 200, Mssg: "email sent"}
	return c.JSON(http.StatusOK, response)
}

func (controller *UserControllerImplementation) UpdateStatusActiveUser(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	accessToken := c.QueryParam("access_token")
	err := controller.UserServiceInterface.UpdateStatusActiveUser(requestId, accessToken)
	if err == nil {
		// return c.JSON(http.StatusOK, "VERIFIKASI SUCCESS SILAKAN LOGIN DI APLIAKSI TEMAN BUNDA")
		return c.File("./template/verifikasi_email_success.html")
	} else {
		return nil
	}
}

func (controller *UserControllerImplementation) FindUserByReferal(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	referal := c.QueryParam("referal")
	referalUppercase := strings.ToUpper(referal)
	userResponse := controller.UserServiceInterface.FindUserByReferal(requestId, referalUppercase)
	response := response.Response{Code: 200, Mssg: "success", Data: userResponse.ReferalCode, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}

func (controller *UserControllerImplementation) FindUserById(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	idUser := middleware.TokenClaimsIdUser(c)
	userResponse := controller.UserServiceInterface.FindUserById(requestId, idUser)
	response := response.Response{Code: 200, Mssg: "success", Data: userResponse, Error: []string{}}
	return c.JSON(http.StatusOK, response)
}
