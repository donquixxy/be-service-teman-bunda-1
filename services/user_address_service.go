package services

import (
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/request"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"github.com/tensuqiuwulu/be-service-teman-bunda/utilities"
	"gorm.io/gorm"
)

type UserShippingAddressServiceInterface interface {
	FindUserShippingAddressByIdUser(requestId string, idUser string) (userAShippingddressResponses []response.FindUserShippingAddress)
	CreateUserShippingAddress(requestId string, idUser string, userShippingAddressRequest *request.CreateUserShippingAddressRequest) error
	DeleteUserShippingAddress(requestId string, idUserShippingAddress string) error
}

type UserShippingAddressServiceImplementation struct {
	ConfigWebserver                        config.Webserver
	DB                                     *gorm.DB
	Validate                               *validator.Validate
	Logger                                 *logrus.Logger
	UserShippingAddressRepositoryInterface mysql.UserShippingAddressRepositoryInterface
}

func NewUserShippingAddressService(configWebserver config.Webserver, DB *gorm.DB, validate *validator.Validate, logger *logrus.Logger, userShippingAddressRepositoryInterface mysql.UserShippingAddressRepositoryInterface) UserShippingAddressServiceInterface {
	return &UserShippingAddressServiceImplementation{
		ConfigWebserver:                        configWebserver,
		DB:                                     DB,
		Validate:                               validate,
		Logger:                                 logger,
		UserShippingAddressRepositoryInterface: userShippingAddressRepositoryInterface,
	}
}

func (service *UserShippingAddressServiceImplementation) DeleteUserShippingAddress(requestId string, idUserShippingAddress string) error {
	err := service.UserShippingAddressRepositoryInterface.DeleteUserShippingAddress(service.DB, idUserShippingAddress)
	exceptions.PanicIfError(err, requestId, service.Logger)
	return err
}

func (service *UserShippingAddressServiceImplementation) CreateUserShippingAddress(requestId string, idUser string, createUserShippingAddressRequest *request.CreateUserShippingAddressRequest) error {
	// validate request
	request.ValidateCreateUserShippingAddressRequest(service.Validate, createUserShippingAddressRequest, requestId, service.Logger)

	userShippingAddressEntity := &entity.UserShippingAddress{}
	userShippingAddressEntity.Id = utilities.RandomUUID()
	userShippingAddressEntity.IdUser = idUser
	userShippingAddressEntity.Status = 0
	userShippingAddressEntity.Address = createUserShippingAddressRequest.Address
	userShippingAddressEntity.Latitude = createUserShippingAddressRequest.Latitude
	userShippingAddressEntity.Longitude = createUserShippingAddressRequest.Longitude
	userShippingAddressEntity.Radius = createUserShippingAddressRequest.Radius
	_, err := service.UserShippingAddressRepositoryInterface.CreateUserShippingAddress(service.DB, *userShippingAddressEntity)
	exceptions.PanicIfError(err, requestId, service.Logger)
	return err
}

func (service *UserShippingAddressServiceImplementation) FindUserShippingAddressByIdUser(requestId string, idUser string) (userShippingAddressResponses []response.FindUserShippingAddress) {
	userShippingAddresss, err := service.UserShippingAddressRepositoryInterface.FindUserShippingAddressByIdUser(service.DB, idUser)
	exceptions.PanicIfError(err, requestId, service.Logger)
	userShippingAddressResponses = response.ToFindUserShippingAddressResponse(userShippingAddresss)
	return userShippingAddressResponses
}
