package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type ProductStockHistoryRepositoryInterface interface {
	AddProductStockHistory(DB *gorm.DB, productStockHistory entity.ProductStockHistory) (entity.ProductStockHistory, error)
}

type ProductStockHistoryRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewProductStockHistoryRepository(configDatabase *config.Database) ProductStockHistoryRepositoryInterface {
	return &ProductStockHistoryRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *ProductStockHistoryRepositoryImplementation) AddProductStockHistory(DB *gorm.DB, productStockHistory entity.ProductStockHistory) (entity.ProductStockHistory, error) {
	results := DB.Create(productStockHistory)
	return productStockHistory, results.Error
}
