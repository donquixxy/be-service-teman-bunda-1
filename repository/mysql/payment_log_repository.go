package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type PaymentLogRepositoryInterface interface {
	CreatePaymentLog(DB *gorm.DB, paymentLog entity.PaymentLog) (entity.PaymentLog, error)
	FindPaymentLogByIdOrder(DB *gorm.DB, idOrder string) (entity.PaymentLog, error)
}

type PaymentLogRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewPaymentLogRepository(configDatabase *config.Database) PaymentLogRepositoryInterface {
	return &PaymentLogRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *PaymentLogRepositoryImplementation) CreatePaymentLog(DB *gorm.DB, paymentLog entity.PaymentLog) (entity.PaymentLog, error) {
	results := DB.Create(paymentLog)
	return paymentLog, results.Error
}

func (repository *PaymentLogRepositoryImplementation) FindPaymentLogByIdOrder(DB *gorm.DB, idOrder string) (entity.PaymentLog, error) {
	var paymentLog entity.PaymentLog
	results := DB.Where("id_order = ?", idOrder).Where("type_log = ?", "Create Trx Ipaymu").First(&paymentLog)
	return paymentLog, results.Error
}
