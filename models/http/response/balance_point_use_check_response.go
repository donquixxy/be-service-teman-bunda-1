package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type BalancePointUseCheck struct {
	Id             string  `json:"id"`
	IdUser         string  `json:"id_user"`
	UsePointAmount float64 `json:"use_point_amount"`
}

func ToBalancePointUseCheck(balancePoint entity.BalancePoint, usePointAmount float64) (balancePointResponse BalancePointUseCheck) {
	balancePointResponse.Id = balancePoint.Id
	balancePointResponse.IdUser = balancePoint.IdUser
	balancePointResponse.UsePointAmount = usePointAmount
	return balancePointResponse
}

// Test CI
