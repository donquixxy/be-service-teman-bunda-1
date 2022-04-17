package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUser(DB *gorm.DB, user entity.User) (entity.User, error)
	FindUserByUsername(DB *gorm.DB, username string) (entity.User, error)
	FindUserByEmail(DB *gorm.DB, email string) (entity.User, error)
	FindUserByPhone(DB *gorm.DB, phone string) (entity.User, error)
	FindUserByReferal(DB *gorm.DB, referalCode string) (entity.User, error)
	FindUserById(DB *gorm.DB, id string) (entity.User, error)
}

type UserRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewUserRepository(configDatabase *config.Database) UserRepositoryInterface {
	return &UserRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *UserRepositoryImplementation) CreateUser(DB *gorm.DB, user entity.User) (entity.User, error) {
	results := DB.Create(user)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) FindUserByUsername(DB *gorm.DB, username string) (entity.User, error) {
	var user entity.User
	results := DB.Where("username = ?", username).First(&user)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) FindUserByReferal(DB *gorm.DB, referalCode string) (entity.User, error) {
	var user entity.User
	results := DB.Where("referal_code = ?", referalCode).First(&user)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) FindUserByEmail(DB *gorm.DB, email string) (entity.User, error) {
	var user entity.User
	results := DB.Joins("FamilyMembers").Find(&user, "FamilyMembers.email = ?", email)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) FindUserByPhone(DB *gorm.DB, phone string) (entity.User, error) {
	var user entity.User
	results := DB.Joins("FamilyMembers").Find(&user, "FamilyMembers.phone = ?", phone)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) FindUserById(DB *gorm.DB, id string) (entity.User, error) {
	var user entity.User
	results := DB.Joins("JOIN family_members ON family_members.id = users.id_family_members").
		Joins("JOIN family ON family.id = family_members.id_family").
		Preload("FamilyMembers.Family").
		Joins("BalancePoint").
		Where("users.id = ?", id).
		Find(&user)
	return user, results.Error
}
