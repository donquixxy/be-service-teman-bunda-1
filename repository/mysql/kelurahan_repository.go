package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type KelurahanRepositoryInterface interface {
	FindAllKelurahanByIdKecamatan(DB *gorm.DB, id int) ([]entity.Kelurahan, error)
	FindKelurahanById(DB *gorm.DB, IdKelurahan int) (entity.Kelurahan, error)
}

type KelurahanRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewKelurahanRepository(configDatabase *config.Database) KelurahanRepositoryInterface {
	return &KelurahanRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *KelurahanRepositoryImplementation) FindAllKelurahanByIdKecamatan(DB *gorm.DB, id int) ([]entity.Kelurahan, error) {
	var kelurahans []entity.Kelurahan
	results := DB.Where("idkeca = ?", id).Find(&kelurahans)
	return kelurahans, results.Error
}

func (repository *KelurahanRepositoryImplementation) FindKelurahanById(DB *gorm.DB, IdKelurahan int) (entity.Kelurahan, error) {
	var kelurahan entity.Kelurahan
	results := DB.Where("master_kelurahan.idkelu = ?", IdKelurahan).Joins("ServiceZonaArea").Find(&kelurahan)
	return kelurahan, results.Error
}
