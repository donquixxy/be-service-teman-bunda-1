package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type AuthRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	AppVer   string `json:"app_ver" form:"app_ver"`
}

func ReadFromAuthRequestBody(c echo.Context, requestId string, logger *logrus.Logger) (authRequest *AuthRequest) {
	authRequest = new(AuthRequest)
	if err := c.Bind(authRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	return authRequest
}

func ValidateAuth(validate *validator.Validate, authRequest *AuthRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	var errorString string
	err := validate.Struct(authRequest)
	if err != nil {
		for _, errorValidation := range err.(validator.ValidationErrors) {
			errorString = errorValidation.Field() + " is " + errorValidation.Tag()
			errorStrings = append(errorStrings, errorString)
		}
		exceptions.PanicIfBadRequest(err, requestId, errorStrings, logger)
	}
}
