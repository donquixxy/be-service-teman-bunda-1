package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindSettingShippingCost struct {
	Value float64 `json:"shipping_cost"`
}

func ToFindSettingShippingCost(shippingCost entity.Settings) (shippingCostResponse FindSettingShippingCost) {
	shippingCostResponse.Value = shippingCost.Value
	return shippingCostResponse
}
