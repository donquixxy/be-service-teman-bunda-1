package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type KecamatanRepositoryInterface interface {
	FindAllKecamatanByIdKabupaten(DB *gorm.DB, id int) ([]entity.Kecamatan, error)
}

type KecamatanRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewKecamatanRepository(configDatabase *config.Database) KecamatanRepositoryInterface {
	return &KecamatanRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *KecamatanRepositoryImplementation) FindAllKecamatanByIdKabupaten(DB *gorm.DB, id int) ([]entity.Kecamatan, error) {
	var kecamatans []entity.Kecamatan
	results := DB.Where("idkabu = ?", id).Find(&kecamatans)
	return kecamatans, results.Error
}
