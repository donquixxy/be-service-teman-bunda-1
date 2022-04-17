package response

import (
	"github.com/shopspring/decimal"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type FindProductResponse struct {
	Id          string          `json:"id"`
	ProductName string          `json:"product_name"`
	Price       decimal.Decimal `json:"price"`
	Description string          `json:"description"`
	PictureUrl  string          `json:"picture_url"`
	Thumbnail   string          `json:"thumbnail"`
	Stock       int             `json:"stock"`
}

func ToFindProductResponses(products []entity.Product) (productResponses []FindProductResponse) {
	for _, product := range products {
		var productResponse FindProductResponse
		productResponse.Id = product.Id
		productResponse.ProductName = product.ProductName
		productResponse.Price = product.Price
		productResponse.Description = product.Description
		productResponse.PictureUrl = product.PictureUrl
		productResponse.Thumbnail = product.Thumbnail
		productResponse.Stock = product.Stock
		productResponses = append(productResponses, productResponse)
	}
	return productResponses
}

func ToFindProductResponse(product entity.Product) (productResponse FindProductResponse) {
	productResponse.Id = product.Id
	productResponse.ProductName = product.ProductName
	productResponse.Price = product.Price
	productResponse.Description = product.Description
	productResponse.PictureUrl = product.PictureUrl
	productResponse.Thumbnail = product.Thumbnail
	productResponse.Stock = product.Stock
	return productResponse
}
