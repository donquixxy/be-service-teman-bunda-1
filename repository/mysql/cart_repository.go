package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type CartRepositoryInterface interface {
	// CreateBalancePoint(DB *gorm.DB, balancePoint entity.BalancePoint) (entity.BalancePoint, error)
	// FindBalancePointWithTxByIdUser(DB *gorm.DB, IdUser string) (entity.BalancePoint, error)
	FindCartByIdUser(DB *gorm.DB, IdUser string) ([]entity.Cart, error)
}

type CartRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewCartRepository(configDatabase *config.Database) CartRepositoryInterface {
	return &CartRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

// func (repository *CartRepositoryImplementation) CreateBalancePoint(DB *gorm.DB, balancePoint entity.BalancePoint) (entity.BalancePoint, error) {
// 	results := DB.Create(balancePoint)
// 	return balancePoint, results.Error
// }

func (repository *CartRepositoryImplementation) FindCartByIdUser(DB *gorm.DB, IdUser string) ([]entity.Cart, error) {
	var cart []entity.Cart
	results := DB.Where("cart.id_user = ?", IdUser).Joins("Product").Find(&cart)
	return cart, results.Error
}

// func (repository *BalancePointRepositoryImplementation) FindBalancePointByIdUser(DB *gorm.DB, IdUser string) (entity.BalancePoint, error) {
// 	var balancePoint entity.BalancePoint
// 	results := DB.Where("balance_point.id_user = ?", IdUser).Find(&balancePoint)
// 	return balancePoint, results.Error
// }
