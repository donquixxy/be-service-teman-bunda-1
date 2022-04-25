package mysql

import (
	"time"

	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type OrderRepositoryInterface interface {
	FindOrderByDate(DB *gorm.DB, idUser string) ([]entity.Order, error)
	CreateOrder(DB *gorm.DB, order entity.Order) (entity.Order, error)
}

type OrderRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewOrderRepository(configDatabase *config.Database) OrderRepositoryInterface {
	return &OrderRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *OrderRepositoryImplementation) FindOrderByDate(DB *gorm.DB, idUser string) ([]entity.Order, error) {
	var order []entity.Order
	now := time.Now()
	month := now.Month()
	results := DB.Where("orders_transaction.id_user = ?", idUser).Where("month(ordered_at) = ?", int(month)).Find(&order)
	return order, results.Error
}

func (repository *OrderRepositoryImplementation) CreateOrder(DB *gorm.DB, order entity.Order) (entity.Order, error) {
	results := DB.Create(order)
	return order, results.Error
}
