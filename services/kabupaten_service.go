package services

import (
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type KabupatenServiceInterface interface {
	FindAllKabupatenByIdProvinsi(requestId string, id int) (kabupatenResponses []response.FindKabupatenByProvinsiResponse)
}

type KabupatenServiceImplementation struct {
	ConfigWebserver              config.Webserver
	DB                           *gorm.DB
	Logger                       *logrus.Logger
	KabupatenRepositoryInterface mysql.KabupatenRepositoryInterface
}

func NewKabupatenService(configWebserver config.Webserver, DB *gorm.DB, logger *logrus.Logger, kabupatenRepositoryInterface mysql.KabupatenRepositoryInterface) KabupatenServiceInterface {
	return &KabupatenServiceImplementation{
		ConfigWebserver:              configWebserver,
		DB:                           DB,
		Logger:                       logger,
		KabupatenRepositoryInterface: kabupatenRepositoryInterface,
	}
}

func (service *KabupatenServiceImplementation) FindAllKabupatenByIdProvinsi(requestId string, id int) (kabupatenResponses []response.FindKabupatenByProvinsiResponse) {
	kabupatens, err := service.KabupatenRepositoryInterface.FindAllKabupatenByIdProvinsi(service.DB, id)
	exceptions.PanicIfRecordNotFound(err, requestId, []string{"Data not found"}, service.Logger)
	exceptions.PanicIfError(err, requestId, service.Logger)
	kabupatenResponses = response.ToFindKabupatenByProvinsiResponse(kabupatens)
	return kabupatenResponses
}
