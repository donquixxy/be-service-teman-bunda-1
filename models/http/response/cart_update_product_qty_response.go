package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type UpdateProductQtyInCartResponse struct {
	Id string `json:"id_cart"`
}

func ToUpdateProductQtyInCartResponse(cart entity.Cart) (updateProductQtyInCartResponse UpdateProductQtyInCartResponse) {
	updateProductQtyInCartResponse.Id = cart.Id
	return updateProductQtyInCartResponse
}
