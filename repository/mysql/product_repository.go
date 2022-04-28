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
	FindProductByIdSubCategory(DB *gorm.DB, idSubCategory string) ([]entity.Product, error)
	FindProductByIdBrand(DB *gorm.DB, idBrand string) ([]entity.Product, error)
	UpdateProductStock(DB *gorm.DB, idProduct string, product entity.Product) (entity.Product, error)
}

type ProductRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewProductRepository(configDatabase *config.Database) ProductRepositoryInterface {
	return &ProductRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *ProductRepositoryImplementation) UpdateProductStock(DB *gorm.DB, idProduct string, product entity.Product) (entity.Product, error) {
	result := DB.
		Model(entity.Product{}).
		Where("id = ?", idProduct).
		Updates(entity.Product{
			Stock: product.Stock,
		})
	return product, result.Error
}

func (repository *ProductRepositoryImplementation) FindAllProducts(DB *gorm.DB, limit int, page int) ([]entity.Product, error) {
	var products []entity.Product
	results := DB.Joins("ProductDiscount").
		Limit(limit).
		Offset(page-1).
		Where("products.published = ?", "1").
		Find(&products)
	return products, results.Error
}

func (repository *ProductRepositoryImplementation) FindProductsBySearch(DB *gorm.DB, productName string) ([]entity.Product, error) {
	var products []entity.Product
	results := DB.Where("products.product_name LIKE ?", "%"+productName+"%").Where("products.published = ?", "1").Joins("ProductDiscount").Find(&products)
	return products, results.Error
}

func (repository *ProductRepositoryImplementation) FindProductById(DB *gorm.DB, id string) (entity.Product, error) {
	var product entity.Product
	results := DB.Where("products.id = ?", id).Where("products.published = ?", "1").Joins("ProductDiscount").Find(&product)
	return product, results.Error
}

func (repository *ProductRepositoryImplementation) FindProductByIdCategory(DB *gorm.DB, idCategory string) ([]entity.Product, error) {
	var products []entity.Product
	results := DB.Where("products.id_category = ?", idCategory).Where("products.published = ?", "1").Joins("ProductDiscount").Find(&products)
	return products, results.Error
}

func (repository *ProductRepositoryImplementation) FindProductByIdSubCategory(DB *gorm.DB, idSubCategory string) ([]entity.Product, error) {
	var products []entity.Product
	results := DB.Where("products.id_sub_category = ?", idSubCategory).Where("products.published = ?", "1").Joins("ProductDiscount").Find(&products)
	return products, results.Error
}

func (repository *ProductRepositoryImplementation) FindProductByIdBrand(DB *gorm.DB, idBrand string) ([]entity.Product, error) {
	var products []entity.Product
	results := DB.Where("products.id_brand = ?", idBrand).Where("products.published = ?", "1").Joins("ProductDiscount").Find(&products)
	return products, results.Error
}
