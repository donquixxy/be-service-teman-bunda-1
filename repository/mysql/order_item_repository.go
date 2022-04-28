package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type OrderItemRepositoryInterface interface {
	CreateOrderItems(DB *gorm.DB, order []entity.OrderItem) error
	FindOrderItemsByIdOrder(DB *gorm.DB, idOrder string) ([]entity.OrderItem, error)
}

type OrderItemRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewOrderItemRepository(configDatabase *config.Database) OrderItemRepositoryInterface {
	return &OrderItemRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *OrderItemRepositoryImplementation) CreateOrderItems(DB *gorm.DB, orderItems []entity.OrderItem) error {
	results := DB.Create(orderItems)
	return results.Error
}

func (repository *OrderItemRepositoryImplementation) FindOrderItemsByIdOrder(DB *gorm.DB, idOrder string) ([]entity.OrderItem, error) {
	var orderItems []entity.OrderItem
	results := DB.Where("id_order = ?", idOrder).Find(&orderItems)
	return orderItems, results.Error
}
