package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type BalanceRepository interface {
	GetBalances() []model.Balances
	CreateBalance(balance model.Balances)
	UpdateBalance(balance model.Balances)
	DeleteBalance(balance model.Balances)
}

type balanceRepository struct {
	BaseConfig config.BaseConfig
}

func (balanceRepository balanceRepository) GetBalances() []model.Balances {
	var balances []model.Balances
	return balances
}

func (balanceRepository balanceRepository) CreateBalance(balance model.Balances) {
	fmt.Println("balance")
}

func (balanceRepository balanceRepository) UpdateBalance(balance model.Balances) {
	fmt.Println("balance")
}

func (balanceRepository balanceRepository) DeleteBalance(balance model.Balances) {
	fmt.Println("balance")
}

func NewBalanceRepository(conf config.BaseConfig) BalanceRepository {
	return &balanceRepository{BaseConfig: conf}
}
