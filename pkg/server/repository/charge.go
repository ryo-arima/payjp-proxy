package repository

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type ChargeRepository interface {
	GetCharges() []model.Charges
	CreateCharge(charge model.Charges) model.Charges
	UpdateCharge(charge model.Charges) model.Charges
	DeleteCharge(charge model.Charges) model.Charges
}

type chargeRepository struct {
	BaseConfig config.BaseConfig
}

func (chargeRepository chargeRepository) GetCharges() []model.Charges {
	var charges []model.Charges
	return charges
}

func (chargeRepository chargeRepository) CreateCharge(charge model.Charges) model.Charges {
	return charge
}

func (chargeRepository chargeRepository) UpdateCharge(charge model.Charges) model.Charges {
	return charge
}

func (chargeRepository chargeRepository) DeleteCharge(charge model.Charges) model.Charges {
	return charge
}

func NewChargeRepository(conf config.BaseConfig) ChargeRepository {
	return &chargeRepository{BaseConfig: conf}
}
