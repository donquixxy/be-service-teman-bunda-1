package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type BalancePointTxRepositoryInterface interface {
	CreateBalancePointTx(DB *gorm.DB, balancePoint entity.BalancePointTx) (entity.BalancePointTx, error)
}

type BalancePointTxRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewBalancePointTxRepository(configDatabase *config.Database) BalancePointTxRepositoryInterface {
	return &BalancePointTxRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *BalancePointTxRepositoryImplementation) CreateBalancePointTx(DB *gorm.DB, balancePointTx entity.BalancePointTx) (entity.BalancePointTx, error) {
	results := DB.Create(balancePointTx)
	return balancePointTx, results.Error
}
