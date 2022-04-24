package mysql

import (
	"strings"

	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type BalancePointTxRepositoryInterface interface {
	CreateBalancePointTx(DB *gorm.DB, balancePoint entity.BalancePointTx) (entity.BalancePointTx, error)
	FindBalancePointTxByIdBalancePoint(DB *gorm.DB, date string, idBalancePoint string) ([]entity.BalancePointTx, error)
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

func (repository *BalancePointTxRepositoryImplementation) FindBalancePointTxByIdBalancePoint(DB *gorm.DB, date string, idBalancePoint string) ([]entity.BalancePointTx, error) {
	var balancePointTx []entity.BalancePointTx
	var dateStart = []string{date, "00:00:00"}
	var dateEnd = []string{date, "23:59:59"}

	results := DB.Where("id_balance_point = ?", idBalancePoint).Where("tx_date >= ?", strings.Join(dateStart, " ")).Where("tx_date <= ?", strings.Join(dateEnd, " ")).Find(&balancePointTx)
	return balancePointTx, results.Error
}
