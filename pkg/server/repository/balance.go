package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type BalanceRepository interface {
	GetBalances() []model.Balances
	CreateBalance(balance model.Balances) *gorm.DB
	UpdateBalance(balance model.Balances) *gorm.DB
	DeleteBalance(uuid string) *gorm.DB
}

type balanceRepository struct {
	BaseConfig config.BaseConfig
}


func (balanceRepository balanceRepository) GetBalances() []model.Balances {
	var balances []model.Balances
	balanceRepository.BaseConfig.DBConnection.Find(&balances)
	return balances
}

func (balanceRepository balanceRepository) CreateBalance(balance model.Balances) *gorm.DB {
	results := balanceRepository.BaseConfig.DBConnection.Create(&balance)
	return results
}

func (balanceRepository balanceRepository) UpdateBalance(balance model.Balances) *gorm.DB {
	results := balanceRepository.BaseConfig.DBConnection.Model(model.Balances{}).Where("uuid = ?", balance.UUID).Updates(balance)
	return results
}

func (balanceRepository balanceRepository) DeleteBalance(uuid string) *gorm.DB {
	results := balanceRepository.BaseConfig.DBConnection.Model(model.Balances{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewBalanceRepository(conf config.BaseConfig) BalanceRepository {
	return &balanceRepository{BaseConfig: conf}
}