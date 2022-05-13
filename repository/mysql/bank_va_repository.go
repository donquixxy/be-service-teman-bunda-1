package mysql

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/config"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
	"gorm.io/gorm"
)

type BankVaRepositoryInterface interface {
	FindBankVaByBankCode(DB *gorm.DB, bankCode string) (entity.BankVa, error)
	FindAllBankVa(DB *gorm.DB) ([]entity.BankVa, error)
}

type BankVaRepositoryImplementation struct {
	configurationDatabase *config.Database
}

func NewBankVaRepository(configDatabase *config.Database) BankVaRepositoryInterface {
	return &BankVaRepositoryImplementation{
		configurationDatabase: configDatabase,
	}
}

func (repository *BankVaRepositoryImplementation) FindBankVaByBankCode(DB *gorm.DB, bankCode string) (entity.BankVa, error) {
	var bankVa entity.BankVa
	results := DB.Where("bank_code = ?", bankCode).Where("is_active = ?", "1").First(&bankVa)
	return bankVa, results.Error
}

func (repository *BankVaRepositoryImplementation) FindAllBankVa(DB *gorm.DB) ([]entity.BankVa, error) {
	var bankVas []entity.BankVa
	results := DB.Where("bank_code != ?", "qris").Where("bank_code != ?", "cod").Where("is_active = ?", "1").Find(&bankVas)
	return bankVas, results.Error
}
