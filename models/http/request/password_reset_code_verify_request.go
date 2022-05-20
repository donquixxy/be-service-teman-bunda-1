package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type PasswordResetCodeVerifyRequest struct {
	Email string `json:"email" form:"email" validate:"required"`
	Code  string `json:"code" form:"code" validate:"required"`
}

func ReadFromPasswordResetCodeVerifyBody(c echo.Context, requestId string, logger *logrus.Logger) (passwordResetCodeVerify *PasswordResetCodeVerifyRequest) {
	passwordResetCodeVerifyRequest := new(PasswordResetCodeVerifyRequest)
	if err := c.Bind(passwordResetCodeVerifyRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	passwordResetCodeVerify = passwordResetCodeVerifyRequest
	return passwordResetCodeVerify
}

func ValidatePasswordResetCodeVerifyRequest(validate *validator.Validate, passwordResetCodeVerify *PasswordResetCodeVerifyRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	err := validate.Struct(passwordResetCodeVerify)
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
