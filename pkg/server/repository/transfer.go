package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type TransferRepository interface {
	GetTransfers() []model.Transfers
	CreateTransfer(transfer model.Transfers) *gorm.DB
	UpdateTransfer(transfer model.Transfers) *gorm.DB
	DeleteTransfer(uuid string) *gorm.DB
}

type transferRepository struct {
	BaseConfig config.BaseConfig
}


func (transferRepository transferRepository) GetTransfers() []model.Transfers {
	var transfers []model.Transfers
	transferRepository.BaseConfig.DBConnection.Find(&transfers)
	return transfers
}

func (transferRepository transferRepository) CreateTransfer(transfer model.Transfers) *gorm.DB {
	results := transferRepository.BaseConfig.DBConnection.Create(&transfer)
	return results
}

func (transferRepository transferRepository) UpdateTransfer(transfer model.Transfers) *gorm.DB {
	results := transferRepository.BaseConfig.DBConnection.Model(model.Transfers{}).Where("uuid = ?", transfer.UUID).Updates(transfer)
	return results
}

func (transferRepository transferRepository) DeleteTransfer(uuid string) *gorm.DB {
	results := transferRepository.BaseConfig.DBConnection.Model(model.Transfers{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewTransferRepository(conf config.BaseConfig) TransferRepository {
	return &transferRepository{BaseConfig: conf}
}