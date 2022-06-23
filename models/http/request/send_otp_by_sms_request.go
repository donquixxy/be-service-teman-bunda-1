package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type SendOtpBySmsRequest struct {
	Phone string `json:"phone" form:"phone" validate:"required"`
}

func ReadFromSendOtpBySmsRequestBody(c echo.Context, requestId string, logger *logrus.Logger) (sendOtpBySms *SendOtpBySmsRequest) {
	sendOtpBySmsRequest := new(SendOtpBySmsRequest)
	if err := c.Bind(sendOtpBySmsRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	sendOtpBySms = sendOtpBySmsRequest
	return sendOtpBySms
}

func ValidateSendOtpBySmsRequest(validate *validator.Validate, sendOtpBySmsRequest *SendOtpBySmsRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	var errorString string
	err := validate.Struct(sendOtpBySmsRequest)
	if err != nil {
		for _, errorValidation := range err.(validator.ValidationErrors) {
			errorString = errorValidation.Field() + " is " + errorValidation.Tag()
			errorStrings = append(errorStrings, errorString)
		}
		// panic(exception.NewError(400, errorStrings))
		exceptions.PanicIfBadRequest(err, requestId, errorStrings, logger)
	}
}
