package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
)

type PaymentChannelRepositoryInterface interface {
	// FindPaymentChannelByMethod(DB *gorm.DB) ([]entity.PaymentChannel, error)
}

type PaymentChannelRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewPaymentChannelRepository(configDatabase *config.Database) PaymentChannelRepositoryInterface {
	return &PaymentChannelRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

// func (repository *PaymentChannelRepositoryImplementation) FindPaymentChannelByMethod(DB *gorm.DB) ([]entity.PaymentChannel, error) {
// 	var paymentChannels []entity.PaymentChannel
// 	var paymentMethods []entity.PaymentMethod
// 	DB.Where("is_active = ?", "1").Find(&paymentMethods)

// 	for _, paymentMethod := range paymentMethods {
// 		var channels []entity.PaymentChannel
// 		DB.Joins("PaymentMethod").Where("id_payment_method = ?", paymentMethod.Id).Where("payment_channel.is_active = ?", "1").Find(&channels)

// 		paymentChannels = append(paymentChannels, channels...)
// 	}

// 	results := DB.Find(&paymentChannels)
// 	return paymentChannels, results.Error
// }
