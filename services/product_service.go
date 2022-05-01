package services

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type ProductServiceInterface interface {
	FindAllProducts(requestId string, limit int, page int) (productResponses []response.FindProductResponse)
	FindProductsBySearch(requestId string, productName string) (productsResponses []response.FindProductResponse)
	FindProductById(requestId string, id string) (productsResponse response.FindProductResponse)
	FindProductByIdCategory(requestId string, idCategory string) (productsResponses []response.FindProductResponse)
	FindProductByIdSubCategory(requestId string, idSubCategory string) (productsResponses []response.FindProductResponse)
	FindProductByIdBrand(requestId string, idBrand string) (productsResponses []response.FindProductResponse)
}

type ProductServiceImplementation struct {
	ConfigWebserver            config.Webserver
	DB                         *gorm.DB
	Logger                     *logrus.Logger
	ProductRepositoryInterface mysql.ProductRepositoryInterface
	ConfigPayment              config.Payment
}

func NewProductService(
	configWebserver config.Webserver,
	DB *gorm.DB, logger *logrus.Logger,
	productRepositoryInterface mysql.ProductRepositoryInterface,
	configPayment config.Payment) ProductServiceInterface {
	return &ProductServiceImplementation{
		ConfigWebserver:            configWebserver,
		DB:                         DB,
		Logger:                     logger,
		ProductRepositoryInterface: productRepositoryInterface,
		ConfigPayment:              configPayment,
	}
}

func (service *ProductServiceImplementation) FindAllProducts(requestId string, limit int, page int) (productResponses []response.FindProductResponse) {
	products, err := service.ProductRepositoryInterface.FindAllProducts(service.DB, limit, page)
	exceptions.PanicIfError(err, requestId, service.Logger)
	productResponses = response.ToFindProductResponses(products)
	return productResponses
}

func (service *ProductServiceImplementation) FindProductsBySearch(requestId string, productName string) (productResponses []response.FindProductResponse) {
	products, err := service.ProductRepositoryInterface.FindProductsBySearch(service.DB, productName)
	exceptions.PanicIfError(err, requestId, service.Logger)
	productResponses = response.ToFindProductResponses(products)
	return productResponses
}

func (service *ProductServiceImplementation) FindProductById(requestId string, id string) (productResponse response.FindProductResponse) {
	product, _ := service.ProductRepositoryInterface.FindProductById(service.DB, id)
	if product.Id == "" {
		err := errors.New("product not found")
		exceptions.PanicIfRecordNotFound(err, requestId, []string{"Not Found"}, service.Logger)
	}
	productResponse = response.ToFindProductResponse(product)
	return productResponse
}

func (service *ProductServiceImplementation) FindProductByIdCategory(requestId string, idCategory string) (productResponses []response.FindProductResponse) {
	products, err := service.ProductRepositoryInterface.FindProductByIdCategory(service.DB, idCategory)
	exceptions.PanicIfError(err, requestId, service.Logger)
	productResponses = response.ToFindProductResponses(products)
	return productResponses
}

func (service *ProductServiceImplementation) FindProductByIdSubCategory(requestId string, idSubCategory string) (productResponses []response.FindProductResponse) {
	products, err := service.ProductRepositoryInterface.FindProductByIdSubCategory(service.DB, idSubCategory)
	exceptions.PanicIfError(err, requestId, service.Logger)
	productResponses = response.ToFindProductResponses(products)
	return productResponses
}

func (service *ProductServiceImplementation) FindProductByIdBrand(requestId string, idBrand string) (productResponses []response.FindProductResponse) {
	products, err := service.ProductRepositoryInterface.FindProductByIdBrand(service.DB, idBrand)
	exceptions.PanicIfError(err, requestId, service.Logger)
	productResponses = response.ToFindProductResponses(products)
	return productResponses
}
