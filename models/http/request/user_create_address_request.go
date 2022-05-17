package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type CreateUserShippingAddressRequest struct {
	Address   string  `json:"address" form:"address" validate:"required"`
	Latitude  float64 `json:"latitude" form:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" form:"longitude" validate:"required"`
	Radius    float64 `json:"radius" form:"radius" validate:"required"`
	Note      string  `json:"note" form:"note"`
}

func ReadFromCreateUserShippingAddressRequestBody(c echo.Context, requestId string, logger *logrus.Logger) (createUserShippingAddress *CreateUserShippingAddressRequest) {
	createUserShippingAddressRequest := new(CreateUserShippingAddressRequest)
	if err := c.Bind(createUserShippingAddressRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	createUserShippingAddress = createUserShippingAddressRequest
	return createUserShippingAddress
}

func ValidateCreateUserShippingAddressRequest(validate *validator.Validate, createUserShippingAddress *CreateUserShippingAddressRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	err := validate.Struct(createUserShippingAddress)
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
