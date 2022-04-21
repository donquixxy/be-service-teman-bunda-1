package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type ShippingRepositoryInterface interface {
	GetShippingCostByIdKelurahan(DB *gorm.DB, ikdKelurahan int) (entity.ShippingCostArea, error)
}

type ShippingRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewShippingRepository(configDatabase *config.Database) ShippingRepositoryInterface {
	return &ShippingRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *ShippingRepositoryImplementation) GetShippingCostByIdKelurahan(DB *gorm.DB, idKelurahan int) (entity.ShippingCostArea, error) {
	var shippingCostArea entity.ShippingCostArea
	results := DB.Joins("JOIN master_kelurahan on master_kelurahan.id_shipping_cost_area = shipping_cost_area.id").
		Where("master_kelurahan.idkelu = ?", idKelurahan).
		Find(&shippingCostArea)
	return shippingCostArea, results.Error
}
