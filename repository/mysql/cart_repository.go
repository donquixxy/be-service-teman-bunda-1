package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type CartRepositoryInterface interface {
	FindCartByIdUser(DB *gorm.DB, IdUser string) ([]entity.Cart, error)
	FindProductInCartByIdUser(DB *gorm.DB, IdUser string, IdProduct string) (entity.Cart, error)
	AddProductToCart(DB *gorm.DB, cart entity.Cart) (entity.Cart, error)
	UpdateProductInCart(DB *gorm.DB, IdCart string, cartEntity entity.Cart) (entity.Cart, error)
	DeleteProductInCart(DB *gorm.DB, IdCart string) (err error)
	FindCartById(DB *gorm.DB, IdCart string) (entity.Cart, error)
	DeleteAllProductInCartByIdUser(DB *gorm.DB, idUser string, carts []entity.Cart) (err error)
}

type CartRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewCartRepository(configDatabase *config.Database) CartRepositoryInterface {
	return &CartRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *CartRepositoryImplementation) AddProductToCart(DB *gorm.DB, cart entity.Cart) (entity.Cart, error) {
	results := DB.Create(cart)
	return cart, results.Error
}

func (repository *CartRepositoryImplementation) FindCartByIdUser(DB *gorm.DB, IdUser string) ([]entity.Cart, error) {
	var cart []entity.Cart
	// results := DB.Where("cart.id_user = ?", IdUser).Joins("Product").Preload("ProductDiscount").Find(&cart)
	results := DB.Where("cart.id_user = ?", IdUser).
		Joins("Product").
		Preload("Product.ProductDiscount").
		Find(&cart)
	// results := DB.Joins("JOIN products on products.id = cart.id_product").
	// 	Joins("JOIN products_discount on products_discount.id_product = products.id").
	// 	Where("cart.id_user = ?", IdUser).
	// 	Find(&cart)
	return cart, results.Error
}

func (repository *CartRepositoryImplementation) FindCartById(DB *gorm.DB, IdCart string) (entity.Cart, error) {
	var cart entity.Cart
	results := DB.Where("id = ?", IdCart).Find(&cart)
	return cart, results.Error
}

func (repository *CartRepositoryImplementation) FindProductInCartByIdUser(DB *gorm.DB, IdUser string, IdProduct string) (entity.Cart, error) {
	var cart entity.Cart
	results := DB.Where("id_user = ?", IdUser).Where("id_product = ?", IdProduct).Find(&cart)
	return cart, results.Error
}

func (repository *CartRepositoryImplementation) UpdateProductInCart(DB *gorm.DB, IdCart string, cart entity.Cart) (entity.Cart, error) {
	result := DB.
		Model(entity.Cart{}).
		Where("id = ?", IdCart).
		Updates(entity.Cart{
			Qty: cart.Qty,
		})
	return cart, result.Error
}

func (repository *CartRepositoryImplementation) DeleteProductInCart(DB *gorm.DB, IdCart string) (err error) {
	result := DB.Where("id = ?", IdCart).Delete(&entity.Cart{})
	return result.Error
}

func (repository *CartRepositoryImplementation) DeleteAllProductInCartByIdUser(DB *gorm.DB, idUser string, carts []entity.Cart) (err error) {
	result := DB.Where("id_user LIKE ?", "%"+idUser+"%").Delete(carts)
	return result.Error
}
