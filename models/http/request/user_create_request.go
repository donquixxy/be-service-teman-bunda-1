package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type CreateUserRequest struct {
	FullName                string `json:"full_name" form:"full_name" validate:"required"`
	Email                   string `json:"email" form:"email" validate:"required"`
	Phone                   string `json:"phone" form:"phone" validate:"required"`
	Password                string `json:"password" form:"password" validate:"required"`
	RegistrationReferalCode string `json:"registration_referal_code" form:"registration_referal_code"`
	FormToken               string `json:"form_token" form:"form_token"`
}

func ReadFromCreateUserRequestBody(c echo.Context, requestId string, logger *logrus.Logger) (createUser *CreateUserRequest) {
	createUserRequest := new(CreateUserRequest)
	if err := c.Bind(createUserRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	createUser = createUserRequest
	return createUser
}

func ValidateCreateUserRequest(validate *validator.Validate, createUser *CreateUserRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	err := validate.Struct(createUser)
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
