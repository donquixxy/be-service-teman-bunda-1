package response

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"github.com/tensuqiuwulu/be-service-teman-bunda/utilities"
)

type CreateOrderResponse struct {
	IdOrder        string  `json:"id_order"`
	ReferenceId    string  `json:"reference_id"`
	PaymentNo      string  `json:"payment_no"`
	PaymentName    string  `json:"payment_name"`
	Total          float64 `json:"total"`
	Expired        string  `json:"expired"`
	PaymentMethod  string  `json:"payment_method"`
	PaymentChannel string  `json:"payment_channel"`
	PaymentStatus  string  `json:"payment_status"`
}

func ToCreateOrderResponse(order entity.Order, orderItems []entity.OrderItem, ipaymuData utilities.IpaymuDirectPaymentResponse) (orderResponse CreateOrderResponse) {
	orderResponse.IdOrder = order.Id
	orderResponse.ReferenceId = ipaymuData.Data.ReferenceId
	orderResponse.PaymentNo = ipaymuData.Data.PaymentNo
	orderResponse.PaymentName = ipaymuData.Data.PaymentName
	orderResponse.Total = ipaymuData.Data.Total
	orderResponse.Expired = ipaymuData.Data.Expired
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentChannel = order.PaymentChannel
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentStatus = order.PaymentStatus
	return orderResponse
}
