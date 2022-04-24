package entity

import (
	"time"
)

type BalancePointTx struct {
	Id               string    `gorm:"primaryKey;column:id;"`
	IdBalancePoint   string    `gorm:"column:id_balance_point;"`
	NoOrder          string    `gorm:"column:no_order;"`
	TxType           string    `gorm:"column:tx_type;"`
	TxDate           time.Time `gorm:"column:tx_date;"`
	TxNominal        float64   `gorm:"column:tx_nominal;"`
	LastPointBalance float64   `gorm:"column:last_point_balance;"`
	NewPointBalance  float64   `gorm:"column:new_point_balance;"`
	Description      string    `gorm:"column:description;"`
	CreatedDate      time.Time `gorm:"column:created_at;"`
}

func (BalancePointTx) TableName() string {
	return "balance_point_transcation"
}
