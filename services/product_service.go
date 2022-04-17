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
}

type ProductServiceImplementation struct {
	ConfigWebserver            config.Webserver
	DB                         *gorm.DB
	Logger                     *logrus.Logger
	ProductRepositoryInterface mysql.ProductRepositoryInterface
}

func NewProductService(configWebserver config.Webserver, DB *gorm.DB, logger *logrus.Logger, productRepositoryInterface mysql.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductServiceImplementation{
		ConfigWebserver:            configWebserver,
		DB:                         DB,
		Logger:                     logger,
		ProductRepositoryInterface: productRepositoryInterface,
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
