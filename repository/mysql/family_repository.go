package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type FamilyRepositoryInterface interface {
	CreateFamily(DB *gorm.DB, user entity.Family) (entity.Family, error)
}

type FamilyRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewFamilyRepository(configDatabase *config.Database) FamilyRepositoryInterface {
	return &FamilyRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *FamilyRepositoryImplementation) CreateFamily(DB *gorm.DB, family entity.Family) (entity.Family, error) {
	results := DB.Create(family)
	return family, results.Error
}
