package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccounts() []model.Accounts
	CreateAccount(account model.Accounts) *gorm.DB
	UpdateAccount(account model.Accounts) *gorm.DB
	DeleteAccount(uuid string) *gorm.DB
}

type accountRepository struct {
	BaseConfig config.BaseConfig
}


func (accountRepository accountRepository) GetAccounts() []model.Accounts {
	var accounts []model.Accounts
	accountRepository.BaseConfig.DBConnection.Find(&accounts)
	return accounts
}

func (accountRepository accountRepository) CreateAccount(account model.Accounts) *gorm.DB {
	results := accountRepository.BaseConfig.DBConnection.Create(&account)
	return results
}

func (accountRepository accountRepository) UpdateAccount(account model.Accounts) *gorm.DB {
	results := accountRepository.BaseConfig.DBConnection.Model(model.Accounts{}).Where("uuid = ?", account.UUID).Updates(account)
	return results
}

func (accountRepository accountRepository) DeleteAccount(uuid string) *gorm.DB {
	results := accountRepository.BaseConfig.DBConnection.Model(model.Accounts{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewAccountRepository(conf config.BaseConfig) AccountRepository {
	return &accountRepository{BaseConfig: conf}
}