package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type UpdateUserTokenDeviceRequest struct {
	TokenDevice string `json:"token_device" form:"token_device" validate:"required"`
}

func ReadFromUpdateUseTokenDevicerRequestBody(c echo.Context, requestId string, logger *logrus.Logger) (updateUserTokenDevice *UpdateUserTokenDeviceRequest) {
	updateUserTokenDeviceRequest := new(UpdateUserTokenDeviceRequest)
	if err := c.Bind(updateUserTokenDeviceRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	updateUserTokenDevice = updateUserTokenDeviceRequest
	return updateUserTokenDevice
}

func ValidateUpdateUserTokenDeviceRequest(validate *validator.Validate, updateUserTokenDevice *UpdateUserTokenDeviceRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	err := validate.Struct(updateUserTokenDevice)
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
