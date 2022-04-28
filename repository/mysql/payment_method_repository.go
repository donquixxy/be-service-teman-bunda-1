package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type PaymentMethodRepositoryInterface interface {
	FindPaymentMethod(DB *gorm.DB) ([]entity.PaymentMethod, error)
}

type PaymentMethodRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewPaymentMethodRepository(configDatabase *config.Database) PaymentMethodRepositoryInterface {
	return &PaymentMethodRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *PaymentMethodRepositoryImplementation) FindPaymentMethod(DB *gorm.DB) ([]entity.PaymentMethod, error) {
	var paymentMethod []entity.PaymentMethod
	results := DB.Where("is_active = ?", "1").Find(&paymentMethod)
	return paymentMethod, results.Error
}
