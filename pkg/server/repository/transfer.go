package repository

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type TransferRepository interface {
	GetTransfers() []model.Transfers
	CreateTransfer(transfer model.Transfers) model.Transfers
	UpdateTransfer(transfer model.Transfers) model.Transfers
	DeleteTransfer(transfer model.Transfers) model.Transfers
}

type transferRepository struct {
	BaseConfig config.BaseConfig
}

func (transferRepository transferRepository) GetTransfers() []model.Transfers {
	var transfers []model.Transfers
	return transfers
}

func (transferRepository transferRepository) CreateTransfer(transfer model.Transfers) model.Transfers {
	return transfer
}

func (transferRepository transferRepository) UpdateTransfer(transfer model.Transfers) model.Transfers {
	return transfer
}

func (transferRepository transferRepository) DeleteTransfer(transfer model.Transfers) model.Transfers {
	return transfer
}

func NewTransferRepository(conf config.BaseConfig) TransferRepository {
	return &transferRepository{BaseConfig: conf}
}
