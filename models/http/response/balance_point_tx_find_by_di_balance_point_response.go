package response

import (
	"time"

	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type FindBalancePointTxByIdBalancePoint struct {
	Id               string    `json:"id"`
	NoOrder          string    `json:"no_order"`
	TxType           string    `json:"tx_type"`
	TxDate           time.Time `json:"tx_date"`
	TxNominal        float64   `json:"tx_nominal"`
	LastPointBalance float64   `json:"last_point_balance"`
	NewPointBalance  float64   `json:"new_point_balance"`
	Description      string    `json:"description"`
}

func ToFindBalancePointTxByIdBalancePoint(balancePointTxs []entity.BalancePointTx) (balancePointWithTxResponses []FindBalancePointTxByIdBalancePoint) {
	for _, balancePointTx := range balancePointTxs {
		var balancePointTxResponse FindBalancePointTxByIdBalancePoint
		balancePointTxResponse.Id = balancePointTx.Id
		balancePointTxResponse.NoOrder = balancePointTx.NoOrder
		balancePointTxResponse.TxType = balancePointTx.TxType
		balancePointTxResponse.TxDate = balancePointTx.TxDate
		balancePointTxResponse.TxNominal = balancePointTx.TxNominal
		balancePointTxResponse.LastPointBalance = balancePointTx.LastPointBalance
		balancePointTxResponse.NewPointBalance = balancePointTx.NewPointBalance
		balancePointTxResponse.Description = balancePointTx.Description
		balancePointWithTxResponses = append(balancePointWithTxResponses, balancePointTxResponse)
	}

	return balancePointWithTxResponses
}
