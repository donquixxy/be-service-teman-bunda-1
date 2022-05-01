package services

import (
	"github.com/sirupsen/logrus"
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/exceptions"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/http/response"
	"github.com/tensuqiuwulu/be-service-teman-bunda/repository/mysql"
	"gorm.io/gorm"
)

type ProductBrandServiceInterface interface {
	FindAllProductBrand(requestId string) (productBrandResponses []response.FindAllProductBrandResponse)
}

type ProductBrandServiceImplementation struct {
	ConfigWebserver                 config.Webserver
	DB                              *gorm.DB
	Logger                          *logrus.Logger
	ProductBrandRepositoryInterface mysql.ProductBrandRepositoryInterface
}

func NewProductBrandService(configWebserver config.Webserver, DB *gorm.DB, logger *logrus.Logger, productBrandRepositoryInterface mysql.ProductBrandRepositoryInterface) ProductBrandServiceInterface {
	return &ProductBrandServiceImplementation{
		ConfigWebserver:                 configWebserver,
		DB:                              DB,
		Logger:                          logger,
		ProductBrandRepositoryInterface: productBrandRepositoryInterface,
	}
}

func (service *ProductBrandServiceImplementation) FindAllProductBrand(requestId string) (productBrandResponses []response.FindAllProductBrandResponse) {
	productBrands, err := service.ProductBrandRepositoryInterface.FindAllProductBrand(service.DB)
	exceptions.PanicIfError(err, requestId, service.Logger)
	productBrandResponses = response.ToFindAllProductBrandResponses(productBrands)
	return productBrandResponses
}
