package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type AddProductToCartResponse struct {
	Id string `json:"id_cart"`
}

func ToAddProductToCartResponse(cart entity.Cart) (addProductToCartResponse AddProductToCartResponse) {
	addProductToCartResponse.Id = cart.Id
	return addProductToCartResponse
}
