package services

import (
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type KelurahanServiceInterface interface {
	FindAllKelurahanByIdKecamatan(requestId string, id int) (kelurahanResponses []response.FindKelurahanByKecamatanResponse)
}

type KelurahanServiceImplementation struct {
	ConfigWebserver              config.Webserver
	DB                           *gorm.DB
	Logger                       *logrus.Logger
	KelurahanRepositoryInterface mysql.KelurahanRepositoryInterface
}

func NewKelurahanService(configWebserver config.Webserver, DB *gorm.DB, logger *logrus.Logger, kelurahanRepositoryInterface mysql.KelurahanRepositoryInterface) KelurahanServiceInterface {
	return &KelurahanServiceImplementation{
		ConfigWebserver:              configWebserver,
		DB:                           DB,
		Logger:                       logger,
		KelurahanRepositoryInterface: kelurahanRepositoryInterface,
	}
}

func (service *KelurahanServiceImplementation) FindAllKelurahanByIdKecamatan(requestId string, id int) (kelurahanResponses []response.FindKelurahanByKecamatanResponse) {
	kelurahans, err := service.KelurahanRepositoryInterface.FindAllKelurahanByIdKecamatan(service.DB, id)
	exceptions.PanicIfRecordNotFound(err, requestId, []string{"Data not found"}, service.Logger)
	exceptions.PanicIfError(err, requestId, service.Logger)
	kelurahanResponses = response.ToFindKelurahanByKecamatanResponse(kelurahans)
	return kelurahanResponses
}
