package response

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	modelService "github.com/tensuqiuwulu/be-service-teman-bunda/models/service"
)

type CreateOrderResponse struct {
	IdOrder        string  `json:"id_order"`
	TrxId          int     `json:"trx_id"`
	ReferenceId    string  `json:"reference_id"`
	PaymentNo      string  `json:"payment_no"`
	PaymentName    string  `json:"payment_name"`
	Total          float64 `json:"total"`
	Expired        string  `json:"expired"`
	PaymentMethod  string  `json:"payment_method"`
	PaymentChannel string  `json:"payment_channel"`
	PaymentStatus  string  `json:"payment_status"`
	BankName       string  `json:"bank_name"`
	BankLogo       string  `json:"bank_logo"`
}

func ToCreateOrderVaResponse(
	order entity.Order,
	TrxId int,
	payment modelService.PaymentResponse,
	bankVa entity.BankVa) (orderResponse CreateOrderResponse) {
	orderResponse.IdOrder = order.Id
	orderResponse.TrxId = TrxId
	orderResponse.ReferenceId = payment.Data.ReferenceId
	orderResponse.PaymentNo = payment.Data.PaymentNo
	orderResponse.PaymentName = "Teman Bunda"
	orderResponse.Total = payment.Data.Total
	orderResponse.Expired = payment.Data.Expired
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentChannel = order.PaymentChannel
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentStatus = order.PaymentStatus
	orderResponse.BankName = bankVa.BankName
	orderResponse.BankLogo = bankVa.BankLogo
	return orderResponse
}

func ToCreateOrderTransferResponse(
	order entity.Order,
	payment modelService.PaymentResponse,
	bankTransfer entity.BankTransfer) (orderResponse CreateOrderResponse) {
	orderResponse.IdOrder = order.Id
	orderResponse.ReferenceId = order.NumberOrder
	orderResponse.PaymentNo = payment.Data.PaymentNo
	orderResponse.Total = payment.Data.Total
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentChannel = order.PaymentChannel
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentStatus = order.PaymentStatus
	orderResponse.BankName = bankTransfer.BankName
	orderResponse.PaymentName = bankTransfer.BankAn
	orderResponse.BankLogo = bankTransfer.BankLogo
	orderResponse.Expired = payment.Data.Expired
	return orderResponse
}

func ToCreateOrderCodResponse(
	order entity.Order) (orderResponse CreateOrderResponse) {
	orderResponse.IdOrder = order.Id
	orderResponse.Total = order.PaymentByCash
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentChannel = order.PaymentChannel
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentStatus = order.PaymentStatus
	return orderResponse
}

func ToCreateOrderFullPointResponse(
	order entity.Order) (orderResponse CreateOrderResponse) {
	orderResponse.IdOrder = order.Id
	orderResponse.Total = order.PaymentByPoint
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentChannel = order.PaymentChannel
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentStatus = order.PaymentStatus
	return orderResponse
}
