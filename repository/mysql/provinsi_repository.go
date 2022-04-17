package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type ProvinsiRepositoryInterface interface {
	FindAllProvinsi(DB *gorm.DB) ([]entity.Provinsi, error)
	FindProvinsiById(DB *gorm.DB, id int) (entity.Provinsi, error)
}

type ProvinsiRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewProvinsiRepository(configDatabase *config.Database) ProvinsiRepositoryInterface {
	return &ProvinsiRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *ProvinsiRepositoryImplementation) FindAllProvinsi(DB *gorm.DB) ([]entity.Provinsi, error) {
	var provinsis []entity.Provinsi
	results := DB.Find(&provinsis)
	return provinsis, results.Error
}

func (repository *ProvinsiRepositoryImplementation) FindProvinsiById(DB *gorm.DB, id int) (entity.Provinsi, error) {
	var provinsi entity.Provinsi
	results := DB.Where("idprop = ?", id).Find(&provinsi)
	return provinsi, results.Error
}
