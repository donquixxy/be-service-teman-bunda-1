package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindOrderByUserResponse struct {
	IdOrder     string  `json:"id_order"`
	NoOrder     string  `json:"no_order"`
	Address     string  `json:"address"`
	OrderStatus string  `json:"order_status"`
	TotalBill   float64 `json:"total_bill"`
}

func ToFindOrderByUserResponse(orders []entity.Order) (orderResponses []FindOrderByUserResponse) {
	for _, order := range orders {
		var orderResponse FindOrderByUserResponse
		orderResponse.IdOrder = order.Id
		orderResponse.NoOrder = order.NumberOrder
		orderResponse.Address = order.Address
		orderResponse.OrderStatus = order.OrderSatus
		orderResponse.TotalBill = order.TotalBill
		orderResponses = append(orderResponses, orderResponse)
	}
	return orderResponses
}
