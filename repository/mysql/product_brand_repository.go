package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type ProductBrandRepositoryInterface interface {
	FindAllProductBrand(DB *gorm.DB) ([]entity.ProductBrand, error)
}

type ProductBrandRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewProductBrandRepository(configDatabase *config.Database) ProductBrandRepositoryInterface {
	return &ProductBrandRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *ProductBrandRepositoryImplementation) FindAllProductBrand(DB *gorm.DB) ([]entity.ProductBrand, error) {
	var productBrands []entity.ProductBrand
	results := DB.Find(&productBrands)
	return productBrands, results.Error
}
