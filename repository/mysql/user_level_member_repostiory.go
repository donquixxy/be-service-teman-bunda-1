package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type UserLevelMemberRepositoryInterface interface {
	FindUserLevelMemberById(DB *gorm.DB, idLevelMember int) (entity.UserLevelMember, error)
}

type UserLevelMemberRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewUserLevelMemberRepository(configDatabase *config.Database) UserLevelMemberRepositoryInterface {
	return &UserLevelMemberRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *UserLevelMemberRepositoryImplementation) FindUserLevelMemberById(DB *gorm.DB, idLevelMember int) (entity.UserLevelMember, error) {
	var userLevelMember entity.UserLevelMember
	results := DB.Where("id = ?", idLevelMember).Find(&userLevelMember)
	return userLevelMember, results.Error
}
