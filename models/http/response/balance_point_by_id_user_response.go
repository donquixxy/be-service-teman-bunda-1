package response

import (
	"github.com/shopspring/decimal"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type FindBalancePointByIdUser struct {
	Id            string          `json:"id"`
	IdUser        string          `json:"id_user"`
	BalancePoints decimal.Decimal `json:"balance_points"`
}

func ToFindBalancePointByIdUser(balancePoint entity.BalancePoint) (balancePointResponse FindBalancePointByIdUser) {
	balancePointResponse.Id = balancePoint.Id
	balancePointResponse.IdUser = balancePoint.IdUser
	balancePointResponse.BalancePoints = balancePoint.BalancePoints
	return balancePointResponse
}
