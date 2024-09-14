package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type CommonRepository interface {
	GetCommons() []model.Commons
	CreateCommon(common model.Commons) *gorm.DB
	UpdateCommon(common model.Commons) *gorm.DB
	DeleteCommon(uuid string) *gorm.DB
}

type commonRepository struct {
	BaseConfig config.BaseConfig
}


func (commonRepository commonRepository) GetCommons() []model.Commons {
	var commons []model.Commons
	commonRepository.BaseConfig.DBConnection.Find(&commons)
	return commons
}

func (commonRepository commonRepository) CreateCommon(common model.Commons) *gorm.DB {
	results := commonRepository.BaseConfig.DBConnection.Create(&common)
	return results
}

func (commonRepository commonRepository) UpdateCommon(common model.Commons) *gorm.DB {
	results := commonRepository.BaseConfig.DBConnection.Model(model.Commons{}).Where("uuid = ?", common.UUID).Updates(common)
	return results
}

func (commonRepository commonRepository) DeleteCommon(uuid string) *gorm.DB {
	results := commonRepository.BaseConfig.DBConnection.Model(model.Commons{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewCommonRepository(conf config.BaseConfig) CommonRepository {
	return &commonRepository{BaseConfig: conf}
}