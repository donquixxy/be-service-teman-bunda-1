package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type PasswordCodeRequest struct {
	Email string `json:"email" form:"email" validate:"required"`
}

func ReadFromPasswordCodeRequestBody(c echo.Context, requestId string, logger *logrus.Logger) (passwordCode *PasswordCodeRequest) {
	passwordCodeRequest := new(PasswordCodeRequest)
	if err := c.Bind(passwordCodeRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	passwordCode = passwordCodeRequest
	return passwordCode
}

func ValidatePasswordCodeRequest(validate *validator.Validate, passwordCode *PasswordCodeRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	err := validate.Struct(passwordCode)
	var errorString string
	if err != nil {
		for _, errorValidation := range err.(validator.ValidationErrors) {
			errorString = errorValidation.Field() + " is " + errorValidation.Tag()
			errorStrings = append(errorStrings, errorString)
		}
		// panic(exception.NewError(400, errorStrings))
		exceptions.PanicIfBadRequest(err, requestId, errorStrings, logger)
	}
}
