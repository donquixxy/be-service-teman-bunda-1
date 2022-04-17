package services

import (
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type ProvinsiServiceInterface interface {
	FindAllProvinsi(requestId string) (provinsiResponses []response.FindProvinsiAllResponse)
}

type ProvinsiServiceImplementation struct {
	ConfigWebserver             config.Webserver
	DB                          *gorm.DB
	Logger                      *logrus.Logger
	ProvinsiRepositoryInterface mysql.ProvinsiRepositoryInterface
}

func NewProvinsiService(configWebserver config.Webserver, DB *gorm.DB, logger *logrus.Logger, provinsiRepositoryInterface mysql.ProvinsiRepositoryInterface) ProvinsiServiceInterface {
	return &ProvinsiServiceImplementation{
		ConfigWebserver:             configWebserver,
		DB:                          DB,
		Logger:                      logger,
		ProvinsiRepositoryInterface: provinsiRepositoryInterface,
	}
}

func (service *ProvinsiServiceImplementation) FindAllProvinsi(requestId string) (provinsiResponses []response.FindProvinsiAllResponse) {
	provinsis, err := service.ProvinsiRepositoryInterface.FindAllProvinsi(service.DB)
	exceptions.PanicIfError(err, requestId, service.Logger)
	provinsiResponses = response.ToProvinsiFindAllResponse(provinsis)
	return provinsiResponses
}
