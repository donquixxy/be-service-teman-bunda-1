package services

import (
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/request"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	modelService "github.com/tensuqiuwulu/be-service-teman-bunda/models/service"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"github.com/tensuqiuwulu/be-service-teman-bunda/utilities"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type AuthServiceInterface interface {
	Login(requestId string, authRequest *request.AuthRequest) (authResponse interface{})
	NewToken(requestId string, refreshToken string) (token string)
	GenerateToken(user modelService.User) (token string, err error)
	GenerateRefreshToken(user modelService.User) (token string, err error)
	VerifyOtp(requestId string, verifyOtpRequest *request.VerifyOtpRequest) error
	SendOtpBySms(requestId string, sendOtpBySmsRequest *request.SendOtpBySmsRequest) error
	SendOtpByEmail(requestId string, sendOtpByEmail *request.SendOtpByEmailRequest) error
}

type AuthServiceImplementation struct {
	ConfigurationWebserver     config.Webserver
	DB                         *gorm.DB
	ConfigJwt                  config.Jwt
	Validate                   *validator.Validate
	Logger                     *logrus.Logger
	UserRepositoryInterface    mysql.UserRepositoryInterface
	SettingRepositoryInterface mysql.SettingRepositoryInterface
}

func NewAuthService(
	configurationWebserver config.Webserver,
	DB *gorm.DB,
	configJwt config.Jwt,
	validate *validator.Validate,
	logger *logrus.Logger,
	userRepositoryInterface mysql.UserRepositoryInterface,
	settingRepositoryInterface mysql.SettingRepositoryInterface) AuthServiceInterface {
	return &AuthServiceImplementation{
		ConfigurationWebserver:     configurationWebserver,
		DB:                         DB,
		ConfigJwt:                  configJwt,
		Validate:                   validate,
		Logger:                     logger,
		UserRepositoryInterface:    userRepositoryInterface,
		SettingRepositoryInterface: settingRepositoryInterface,
	}
}

func (service *AuthServiceImplementation) SendOtpByEmail(requestId string, sendOtpByEmailRequest *request.SendOtpByEmailRequest) error {
	request.ValidateSendOtpByEmailRequest(service.Validate, sendOtpByEmailRequest, requestId, service.Logger)

	emailLowerCase := strings.ToLower(sendOtpByEmailRequest.Email)
	user, _ := service.UserRepositoryInterface.FindUserByEmail(service.DB, emailLowerCase)
	if user.Id == "" {
		exceptions.PanicIfRecordNotFound(errors.New("record not found"), requestId, []string{"user not found"}, service.Logger)
	}

	otpCode := utilities.GenerateRandomCode()
	bcryptOtpCode, err := bcrypt.GenerateFromPassword([]byte(otpCode), bcrypt.DefaultCost)
	exceptions.PanicIfBadRequest(err, requestId, []string{"Error Generate otp code"}, service.Logger)

	userEntity := &entity.User{}
	userEntity.OtpCode = string(bcryptOtpCode)
	userEntity.OtpCodeExpiredDueDate = null.NewTime(time.Now().Add(time.Minute*5), true)
	errUpdateOtpCodeUser := service.UserRepositoryInterface.UpdateOtpCodeUser(service.DB, user.Id, *userEntity)
	exceptions.PanicIfError(errUpdateOtpCodeUser, requestId, service.Logger)

	dataEmail := modelService.BodyCodeEmail{
		Code:     otpCode,
		FullName: user.FamilyMembers.FullName,
	}

	template := "./template/verifikasi_code_password.html"
	subject := "Kode OTP Teman Bunda"
	go utilities.SendEmail(user.FamilyMembers.Email, subject, dataEmail, template)
	return nil
}

func (service *AuthServiceImplementation) SendOtpBySms(requestId string, sendOtpBySmsRequest *request.SendOtpBySmsRequest) error {
	request.ValidateSendOtpBySmsRequest(service.Validate, sendOtpBySmsRequest, requestId, service.Logger)

	user, _ := service.UserRepositoryInterface.FindUserByPhone(service.DB, sendOtpBySmsRequest.Phone)
	if user.Id == "" {
		exceptions.PanicIfRecordNotFound(errors.New("record not found"), requestId, []string{"user not found"}, service.Logger)
	}

	// make otp code
	otpCode := utilities.GenerateRandomCode()
	bcryptOtpCode, err := bcrypt.GenerateFromPassword([]byte(otpCode), bcrypt.DefaultCost)
	exceptions.PanicIfBadRequest(err, requestId, []string{"Error Generate otp code"}, service.Logger)

	userEntity := &entity.User{}
	userEntity.OtpCode = string(bcryptOtpCode)
	userEntity.OtpCodeExpiredDueDate = null.NewTime(time.Now().Add(time.Minute*5), true)
	errUpdateOtpCodeUser := service.UserRepositoryInterface.UpdateOtpCodeUser(service.DB, user.Id, *userEntity)
	exceptions.PanicIfError(errUpdateOtpCodeUser, requestId, service.Logger)

	go utilities.SendSmsOtp(sendOtpBySmsRequest.Phone, otpCode)
	return nil
}

func (service *AuthServiceImplementation) VerifyOtp(requestId string, verifyOtpRequest *request.VerifyOtpRequest) error {
	request.ValidateVerifyOtpByPhoneRequest(service.Validate, verifyOtpRequest, requestId, service.Logger)

	var user entity.User

	user, _ = service.UserRepositoryInterface.FindUserByPhone(service.DB, verifyOtpRequest.Credential)

	if user.Id == "" {
		emailLowerCase := strings.ToLower(verifyOtpRequest.Credential)
		user, _ = service.UserRepositoryInterface.FindUserByEmail(service.DB, emailLowerCase)
		if user.Id == "" {
			exceptions.PanicIfRecordNotFound(errors.New("user not found"), requestId, []string{"user not found"}, service.Logger)
		}
	}

	// cek if otp code not exist
	if user.OtpCode == " " {
		exceptions.PanicIfBadRequest(errors.New("otp code null"), requestId, []string{"otp code null"}, service.Logger)
	}

	// cek expired token
	if time.Now().After(user.OtpCodeExpiredDueDate.Time) {
		exceptions.PanicIfBadRequest(errors.New("otp code has expired"), requestId, []string{"otp code has expired"}, service.Logger)
	}

	// verify Otp code
	err := bcrypt.CompareHashAndPassword([]byte(user.OtpCode), []byte(verifyOtpRequest.OtpCode))
	exceptions.PanicIfBadRequest(err, requestId, []string{"Invalid Credentials"}, service.Logger)

	if user.IsActive == 0 && user.NotVerification != 1 {
		userEntity := &entity.User{}
		userEntity.OtpCode = " "
		userEntity.IsActive = 1
		userEntity.VerificationDate = null.NewTime(time.Now(), true)
		_, err := service.UserRepositoryInterface.UpdateStatusActiveUser(service.DB, user.Id, *userEntity)
		exceptions.PanicIfError(err, requestId, service.Logger)
		return nil
	} else if user.IsActive == 1 {
		userEntity := &entity.User{}
		userEntity.OtpCode = " "
		errUpdateOtpCodeUser := service.UserRepositoryInterface.UpdateOtpCodeUser(service.DB, user.Id, *userEntity)
		exceptions.PanicIfError(errUpdateOtpCodeUser, requestId, service.Logger)
		return nil
	} else {
		err := errors.New("bad request")
		exceptions.PanicIfBadRequest(err, requestId, []string{"bad request"}, service.Logger)
		return err
	}
}

func (service *AuthServiceImplementation) Login(requestId string, authRequest *request.AuthRequest) (authResponse interface{}) {
	var userModelService modelService.User
	var user entity.User

	request.ValidateAuth(service.Validate, authRequest, requestId, service.Logger)

	// jika username tidak ditemukan
	user, _ = service.UserRepositoryInterface.FindUserByPhone(service.DB, authRequest.Credential)
	if user.Id == "" {
		// cek apakah yg di input email
		emailLowerCase := strings.ToLower(authRequest.Credential)
		user, _ = service.UserRepositoryInterface.FindUserByEmail(service.DB, emailLowerCase)
		if user.Id == "" {
			user, _ = service.UserRepositoryInterface.FindUserByUsername(service.DB, authRequest.Credential)
			if len(user.Id) == 0 {
				exceptions.PanicIfRecordNotFound(errors.New("user not found"), requestId, []string{"not found"}, service.Logger)
			}
		}
	}

	if user.IsActive == 1 {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authRequest.Password))
		exceptions.PanicIfBadRequest(err, requestId, []string{"Invalid Credentials"}, service.Logger)

		userModelService.Id = user.Id
		userModelService.Username = user.Username
		userModelService.IdKelurahan = user.FamilyMembers.IdKelurahan

		token, err := service.GenerateToken(userModelService)
		exceptions.PanicIfError(err, requestId, service.Logger)

		refreshToken, err := service.GenerateRefreshToken(userModelService)
		exceptions.PanicIfError(err, requestId, service.Logger)

		_, err = service.UserRepositoryInterface.SaveUserRefreshToken(service.DB, userModelService.Id, refreshToken)
		exceptions.PanicIfError(err, requestId, service.Logger)

		setting, _ := service.SettingRepositoryInterface.FindSettingsByName(service.DB, "ver_app")

		authResponse = response.ToAuthResponse(userModelService.Id, userModelService.Username, token, refreshToken, setting.SettingsTitle)

		return authResponse
	} else {
		exceptions.PanicIfUnauthorized(errors.New("account is not active"), requestId, []string{"not active"}, service.Logger)
		return nil
	}

}

func (service *AuthServiceImplementation) NewToken(requestId string, refreshToken string) (token string) {
	tokenParse, err := jwt.ParseWithClaims(refreshToken, &modelService.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.ConfigJwt.Key), nil
	})

	if !tokenParse.Valid {
		exceptions.PanicIfUnauthorized(err, requestId, []string{"invalid token"}, service.Logger)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			exceptions.PanicIfUnauthorized(err, requestId, []string{"invalid token"}, service.Logger)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			exceptions.PanicIfUnauthorized(err, requestId, []string{"expired token"}, service.Logger)
		} else {
			exceptions.PanicIfError(err, requestId, service.Logger)
		}
	}

	if claims, ok := tokenParse.Claims.(*modelService.TokenClaims); ok && tokenParse.Valid {
		//fmt.Printf("%v %v", claims, ok)
		user, err := service.UserRepositoryInterface.FindUserByUsernameAndRefreshToken(service.DB, claims.Username, refreshToken)
		exceptions.PanicIfRecordNotFound(err, requestId, []string{"User tidak ada"}, service.Logger)

		var userModelService modelService.User
		userModelService.Id = user.Id
		userModelService.Username = user.Username
		// userModelService.CreatedDate = user.CreatedDate
		token, err := service.GenerateRefreshToken(userModelService)
		exceptions.PanicIfError(err, requestId, service.Logger)
		return token
	} else {
		err := errors.New("no claims")
		exceptions.PanicIfBadRequest(err, requestId, []string{"no claims"}, service.Logger)
		return ""
	}
}

func (service *AuthServiceImplementation) GenerateToken(user modelService.User) (token string, err error) {
	// Create the Claims
	claims := modelService.TokenClaims{
		Id:          user.Id,
		Username:    user.Username,
		IdKelurahan: user.IdKelurahan,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(service.ConfigJwt.Tokenexpiredtime)).Unix(),
			Issuer:    "aether",
		},
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenWithClaims.SignedString([]byte(service.ConfigJwt.Key))
	if err != nil {
		return "", err
	}
	return token, err
}

func (service *AuthServiceImplementation) GenerateRefreshToken(user modelService.User) (token string, err error) {
	// Create the Claims
	claims := modelService.TokenClaims{
		Id:       user.Id,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, int(service.ConfigJwt.Refreshtokenexpiredtime)).Unix(),
			Issuer:    "aether",
		},
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenWithClaims.SignedString([]byte(service.ConfigJwt.Key))
	if err != nil {
		return "", err
	}
	return token, err
}
