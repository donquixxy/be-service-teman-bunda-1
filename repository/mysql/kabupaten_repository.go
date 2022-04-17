package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type KabupatenRepositoryInterface interface {
	FindAllKabupatenByIdProvinsi(DB *gorm.DB, id int) ([]entity.Kabupaten, error)
}

type KabupatenRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewKabupatenRepository(configDatabase *config.Database) KabupatenRepositoryInterface {
	return &KabupatenRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *KabupatenRepositoryImplementation) FindAllKabupatenByIdProvinsi(DB *gorm.DB, id int) ([]entity.Kabupaten, error) {
	var kabupatens []entity.Kabupaten
	results := DB.Where("idprop = ?", id).Find(&kabupatens)
	return kabupatens, results.Error
}
