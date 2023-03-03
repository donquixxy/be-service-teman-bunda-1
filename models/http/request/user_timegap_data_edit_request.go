package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
)
type UserTimegapDataEditRequest struct {
	// FullName                string `json:"full_name" form:"full_name" validate:"required"`
	// Email                   string `json:"email" form:"email" validate:"required"`
	// Phone                   string `json:"phone" form:"phone" validate:"required"`
	// Password                string `json:"password" form:"password" validate:"required"`
	// RegistrationReferalCode string `json:"registration_referal_code" form:"registration_referal_code"`
	// FormToken               string `json:"form_token" form:"form_token"`
	TimegapData             string `json:"timegap_data" form:"timegap_data" validate:"required"` 
}
// First binding data
func ReadTimegapEditRequest(c echo.Context, requestId string, logger *logrus.Logger) (editUserTimegap *UserTimegapDataEditRequest) {
	editUserTimegapRequest := new(UserTimegapDataEditRequest)
	if err := c.Bind(editUserTimegapRequest); err != nil {
		exceptions.PanicIfError(err, requestId, logger)
	}
	editUserTimegap = editUserTimegapRequest
	return editUserTimegap
}
// Second, do validation
func ValidateEditUserTimegapRequest(validate *validator.Validate, editUserTimegap *UserTimegapDataEditRequest, requestId string, logger *logrus.Logger) {
	var errorStrings []string
	err := validate.Struct(editUserTimegap)
	var errorString string
	if err != nil {
		for _, errorValidation := range err.(validator.ValidationErrors) {
			errorString = errorValidation.Field() + " is " + errorValidation.Tag()
			errorStrings = append(errorStrings, errorString)
		}
		exceptions.PanicIfBadRequest(err, requestId, errorStrings, logger)
	}
}
