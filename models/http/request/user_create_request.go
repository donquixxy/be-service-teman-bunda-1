package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)

type CreateUserRequest struct {
	FullName                string `json:"full_name" form:"username" validate:"required"`
	Email                   string `json:"email" form:"email" validate:"required"`
	Address                 string `json:"address" form:"address" validate:"required"`
	Phone                   string `json:"phone" form:"phone" validate:"required"`
	Username                string `json:"username" form:"username" validate:"required"`
	Password                string `json:"password" form:"password" validate:"required"`
	IdProvinsi              int    `json:"id_provinsi" form:"id_provinsi" validate:"required"`
	IdKabupaten             int    `json:"id_kabupaten" form:"id_kabupaten" validate:"required"`
	IdKecamatan             int    `json:"id_kecamatan" form:"id_kecamatan" validate:"required"`
	IdKelurahan             int    `json:"id_kelurahan" form:"id_kelurahan" validate:"required"`
	RegistrationReferalCode string `json:"registration_referal_code" form:"registration_referal_code"`
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
