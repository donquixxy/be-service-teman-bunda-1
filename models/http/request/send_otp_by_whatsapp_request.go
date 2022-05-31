package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type SendOtpByWhatsappRequest struct {
	Phone string `json:"phone" form:"phone" validate:"required"`
}

func ReadFromSendOtpByWhatsappRequestBody(c echo.Context, requestId string, logger *logrus.Logger) (sendOtpByWhatsapp *SendOtpByWhatsappRequest) {
	sendOtpByWhatsappRequest := new(SendOtpByWhatsappRequest)
	if err := c.Bind(sendOtpByWhatsappRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	sendOtpByWhatsapp = sendOtpByWhatsappRequest
	return sendOtpByWhatsapp
}

func ValidateSendOtpByWhatsapRequest(validate *validator.Validate, sendOtpByWhatsappRequest *SendOtpByWhatsappRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	var errorString string
	err := validate.Struct(sendOtpByWhatsappRequest)
	if err != nil {
		for _, errorValidation := range err.(validator.ValidationErrors) {
			errorString = errorValidation.Field() + " is " + errorValidation.Tag()
			errorStrings = append(errorStrings, errorString)
		}
		// panic(exception.NewError(400, errorStrings))
		exceptions.PanicIfBadRequest(err, requestId, errorStrings, logger)
	}
}
