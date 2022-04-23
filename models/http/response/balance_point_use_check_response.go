package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type BalancePointUseCheck struct {
	Id            string  `json:"id"`
	IdUser        string  `json:"id_user"`
	BalancePoints float64 `json:"balance_points"`
}

func ToBalancePointUseCheck(balancePoint entity.BalancePoint) (balancePointResponse BalancePointUseCheck) {
	balancePointResponse.Id = balancePoint.Id
	balancePointResponse.IdUser = balancePoint.IdUser
	balancePointResponse.BalancePoints = balancePoint.BalancePoints
	return balancePointResponse
}

// Test CI
