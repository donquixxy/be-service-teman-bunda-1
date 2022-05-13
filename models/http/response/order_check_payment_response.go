package response

import (
	"time"

	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type OrderCheckPayment struct {
	IdOrder        string    `json:"id_order"`
	PaymentNo      string    `json:"payment_no"`
	PaymentName    string    `json:"payment_name"`
	Total          float64   `json:"total"`
	Expired        time.Time `json:"expired"`
	PaymentMethod  string    `json:"payment_method"`
	PaymentChannel string    `json:"payment_channel"`
	PaymentStatus  string    `json:"payment_status"`
	BankName       string    `json:"bank_name"`
	BankLogo       string    `json:"bank_logo"`
}

func ToOrderCheckVaPaymentResponse(
	order entity.Order,
	bankVa entity.BankVa) (orderResponse OrderCheckPayment) {
	orderResponse.IdOrder = order.Id
	orderResponse.PaymentNo = order.PaymentNo
	orderResponse.PaymentName = "Teman Bunda"
	orderResponse.Total = order.PaymentByCash
	// orderResponse.Expired = order.PaymentDueDate.Time
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentChannel = order.PaymentChannel
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentStatus = order.PaymentStatus
	orderResponse.BankName = bankVa.BankName
	orderResponse.BankLogo = bankVa.BankLogo
	return orderResponse
}

func ToOrderCheckTransferPaymentResponse(
	order entity.Order,
	bankTransfer entity.BankTransfer) (orderResponse OrderCheckPayment) {
	orderResponse.IdOrder = order.Id
	orderResponse.PaymentNo = order.PaymentNo
	orderResponse.Total = order.PaymentByCash
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentChannel = order.PaymentChannel
	orderResponse.PaymentMethod = order.PaymentMethod
	orderResponse.PaymentStatus = order.PaymentStatus
	orderResponse.BankName = bankTransfer.BankName
	orderResponse.PaymentName = bankTransfer.BankAn
	orderResponse.BankLogo = bankTransfer.BankLogo
	return orderResponse
}
