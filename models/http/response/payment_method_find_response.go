package response

import (
	modelService "github.com/tensuqiuwulu/be-service-teman-bunda/models/service"
)

type FindListPaymentChannelResponse struct {
	PaymentMethod      string  `json:"payment_method"`
	BankCode           string  `json:"bank_code"`
	BankName           string  `json:"bank_name"`
	BankLogo           string  `json:"bank_logo"`
	AdminFee           float64 `json:"admin_fee"`
	AdminFeePercentage float64 `json:"admin_fee_percentage"`
}

func ToFindPaymentMethodResponses(paymentChannelLists []modelService.ListPaymentChannelPayment) (paymentChannelResponses []FindListPaymentChannelResponse) {
	for _, paymentChannelList := range paymentChannelLists {
		var listPaymentChannelResponse FindListPaymentChannelResponse
		listPaymentChannelResponse.PaymentMethod = paymentChannelList.PaymentMethod
		listPaymentChannelResponse.BankCode = paymentChannelList.BankCode
		listPaymentChannelResponse.BankLogo = paymentChannelList.BankLogo
		listPaymentChannelResponse.BankName = paymentChannelList.BankName
		listPaymentChannelResponse.AdminFee = paymentChannelList.AdminFee
		listPaymentChannelResponse.AdminFeePercentage = paymentChannelList.AdminFeePercentage
		paymentChannelResponses = append(paymentChannelResponses, listPaymentChannelResponse)
	}

	return paymentChannelResponses
}
