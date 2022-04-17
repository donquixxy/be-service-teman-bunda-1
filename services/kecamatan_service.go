package services

import (
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type KecamatanServiceInterface interface {
	FindAllKecamatanByIdKabupaten(requestId string, id int) (kecamatanResponses []response.FindKecamatanByKabupatenResponse)
}

type KecamatanServiceImplementation struct {
	ConfigWebserver              config.Webserver
	DB                           *gorm.DB
	Logger                       *logrus.Logger
	KecamatanRepositoryInterface mysql.KecamatanRepositoryInterface
}

func NewKecamatanService(configWebserver config.Webserver, DB *gorm.DB, logger *logrus.Logger, kecamatanRepositoryInterface mysql.KecamatanRepositoryInterface) KecamatanServiceInterface {
	return &KecamatanServiceImplementation{
		ConfigWebserver:              configWebserver,
		DB:                           DB,
		Logger:                       logger,
		KecamatanRepositoryInterface: kecamatanRepositoryInterface,
	}
}

func (service *KecamatanServiceImplementation) FindAllKecamatanByIdKabupaten(requestId string, id int) (kecamatanResponses []response.FindKecamatanByKabupatenResponse) {
	kecamatans, err := service.KecamatanRepositoryInterface.FindAllKecamatanByIdKabupaten(service.DB, id)
	exceptions.PanicIfRecordNotFound(err, requestId, []string{"Data not found"}, service.Logger)
	exceptions.PanicIfError(err, requestId, service.Logger)
	kecamatanResponses = response.ToFindKecamatanByKabuaptenResponse(kecamatans)
	return kecamatanResponses
}
