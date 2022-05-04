package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type PaymentStatusRequest struct {
	TranscationId string `json:"transactionId" form:"transactionId" validate:"required"`
	IdOrder       string `json:"id_order" form:"id_order" validate:"required"`
}

func ReadFromPaymentStatusRequestBody(c echo.Context, requestId string, logger *logrus.Logger) (paymentStatus *PaymentStatusRequest) {
	paymentStatusRequest := new(PaymentStatusRequest)
	if err := c.Bind(paymentStatusRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	paymentStatus = paymentStatusRequest
	return paymentStatus
}

func ValidatePaymentStatusRequest(validate *validator.Validate, paymentStatus *PaymentStatusRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	err := validate.Struct(paymentStatus)
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
