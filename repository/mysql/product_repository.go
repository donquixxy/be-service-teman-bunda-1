package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	FindAllProducts(DB *gorm.DB, limit int, page int) ([]entity.Product, error)
	FindProductsBySearch(DB *gorm.DB, productName string) ([]entity.Product, error)
	FindProductById(DB *gorm.DB, id string) (entity.Product, error)
	FindProductByIdCategory(DB *gorm.DB, idCategory string) ([]entity.Product, error)
}

type ProductRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewProductRepository(configDatabase *config.Database) ProductRepositoryInterface {
	return &ProductRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *ProductRepositoryImplementation) FindAllProducts(DB *gorm.DB, limit int, page int) ([]entity.Product, error) {
	var products []entity.Product
	results := DB.Joins("ProductDiscount").
		Limit(limit).
		Offset(page - 1).
		Find(&products)
	return products, results.Error
}

func (repository *ProductRepositoryImplementation) FindProductsBySearch(DB *gorm.DB, productName string) ([]entity.Product, error) {
	var products []entity.Product
	results := DB.Where("products.product_name LIKE ?", "%"+productName+"%").Joins("ProductDiscount").Find(&products)
	return products, results.Error
}

func (repository *ProductRepositoryImplementation) FindProductById(DB *gorm.DB, id string) (entity.Product, error) {
	var product entity.Product
	results := DB.Where("products.id = ?", id).Joins("ProductDiscount").Find(&product)
	return product, results.Error
}

func (repository *ProductRepositoryImplementation) FindProductByIdCategory(DB *gorm.DB, idCategory string) ([]entity.Product, error) {
	var products []entity.Product
	results := DB.Where("products.id_category = ?", idCategory).Joins("ProductDiscount").Find(&products)
	return products, results.Error
}
