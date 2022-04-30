package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindAllProductBrandResponse struct {
	Id           string `json:"id"`
	BrandName    string `json:"brand_name"`
	BrandLogoUrl string `json:"brand_logo_url"`
}

func ToFindAllProductBrandResponses(productBrands []entity.ProductBrand) (productBrandResponses []FindAllProductBrandResponse) {
	for _, productBrand := range productBrands {
		var productBrandResponse FindAllProductBrandResponse
		productBrandResponse.Id = productBrand.Id
		productBrandResponse.BrandName = productBrand.BrandName
		productBrandResponse.BrandLogoUrl = productBrand.BrandLogoUrl
		productBrandResponses = append(productBrandResponses, productBrandResponse)
	}
	return productBrandResponses
}
