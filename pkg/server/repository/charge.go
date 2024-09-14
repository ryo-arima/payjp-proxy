package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type ChargeRepository interface {
	GetCharges() []model.Charges
	CreateCharge(charge model.Charges) *gorm.DB
	UpdateCharge(charge model.Charges) *gorm.DB
	DeleteCharge(uuid string) *gorm.DB
}

type chargeRepository struct {
	BaseConfig config.BaseConfig
}


func (chargeRepository chargeRepository) GetCharges() []model.Charges {
	var charges []model.Charges
	chargeRepository.BaseConfig.DBConnection.Find(&charges)
	return charges
}

func (chargeRepository chargeRepository) CreateCharge(charge model.Charges) *gorm.DB {
	results := chargeRepository.BaseConfig.DBConnection.Create(&charge)
	return results
}

func (chargeRepository chargeRepository) UpdateCharge(charge model.Charges) *gorm.DB {
	results := chargeRepository.BaseConfig.DBConnection.Model(model.Charges{}).Where("uuid = ?", charge.UUID).Updates(charge)
	return results
}

func (chargeRepository chargeRepository) DeleteCharge(uuid string) *gorm.DB {
	results := chargeRepository.BaseConfig.DBConnection.Model(model.Charges{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewChargeRepository(conf config.BaseConfig) ChargeRepository {
	return &chargeRepository{BaseConfig: conf}
}