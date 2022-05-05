package services

import (
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"
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
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	CreateUser(requestId string, userRequest *request.CreateUserRequest) (userResponse response.CreateUserResponse)
	FindUserByReferal(requestId string, referalCode string) (userResponse response.FindUserByReferalResponse)
	FindUserById(requestId string, id string) (userResponse response.FindUserByIdResponse)
	UpdateUser(requestId string, idUser string, userRequest *request.UpdateUserRequest) error
	UpdateStatusActiveUser(requestId string, accessToken string) error
	PasswordCodeRequest(requestId string, passwordRequest *request.PasswordCodeRequest) error
	PasswordResetCodeVerify(requestId string, passwordResetCodeVerifyRequest *request.PasswordResetCodeVerifyRequest) error
	UpdateUserPassword(requestId string, updateUserPasswordRequest *request.UpdateUserPasswordRequest) error
}

type UserServiceImplementation struct {
	ConfigurationWebserver            config.Webserver
	DB                                *gorm.DB
	ConfigJwt                         config.Jwt
	Validate                          *validator.Validate
	Logger                            *logrus.Logger
	ConfigEmail                       config.Email
	UserRepositoryInterface           mysql.UserRepositoryInterface
	ProvinsiRepositoryInterface       mysql.ProvinsiRepositoryInterface
	FamilyRepositoryInterface         mysql.FamilyRepositoryInterface
	FamilyMembersRepositoryInterface  mysql.FamilyMembersRepositoryInterface
	BalancePointRepositoryInterface   mysql.BalancePointRepositoryInterface
	BalancePointTxRepositoryInterface mysql.BalancePointTxRepositoryInterface
}

func NewUserService(
	configurationWebserver config.Webserver,
	DB *gorm.DB, configJwt config.Jwt,
	validate *validator.Validate,
	logger *logrus.Logger,
	configEmail config.Email,
	userRepositoryInterface mysql.UserRepositoryInterface,
	provinsiRepositoryInterface mysql.ProvinsiRepositoryInterface,
	familyRepositoryInterface mysql.FamilyRepositoryInterface,
	familyMembersRepositoryInterface mysql.FamilyMembersRepositoryInterface,
	balancePointRepositoryInterface mysql.BalancePointRepositoryInterface,
	balancePointTxRepositoryInterface mysql.BalancePointTxRepositoryInterface) UserServiceInterface {
	return &UserServiceImplementation{
		ConfigurationWebserver:            configurationWebserver,
		DB:                                DB,
		ConfigJwt:                         configJwt,
		Validate:                          validate,
		Logger:                            logger,
		ConfigEmail:                       configEmail,
		UserRepositoryInterface:           userRepositoryInterface,
		ProvinsiRepositoryInterface:       provinsiRepositoryInterface,
		FamilyRepositoryInterface:         familyRepositoryInterface,
		FamilyMembersRepositoryInterface:  familyMembersRepositoryInterface,
		BalancePointRepositoryInterface:   balancePointRepositoryInterface,
		BalancePointTxRepositoryInterface: balancePointTxRepositoryInterface,
	}
}

func (service *UserServiceImplementation) PasswordResetCodeVerify(requestId string, passwordResetCodeVerifyRequest *request.PasswordResetCodeVerifyRequest) error {
	// Validate request
	request.ValidatePasswordResetCodeVerifyRequest(service.Validate, passwordResetCodeVerifyRequest, requestId, service.Logger)

	user, _ := service.UserRepositoryInterface.FindUserByEmail(service.DB, passwordResetCodeVerifyRequest.Email)

	if user.PasswordResetCode == passwordResetCodeVerifyRequest.Code {
		return nil
	} else {
		err := errors.New("email and code not match")
		exceptions.PanicIfBadRequest(err, requestId, []string{"email and code not match"}, service.Logger)
		return err
	}
}

func (service *UserServiceImplementation) PasswordCodeRequest(requestId string, passwordRequest *request.PasswordCodeRequest) error {
	user, _ := service.UserRepositoryInterface.FindUserByEmail(service.DB, passwordRequest.Email)

	if user.Id == "" {
		exceptions.PanicIfRecordNotFound(errors.New("email not found"), requestId, []string{"Email not registered"}, service.Logger)
	}

	rand.Seed(time.Now().Unix())
	charSet := "1234567890"
	var output strings.Builder
	length := 6

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}

	userEntity := &entity.User{}
	userEntity.PasswordResetCode = output.String()

	_, errUpdateUser := service.UserRepositoryInterface.UpdatePasswordResetCodeUser(service.DB, user.Id, *userEntity)
	exceptions.PanicIfError(errUpdateUser, requestId, service.Logger)

	service.SendEmailPasswordResetCode(user.FamilyMembers.Email, userEntity.PasswordResetCode)

	return nil
}

func (service *UserServiceImplementation) UpdateStatusActiveUser(requestId string, accessToken string) error {
	tokenParse, err := jwt.ParseWithClaims(accessToken, &modelService.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.ConfigJwt.VerifyKey), nil
	})

	if !tokenParse.Valid {
		exceptions.PanicIfUnauthorized(err, requestId, []string{"invalid token"}, service.Logger)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			exceptions.PanicIfUnauthorized(err, requestId, []string{"invalid token"}, service.Logger)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			exceptions.PanicIfUnauthorized(err, requestId, []string{"invalid token"}, service.Logger)
		} else {
			exceptions.PanicIfUnauthorized(err, requestId, []string{"invalid token"}, service.Logger)
		}
	}

	if claims, ok := tokenParse.Claims.(*modelService.TokenClaims); ok && tokenParse.Valid {
		user, _ := service.UserRepositoryInterface.FindUserById(service.DB, claims.Id)
		userEntity := &entity.User{}
		userEntity.IsActive = 1

		_, errUpdateUser := service.UserRepositoryInterface.UpdateStatusActiveUser(service.DB, user.Id, *userEntity)
		exceptions.PanicIfError(errUpdateUser, requestId, service.Logger)
		return nil
	} else {
		err := errors.New("no claims")
		exceptions.PanicIfBadRequest(err, requestId, []string{"no claims"}, service.Logger)
		return nil
	}
}

func (service *UserServiceImplementation) UpdateUserPassword(requestId string, updateUserPasswordRequest *request.UpdateUserPasswordRequest) error {
	// Validate request
	request.ValidateUpdateUserPasswordRequest(service.Validate, updateUserPasswordRequest, requestId, service.Logger)

	user, _ := service.UserRepositoryInterface.FindUserByEmail(service.DB, updateUserPasswordRequest.Email)

	password := updateUserPasswordRequest.Password
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	exceptions.PanicIfBadRequest(err, requestId, []string{"Error Generate Password"}, service.Logger)

	if user.PasswordResetCode == updateUserPasswordRequest.Code {
		userEntity := &entity.User{}
		userEntity.Password = string(bcryptPassword)
		userEntity.PasswordResetCode = " "
		_, errUpdateUser := service.UserRepositoryInterface.UpdateUserPassword(service.DB, user.Id, *userEntity)
		exceptions.PanicIfError(errUpdateUser, requestId, service.Logger)
	} else {
		err := errors.New("email and code not match")
		exceptions.PanicIfBadRequest(err, requestId, []string{"email and code not match"}, service.Logger)
	}

	return nil
}

func (service *UserServiceImplementation) UpdateUser(requestId string, idUser string, userRequest *request.UpdateUserRequest) error {

	// Validate request
	request.ValidateUpdateUserRequest(service.Validate, userRequest, requestId, service.Logger)

	user, _ := service.UserRepositoryInterface.FindUserById(service.DB, idUser)

	if userRequest.Username != user.Username {
		// Check username if exsict
		checkUsername, _ := service.UserRepositoryInterface.FindUserByUsername(service.DB, userRequest.Username)
		if checkUsername.Id != "" {
			err := errors.New("username already exist")
			exceptions.PanicIfRecordAlreadyExists(err, requestId, []string{"Username already exist"}, service.Logger)
		}
	}

	if userRequest.Email != user.FamilyMembers.Email {
		// Check email if exsict
		checkEmail, _ := service.UserRepositoryInterface.FindUserByEmail(service.DB, userRequest.Email)
		if checkEmail.Id != "" {
			err := errors.New("email already exist")
			exceptions.PanicIfRecordAlreadyExists(err, requestId, []string{"Email already exist"}, service.Logger)
		}
	}

	if userRequest.Phone != user.FamilyMembers.Phone {
		// Check phone if exsict
		checkPhone, _ := service.UserRepositoryInterface.FindUserByPhone(service.DB, userRequest.Phone)
		if checkPhone.Id != "" {
			err := errors.New("phone already exist")
			exceptions.PanicIfRecordAlreadyExists(err, requestId, []string{"Phone already exist"}, service.Logger)
		}
	}

	tx := service.DB.Begin()

	// Generate Password
	password := userRequest.Password
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	exceptions.PanicIfBadRequest(err, requestId, []string{"Error Generate Password"}, service.Logger)

	// Create family members profile
	familyMembersEntity := &entity.FamilyMembers{}
	familyMembersEntity.FullName = userRequest.FullName
	familyMembersEntity.Email = userRequest.Email
	familyMembersEntity.Address = userRequest.Address
	familyMembersEntity.Phone = userRequest.Phone
	familyMembersEntity.IdProvinsi = userRequest.IdProvinsi
	familyMembersEntity.IdKabupaten = userRequest.IdKabupaten
	familyMembersEntity.IdKecamatan = userRequest.IdKecamatan
	familyMembersEntity.IdKelurahan = userRequest.IdKelurahan
	familyMembers, errUpdateFamilyMembers := service.FamilyMembersRepositoryInterface.UpdateFamilyMembers(tx, user.IdFamilyMembers, *familyMembersEntity)
	exceptions.PanicIfErrorWithRollback(errUpdateFamilyMembers, requestId, []string{"Error update family members"}, service.Logger, tx)

	// Crate user profile
	userEntity := &entity.User{}
	userEntity.IdFamilyMembers = familyMembers.Id
	userEntity.Username = userRequest.Username
	if userRequest.Password != "" {
		userEntity.Password = string(bcryptPassword)
	}
	_, errUpdateUser := service.UserRepositoryInterface.UpdateUser(tx, idUser, *userEntity)
	exceptions.PanicIfErrorWithRollback(errUpdateUser, requestId, []string{"Error update user"}, service.Logger, tx)

	commit := tx.Commit()
	exceptions.PanicIfError(commit.Error, requestId, service.Logger)
	return nil
}

func (service *UserServiceImplementation) CreateUser(requestId string, userRequest *request.CreateUserRequest) (userResponse response.CreateUserResponse) {

	// Validate request
	request.ValidateCreateUserRequest(service.Validate, userRequest, requestId, service.Logger)

	// Check username if exsict
	checkUsername, _ := service.UserRepositoryInterface.FindUserByUsername(service.DB, userRequest.Username)
	if checkUsername.Id != "" {
		err := errors.New("username already exist")
		exceptions.PanicIfRecordAlreadyExists(err, requestId, []string{"Username sudah digunakan"}, service.Logger)
	}

	// Check email if exsict
	checkEmail, _ := service.UserRepositoryInterface.FindUserByEmail(service.DB, userRequest.Email)
	if checkEmail.Id != "" {
		err := errors.New("email already exist")
		exceptions.PanicIfRecordAlreadyExists(err, requestId, []string{"Email sudah digunakan"}, service.Logger)
	}

	// Check phone if exsict
	checkPhone, _ := service.UserRepositoryInterface.FindUserByPhone(service.DB, userRequest.Phone)
	if checkPhone.Id != "" {
		err := errors.New("phone already exist")
		exceptions.PanicIfRecordAlreadyExists(err, requestId, []string{"Phone sudah digunakan"}, service.Logger)
	}

	// Begin Transcation
	tx := service.DB.Begin()
	exceptions.PanicIfError(tx.Error, requestId, service.Logger)

	// Generate Password
	password := userRequest.Password
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	exceptions.PanicIfBadRequest(err, requestId, []string{"Error Generate Password"}, service.Logger)

	// Generate referal code
	referalCode := service.GenerateReferalCode(userRequest.IdProvinsi)

	// Create family profile
	familyEntity := &entity.Family{}
	familyEntity.Id = utilities.RandomUUID()
	familyEntity.IdProvinsi = userRequest.IdProvinsi
	familyEntity.IdKabupaten = userRequest.IdKabupaten
	familyEntity.IdKecamatan = userRequest.IdKecamatan
	familyEntity.IdKelurahan = userRequest.IdKelurahan
	family, err := service.FamilyRepositoryInterface.CreateFamily(tx, *familyEntity)
	exceptions.PanicIfErrorWithRollback(err, requestId, []string{"Error create family"}, service.Logger, tx)

	// Create family members profile
	familyMembersEntity := &entity.FamilyMembers{}
	familyMembersEntity.Id = utilities.RandomUUID()
	familyMembersEntity.IdFamily = familyEntity.Id
	familyMembersEntity.FullName = userRequest.FullName
	familyMembersEntity.Email = userRequest.Email
	familyMembersEntity.Address = userRequest.Address
	familyMembersEntity.Phone = userRequest.Phone
	familyMembersEntity.IdProvinsi = userRequest.IdProvinsi
	familyMembersEntity.IdKabupaten = userRequest.IdKabupaten
	familyMembersEntity.IdKecamatan = userRequest.IdKecamatan
	familyMembersEntity.IdKelurahan = userRequest.IdKelurahan
	familyMembers, err := service.FamilyMembersRepositoryInterface.CreateFamilyMembers(tx, *familyMembersEntity)
	exceptions.PanicIfErrorWithRollback(err, requestId, []string{"Error create family members"}, service.Logger, tx)

	// Crate user profile
	userEntity := &entity.User{}
	userEntity.Id = utilities.RandomUUID()
	userEntity.IdFamilyMembers = familyMembers.Id
	userEntity.IdLevelMember = 1
	userEntity.Username = userRequest.Username
	userEntity.Password = string(bcryptPassword)
	if userRequest.RegistrationReferalCode == "" {
		userEntity.RegistrationReferalCode = "0X0ROQIBA"
	} else {
		userEntity.RegistrationReferalCode = userRequest.RegistrationReferalCode
	}

	userEntity.CreatedDate = time.Now()
	userEntity.ReferalCode = referalCode
	userEntity.RefreshToken = ""
	user, err := service.UserRepositoryInterface.CreateUser(tx, *userEntity)
	exceptions.PanicIfErrorWithRollback(err, requestId, []string{"Error insert user"}, service.Logger, tx)

	// Create user balance points
	balancePointEntity := &entity.BalancePoint{}
	balancePointEntity.Id = utilities.RandomUUID()
	balancePointEntity.IdUser = userEntity.Id
	balancePointEntity.CreatedDate = time.Now()
	balancePoint, err := service.BalancePointRepositoryInterface.CreateBalancePoint(tx, *balancePointEntity)
	exceptions.PanicIfErrorWithRollback(err, requestId, []string{"Error insert balance point"}, service.Logger, tx)

	var userModelService modelService.User
	userModelService.Id = user.Id
	userModelService.Username = user.Username
	userModelService.IdKelurahan = user.FamilyMembers.IdKelurahan

	token, err := service.GenerateTokenVerify(userModelService)
	exceptions.PanicIfError(err, requestId, service.Logger)

	service.SendEmail(userRequest.Email, userEntity.Id, token)

	commit := tx.Commit()
	exceptions.PanicIfError(commit.Error, requestId, service.Logger)
	userResponse = response.ToUserCreateUserResponse(user, family, familyMembers, balancePoint)

	return userResponse
}

func (service *UserServiceImplementation) GenerateReferalCode(idProvinsi int) (referalCode string) {
	referalCodeEntity := &entity.ReferalCode{}
	provinsi, _ := service.ProvinsiRepositoryInterface.FindProvinsiById(service.DB, idProvinsi)
	for {
		rand.Seed(time.Now().Unix())
		charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		var output strings.Builder
		length := 7

		for i := 0; i < length; i++ {
			random := rand.Intn(len(charSet))
			randomChar := charSet[random]
			output.WriteString(string(randomChar))
		}

		referalCodeEntity.ReferalCode = output.String() + provinsi.KodeArea

		// Check referal code if exist
		checkUser, _ := service.UserRepositoryInterface.FindUserByReferal(service.DB, referalCodeEntity.ReferalCode)
		if checkUser.Id == "" {
			break
		}
	}
	return referalCodeEntity.ReferalCode
}

func (service *UserServiceImplementation) FindUserByReferal(requestId string, referal string) (userResponse response.FindUserByReferalResponse) {
	user, err := service.UserRepositoryInterface.FindUserByReferal(service.DB, referal)
	exceptions.PanicIfRecordNotFound(err, requestId, []string{"Data Not Found"}, service.Logger)
	userResponse = response.ToUserFindByReferalResponse(user)
	return userResponse
}

func (service *UserServiceImplementation) FindUserById(requestId string, id string) (userResponse response.FindUserByIdResponse) {
	user, _ := service.UserRepositoryInterface.FindUserById(service.DB, id)
	if user.Id == "" {
		err := errors.New("user not found")
		exceptions.PanicIfRecordNotFound(err, requestId, []string{"Not Found"}, service.Logger)
	}
	userCount, _ := service.UserRepositoryInterface.CountUserByRegistrationReferal(service.DB, user.ReferalCode)
	userResponse = response.ToUserFindByIdResponse(user, userCount)
	return userResponse
}

func (service *UserServiceImplementation) SendEmail(toEmail string, idUser string, token string) error {
	fromEmail := string(service.ConfigEmail.FromEmail)
	fromPasswordEmail := string(service.ConfigEmail.FromEmailPassword)

	to := []string{toEmail}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: Email Verification\r\n\r\n"
	body := "Silakan klik link berikut untuk melakukan verifikasi \n" +
		"Link : " + service.ConfigEmail.LinkVerifyEmail + token
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", fromEmail, fromPasswordEmail, host)
	fmt.Println("message : ", string(message))
	err := smtp.SendMail(address, auth, fromEmail, to, message)
	fmt.Println(err)
	return err
}

func (service *UserServiceImplementation) SendEmailPasswordResetCode(toEmail string, code string) error {
	fromEmail := string(service.ConfigEmail.FromEmail)
	fromPasswordEmail := string(service.ConfigEmail.FromEmailPassword)

	to := []string{toEmail}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: Reset Password Code\r\n\r\n"
	body := "Berikut merupakan code untuk reset password \n" +
		"CODE : " + code
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", fromEmail, fromPasswordEmail, host)
	fmt.Println("message : ", string(message))
	err := smtp.SendMail(address, auth, fromEmail, to, message)
	fmt.Println(err)
	return err
}
