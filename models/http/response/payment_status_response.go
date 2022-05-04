package response

import (
	modelService "github.com/tensuqiuwulu/be-service-teman-bunda/models/service"
)

type PaymentStatusResponse struct {
	Status int `json:"status"`
}

func ToPaymentStatusResponse(paymentStatus modelService.PaymentStatusResponse) (paymentStatusResponse PaymentStatusResponse) {
	paymentStatusResponse.Status = paymentStatus.Data.Status
	return paymentStatusResponse
}
