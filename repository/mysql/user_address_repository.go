package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type UserShippingAddressRepositoryInterface interface {
	CreateUserShippingAddress(DB *gorm.DB, userShippingAddress entity.UserShippingAddress) (entity.UserShippingAddress, error)
	FindUserShippingAddressByIdUser(DB *gorm.DB, idUser string) ([]entity.UserShippingAddress, error)
	DeleteUserShippingAddress(DB *gorm.DB, idUserShippingAddress string) error
}

type UserShippingAddressRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewUserShippingAddressRepository(configDatabase *config.Database) UserShippingAddressRepositoryInterface {
	return &UserShippingAddressRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *UserShippingAddressRepositoryImplementation) DeleteUserShippingAddress(DB *gorm.DB, idUserShippingAddress string) error {
	result := DB.Where("id = ?", idUserShippingAddress).Delete(&entity.UserShippingAddress{})
	return result.Error
}

func (repository *UserShippingAddressRepositoryImplementation) CreateUserShippingAddress(DB *gorm.DB, userAddress entity.UserShippingAddress) (entity.UserShippingAddress, error) {
	results := DB.Create(userAddress)
	return userAddress, results.Error
}

func (repository *UserShippingAddressRepositoryImplementation) FindUserShippingAddressByIdUser(DB *gorm.DB, idUser string) ([]entity.UserShippingAddress, error) {
	var userShippingAddresss []entity.UserShippingAddress
	results := DB.Where("id_user = ?", idUser).Find(&userShippingAddresss)
	return userShippingAddresss, results.Error
}
