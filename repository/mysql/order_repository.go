package mysql

import (
	"time"

	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type OrderRepositoryInterface interface {
	FindOrderByUser(DB *gorm.DB, idUser string, orderStatus string) ([]entity.Order, error)
	FindOrderByDate(DB *gorm.DB, idUser string) ([]entity.Order, error)
	FindOrderByNumberOrder(DB *gorm.DB, numberOrder string) (entity.Order, error)
	FindOrderById(DB *gorm.DB, idOrder string) (entity.Order, error)
	CreateOrder(DB *gorm.DB, order entity.Order) (entity.Order, error)
	UpdateOrderStatus(DB *gorm.DB, numberOrder string, order entity.Order) (entity.Order, error)
	UpdateOrderPayment(DB *gorm.DB, numberOrder string, order entity.Order) (entity.Order, error)
}

type OrderRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewOrderRepository(configDatabase *config.Database) OrderRepositoryInterface {
	return &OrderRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *OrderRepositoryImplementation) FindOrderByUser(DB *gorm.DB, idUser string, orderStatus string) ([]entity.Order, error) {
	var order []entity.Order
	if orderStatus == "" {
		results := DB.Where("id_user = ?", idUser).Find(&order)
		return order, results.Error
	} else {
		results := DB.Where("id_user = ?", idUser).Where("order_status = ?", order).Find(&order)
		return order, results.Error
	}
}

func (repository *OrderRepositoryImplementation) FindOrderByDate(DB *gorm.DB, idUser string) ([]entity.Order, error) {
	var order []entity.Order
	now := time.Now()
	month := now.Month()
	results := DB.Where("orders_transaction.id_user = ?", idUser).Where("month(ordered_at) = ?", int(month)).Where("order_status = ?", "Selesai").Find(&order)
	return order, results.Error
}

func (repository *OrderRepositoryImplementation) FindOrderByNumberOrder(DB *gorm.DB, numberOrder string) (entity.Order, error) {
	var order entity.Order
	results := DB.Where("orders_transaction.number_order = ?", numberOrder).First(&order)
	return order, results.Error
}

func (repository *OrderRepositoryImplementation) FindOrderById(DB *gorm.DB, idOrder string) (entity.Order, error) {
	var order entity.Order
	results := DB.Where("orders_transaction.id = ?", idOrder).First(&order)
	return order, results.Error
}

func (repository *OrderRepositoryImplementation) CreateOrder(DB *gorm.DB, order entity.Order) (entity.Order, error) {
	results := DB.Create(order)
	return order, results.Error
}

func (repository *OrderRepositoryImplementation) UpdateOrderStatus(DB *gorm.DB, NumberOrder string, order entity.Order) (entity.Order, error) {
	result := DB.
		Model(entity.Order{}).
		Where("number_order = ?", NumberOrder).
		Updates(entity.Order{
			PaymentStatus:    order.PaymentStatus,
			OrderSatus:       order.OrderSatus,
			PaymentSuccessAt: order.PaymentSuccessAt,
		})
	return order, result.Error
}

func (repository *OrderRepositoryImplementation) UpdateOrderPayment(DB *gorm.DB, NumberOrder string, order entity.Order) (entity.Order, error) {
	result := DB.
		Model(entity.Order{}).
		Where("number_order = ?", NumberOrder).
		Updates(entity.Order{
			PaymentNo:      order.PaymentNo,
			PaymentExpired: order.PaymentExpired,
			PaymentName:    order.PaymentName,
		})
	return order, result.Error
}
