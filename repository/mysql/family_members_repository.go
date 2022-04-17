package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type FamilyMembersRepositoryInterface interface {
	CreateFamilyMembers(DB *gorm.DB, user entity.FamilyMembers) (entity.FamilyMembers, error)
}

type FamilyMembersRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewFamilyMembersRepository(configDatabase *config.Database) FamilyMembersRepositoryInterface {
	return &FamilyMembersRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *FamilyMembersRepositoryImplementation) CreateFamilyMembers(DB *gorm.DB, familyMembers entity.FamilyMembers) (entity.FamilyMembers, error) {
	results := DB.Create(familyMembers)
	return familyMembers, results.Error
}
