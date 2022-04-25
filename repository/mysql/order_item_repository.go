package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type OrderItemRepositoryInterface interface {
	CreateOrderItems(DB *gorm.DB, order []entity.OrderItem) ([]entity.OrderItem, error)
}

type OrderItemRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewOrderItemRepository(configDatabase *config.Database) OrderItemRepositoryInterface {
	return &OrderItemRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *OrderItemRepositoryImplementation) CreateOrderItems(DB *gorm.DB, orderItems []entity.OrderItem) ([]entity.OrderItem, error) {
	results := DB.Create(orderItems)
	return orderItems, results.Error
}
