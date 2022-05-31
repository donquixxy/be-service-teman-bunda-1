package response

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type FindOrderByUserResponse struct {
	IdOrder     string  `json:"id_order"`
	NoOrder     string  `json:"no_order"`
	Address     string  `json:"address"`
	OrderStatus string  `json:"order_status"`
	TotalBill   float64 `json:"total_bill"`
	OrderedAt   string  `json:"order_date"`
}

func ToFindOrderByUserResponse(orders []entity.Order) (orderResponses []FindOrderByUserResponse) {
	for _, order := range orders {
		var orderResponse FindOrderByUserResponse
		orderResponse.IdOrder = order.Id
		orderResponse.NoOrder = order.NumberOrder
		orderResponse.Address = order.Address
		orderResponse.OrderStatus = order.OrderSatus
		orderResponse.TotalBill = order.PaymentByCash
		orderResponse.OrderedAt = order.OrderedAt.Format("2006-01-02 15:04:05")
		orderResponses = append(orderResponses, orderResponse)
	}
	return orderResponses
}
