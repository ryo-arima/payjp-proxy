package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type AccountRepository interface {
	GetAccounts() []model.Accounts
	CreateAccount(account model.Accounts)
	UpdateAccount(account model.Accounts)
	DeleteAccount(account model.Accounts)
}

type accountRepository struct {
	BaseConfig config.BaseConfig
}

func (accountRepository accountRepository) GetAccounts() []model.Accounts {
	var accounts []model.Accounts
	return accounts
}

func (accountRepository accountRepository) CreateAccount(account model.Accounts) {
	fmt.Println("account")
}

func (accountRepository accountRepository) UpdateAccount(account model.Accounts) {
	fmt.Println("account")
}

func (accountRepository accountRepository) DeleteAccount(account model.Accounts) {
	fmt.Println("account")
}

func NewAccountRepository(conf config.BaseConfig) AccountRepository {
	return &accountRepository{BaseConfig: conf}
}
