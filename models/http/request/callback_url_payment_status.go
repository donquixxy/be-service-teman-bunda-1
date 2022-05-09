package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type CallBackIpaymuRequest struct {
	TrxId       int    `json:"trx_id" form:"trx_id" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
	StatusCode  int    `json:"status_code" form:"status_code" validate:"required"`
	ReferenceId string `json:"reference_id" form:"reference_id" validate:"required"`
}

func ReadFromCallBackIpaymuRequest(c echo.Context, requestId string, logger *logrus.Logger) (callBackIpaymu *CallBackIpaymuRequest) {
	callBackIpaymuRequest := new(CallBackIpaymuRequest)
	if err := c.Bind(callBackIpaymuRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	callBackIpaymu = callBackIpaymuRequest
	return callBackIpaymu
}

func ValidateCallBackIpaymuRequest(validate *validator.Validate, callBackIpaymu *CallBackIpaymuRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	var errorString string
	err := validate.Struct(callBackIpaymu)
	if err != nil {
		for _, errorValidation := range err.(validator.ValidationErrors) {
			errorString = errorValidation.Field() + " is " + errorValidation.Tag()
			errorStrings = append(errorStrings, errorString)
		}
		// panic(exception.NewError(400, errorStrings))
		exceptions.PanicIfBadRequest(err, requestId, errorStrings, logger)
	}
}
