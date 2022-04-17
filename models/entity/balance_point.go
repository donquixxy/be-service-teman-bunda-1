package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type BalancePoint struct {
	Id              string           `gorm:"primaryKey;column:id;"`
	IdUser          string           `gorm:"column:id_user;"`
	BalancePoints   decimal.Decimal  `gorm:"column:balance_points;"`
	BalancePointTxs []BalancePointTx `gorm:"foreignKey:IdBalancePoint;"`
	CreatedDate     time.Time        `gorm:"column:created_at;"`
}

func (BalancePoint) TableName() string {
	return "balance_point"
}
