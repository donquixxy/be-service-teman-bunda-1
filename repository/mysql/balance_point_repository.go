package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type BalancePointRepositoryInterface interface {
	CreateBalancePoint(DB *gorm.DB, balancePoint entity.BalancePoint) (entity.BalancePoint, error)
	FindBalancePointByIdUser(DB *gorm.DB, IdUser string) (entity.BalancePoint, error)
	BalancePointUseCheck(DB *gorm.DB, IdUser string) (entity.BalancePoint, error)
}

type BalancePointRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewBalancePointRepository(configDatabase *config.Database) BalancePointRepositoryInterface {
	return &BalancePointRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *BalancePointRepositoryImplementation) CreateBalancePoint(DB *gorm.DB, balancePoint entity.BalancePoint) (entity.BalancePoint, error) {
	results := DB.Create(balancePoint)
	return balancePoint, results.Error
}

func (repository *BalancePointRepositoryImplementation) FindBalancePointByIdUser(DB *gorm.DB, IdUser string) (entity.BalancePoint, error) {
	var balancePoint entity.BalancePoint
	results := DB.Where("balance_point.id_user = ?", IdUser).Find(&balancePoint)
	return balancePoint, results.Error
}

func (repository *BalancePointRepositoryImplementation) BalancePointUseCheck(DB *gorm.DB, IdUser string) (entity.BalancePoint, error) {
	var balancePoint entity.BalancePoint
	results := DB.Where("balance_point.id_user = ?", IdUser).Find(&balancePoint)
	return balancePoint, results.Error
}
