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
	CountUserByRegistrationReferal(DB *gorm.DB, referal string) (userCount int, err error)
	SaveUserRefreshToken(DB *gorm.DB, id string, refreshToken string) (int64, error)
	FindUserByUsernameAndRefreshToken(DB *gorm.DB, username string, refresh_token string) (entity.User, error)
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
	results := DB.Where("users.id = ?", id).
		Joins("FamilyMembers").
		Joins("BalancePoint").
		Find(&user)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) CountUserByRegistrationReferal(DB *gorm.DB, referalCode string) (countUser int, err error) {
	var user entity.User
	results := DB.Model(&entity.User{}).Where("registration_referal_code = ?", referalCode).Find(&user)
	return int(results.RowsAffected), results.Error
}

func (repository *UserRepositoryImplementation) FindUserByUsernameAndRefreshToken(DB *gorm.DB, username string, refresh_token string) (entity.User, error) {
	var user entity.User
	results := DB.Where("username = ?", username).Where("refresh_token = ?", refresh_token).First(&user)
	return user, results.Error
}

func (repository *UserRepositoryImplementation) SaveUserRefreshToken(DB *gorm.DB, id string, refreshToken string) (int64, error) {
	results := DB.Exec("UPDATE `users` SET refresh_token = ? WHERE id = ?", refreshToken, id)
	return results.RowsAffected, results.Error
}
