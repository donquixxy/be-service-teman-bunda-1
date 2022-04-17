package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type BalancePointTx struct {
	Id               string          `gorm:"primaryKey;column:id;"`
	IdBalancePoint   string          `gorm:"column:id_balance_point;"`
	TxType           string          `gorm:"column:tx_type;"`
	TxDate           time.Time       `gorm:"column:tx_date;"`
	TxNominal        decimal.Decimal `gorm:"column:tx_nominal;"`
	LastPointBalance decimal.Decimal `gorm:"column:last_point_balance;"`
	NewPointBalance  decimal.Decimal `gorm:"column:new_point_balance;"`
	Description      string          `gorm:"column:description;"`
	CreatedDate      time.Time       `gorm:"column:created_at;"`
}

func (BalancePointTx) TableName() string {
	return "balance_point_transcation"
}
