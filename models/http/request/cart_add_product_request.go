package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type AddProductToCartRequest struct {
	IdProduct string `json:"full_name" form:"id_product" validate:"required"`
	Qty       int    `json:"email" form:"qty" validate:"required"`
}

func ReadFromAddProductToCartRequestBody(c echo.Context, requestId string, logger *logrus.Logger) (addProductToCart *AddProductToCartRequest) {
	addProductToCartRequest := new(AddProductToCartRequest)
	if err := c.Bind(addProductToCartRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	addProductToCart = addProductToCartRequest
	return addProductToCart
}

func ValidateAddProductToCartRequest(validate *validator.Validate, addProductToCart *AddProductToCartRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	err := validate.Struct(addProductToCart)
	if err != nil {
		for _, errorValidation := range err.(validator.ValidationErrors) {
			var errorString string
			errorString = errorValidation.Field() + " is " + errorValidation.Tag()
			errorStrings = append(errorStrings, errorString)
		}
		// panic(exception.NewError(400, errorStrings))
		exceptions.PanicIfBadRequest(err, requestId, errorStrings, logger)
	}
}
