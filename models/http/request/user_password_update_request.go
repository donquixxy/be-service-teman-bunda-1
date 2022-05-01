package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type UpdateUserPasswordRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Code     string `json:"code" form:"code" validate:"required"`
}

func ReadFromUpdateUserPasswordRequestBody(c echo.Context, requestId string, logger *logrus.Logger) (updateUserPassword *UpdateUserPasswordRequest) {
	updateUserPasswordRequest := new(UpdateUserPasswordRequest)
	if err := c.Bind(updateUserPasswordRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	updateUserPassword = updateUserPasswordRequest
	return updateUserPassword
}

func ValidateUpdateUserPasswordRequest(validate *validator.Validate, updateUserPassword *UpdateUserPasswordRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	err := validate.Struct(updateUserPassword)
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
