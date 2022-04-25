package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type CreateOrderResponse struct {
	IdOrder string `json:"id_order"`
}

func ToCreateOrderResponse(order entity.Order, orderItems []entity.OrderItem) (orderResponse CreateOrderResponse) {
	orderResponse.IdOrder = order.Id
	return orderResponse
}
