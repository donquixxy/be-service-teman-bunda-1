package response

import (
	"time"

	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type UpdateOrderStatusResponse struct {
	OrderStatus      string    `json:"order_status"`
	PaymentStatus    string    `json:"payment_status"`
	PaymentSuccessAt time.Time `json:"payment_success_at"`
}

func ToUpdateOrderStatusResponse(order entity.Order) (orderResponse UpdateOrderStatusResponse) {
	orderResponse.OrderStatus = order.OrderSatus
	orderResponse.PaymentStatus = order.PaymentStatus
	orderResponse.PaymentSuccessAt = order.PaymentSuccessAt.Time
	return orderResponse
}
