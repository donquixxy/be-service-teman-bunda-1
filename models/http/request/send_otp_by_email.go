package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type SendOtpByEmailRequest struct {
	Email string `json:"email" form:"email" validate:"required"`
}

func ReadFromSendOtpByEmailRequestBody(c echo.Context, requestId string, logger *logrus.Logger) (sendOtpByEmail *SendOtpByEmailRequest) {
	sendOtpByEmailRequest := new(SendOtpByEmailRequest)
	if err := c.Bind(sendOtpByEmailRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	sendOtpByEmail = sendOtpByEmailRequest
	return sendOtpByEmail
}

func ValidateSendOtpByEmailRequest(validate *validator.Validate, sendOtpByEmailRequest *SendOtpByEmailRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	var errorString string
	err := validate.Struct(sendOtpByEmailRequest)
	if err != nil {
		for _, errorValidation := range err.(validator.ValidationErrors) {
			errorString = errorValidation.Field() + " is " + errorValidation.Tag()
			errorStrings = append(errorStrings, errorString)
		}
		// panic(exception.NewError(400, errorStrings))
		exceptions.PanicIfBadRequest(err, requestId, errorStrings, logger)
	}
}
