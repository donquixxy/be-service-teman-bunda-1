package response

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type UpdateOrderStatusResponse struct {
	OrderStatus      string `json:"order_status"`
	PaymentStatus    string `json:"payment_status"`
	PaymentSuccessAt string `json:"payment_success_at"`
}

func ToUpdateOrderStatusResponse(order entity.Order) (orderResponse UpdateOrderStatusResponse) {
	orderResponse.OrderStatus = order.OrderSatus
	orderResponse.PaymentStatus = order.PaymentStatus
	orderResponse.PaymentSuccessAt = order.PaymentSuccessAt.Time.Format("2006-01-02 15:04:05")
	return orderResponse
}
