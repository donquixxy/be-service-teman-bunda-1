package response

import (
	"github.com/shopspring/decimal"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type FindProductResponse struct {
	Id          string          `json:"id"`
	IdCategory  int             `json:"id_category"`
	ProductName string          `json:"product_name"`
	Price       decimal.Decimal `json:"price"`
	Description string          `json:"description"`
	PictureUrl  string          `json:"picture_url"`
	Thumbnail   string          `json:"thumbnail"`
	Stock       int             `json:"stock"`
	Percentage  decimal.Decimal `json:"discount_percentage"`
	Nominal     decimal.Decimal `json:"discount_nominal"`
}

func ToFindProductResponses(products []entity.Product) (productResponses []FindProductResponse) {
	for _, product := range products {
		var productResponse FindProductResponse
		productResponse.Id = product.Id
		productResponse.IdCategory = product.IdCategory
		productResponse.ProductName = product.ProductName
		productResponse.Price = product.Price
		productResponse.Description = product.Description
		productResponse.PictureUrl = product.PictureUrl
		productResponse.Thumbnail = product.Thumbnail
		productResponse.Stock = product.Stock
		productResponse.Percentage = product.ProductDiscount.Percentage
		productResponse.Nominal = product.ProductDiscount.Nominal
		productResponses = append(productResponses, productResponse)
	}
	return productResponses
}

func ToFindProductResponse(product entity.Product) (productResponse FindProductResponse) {
	productResponse.Id = product.Id
	productResponse.IdCategory = product.IdCategory
	productResponse.ProductName = product.ProductName
	productResponse.Price = product.Price
	productResponse.Description = product.Description
	productResponse.PictureUrl = product.PictureUrl
	productResponse.Thumbnail = product.Thumbnail
	productResponse.Stock = product.Stock
	productResponse.Percentage = product.ProductDiscount.Percentage
	productResponse.Nominal = product.ProductDiscount.Nominal
	return productResponse
}
