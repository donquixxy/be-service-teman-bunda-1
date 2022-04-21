package response

import (
	"time"

	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type FindBalancePointWithTxByIdUser struct {
	Id             string           `json:"id"`
	IdUser         string           `json:"id_user"`
	BalancePoints  float64          `json:"balance_points"`
	BalancePointTx []BalancePointTx `json:"balance_point_tx"`
}

type BalancePointTx struct {
	Id               string    `json:"id"`
	TxType           string    `json:"tx_type"`
	TxDate           time.Time `json:"tx_date"`
	TxNominal        float64   `json:"tx_nominal"`
	LastPointBalance float64   `json:"last_point_balance"`
	NewPointBalance  float64   `json:"new_point_balance"`
	Description      string    `json:"description"`
}

func ToFindBalancePointWithTxByIdUser(balancePoint entity.BalancePoint) (balancePointWithTxResponse FindBalancePointWithTxByIdUser) {
	balancePointWithTxResponse.Id = balancePoint.Id
	balancePointWithTxResponse.IdUser = balancePoint.IdUser
	balancePointWithTxResponse.BalancePoints = balancePoint.BalancePoints

	var balancePointTxResponses []BalancePointTx
	for _, balancePointTx := range balancePoint.BalancePointTxs {
		var balancePointTxResponse BalancePointTx
		balancePointTxResponse.Id = balancePointTx.Id
		balancePointTxResponse.TxType = balancePointTx.TxType
		balancePointTxResponse.TxDate = balancePointTx.TxDate
		balancePointTxResponse.TxNominal = balancePointTx.TxNominal
		balancePointTxResponse.LastPointBalance = balancePointTx.LastPointBalance
		balancePointTxResponse.NewPointBalance = balancePointTx.NewPointBalance
		balancePointTxResponse.Description = balancePointTx.Id
		balancePointTxResponses = append(balancePointTxResponses, balancePointTxResponse)
	}

	balancePointWithTxResponse.BalancePointTx = balancePointTxResponses

	return balancePointWithTxResponse
}
