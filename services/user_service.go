package services

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"

	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/request"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"github.com/tensuqiuwulu/be-service-teman-bunda/utilities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	CreateUser(requestId string, userRequest *request.CreateUserRequest) (userResponse response.CreateUserResponse)
	FindUserByReferal(requestId string, referalCode string) (userResponse response.FindUserByReferalResponse)
	FindUserById(requestId string, id string) (userResponse response.FindUserByIdResponse)
}

type UserServiceImplementation struct {
	ConfigurationWebserver            config.Webserver
	DB                                *gorm.DB
	ConfigJwt                         config.Jwt
	Validate                          *validator.Validate
	Logger                            *logrus.Logger
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
		UserRepositoryInterface:           userRepositoryInterface,
		ProvinsiRepositoryInterface:       provinsiRepositoryInterface,
		FamilyRepositoryInterface:         familyRepositoryInterface,
		FamilyMembersRepositoryInterface:  familyMembersRepositoryInterface,
		BalancePointRepositoryInterface:   balancePointRepositoryInterface,
		BalancePointTxRepositoryInterface: balancePointTxRepositoryInterface,
	}
}

func (service *UserServiceImplementation) CreateUser(requestId string, userRequest *request.CreateUserRequest) (userResponse response.CreateUserResponse) {

	// Validate request
	request.ValidateCreateUserRequest(service.Validate, userRequest, requestId, service.Logger)

	// Check username if exsict
	checkUsername, _ := service.UserRepositoryInterface.FindUserByUsername(service.DB, userRequest.Username)
	if checkUsername.Id != "" {
		err := errors.New("username already exist")
		exceptions.PanicIfRecordAlreadyExists(err, requestId, []string{"Username already exist"}, service.Logger)
	}

	// Check email if exsict
	checkEmail, _ := service.UserRepositoryInterface.FindUserByEmail(service.DB, userRequest.Email)
	if checkEmail.Id != "" {
		err := errors.New("email already exist")
		exceptions.PanicIfRecordAlreadyExists(err, requestId, []string{"Email already exist"}, service.Logger)
	}

	// Check phone if exsict
	checkPhone, _ := service.UserRepositoryInterface.FindUserByPhone(service.DB, userRequest.Phone)
	if checkPhone.Id != "" {
		err := errors.New("phone already exist")
		exceptions.PanicIfRecordAlreadyExists(err, requestId, []string{"Phone already exist"}, service.Logger)
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
	userEntity.RegistrationReferalCode = userRequest.RegistrationReferalCode
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
