package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type GetShippingCostByIdKelurahanResponse struct {
	ShippingCost float64 `json:"shipping_cost"`
}

func ToGetShippingCostByIdKelurahanResponse(shippingCost entity.ShippingCostArea) (getShippingCostResponse GetShippingCostByIdKelurahanResponse) {
	getShippingCostResponse.ShippingCost = shippingCost.ShippingCost
	return getShippingCostResponse
}
