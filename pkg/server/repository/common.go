package repository

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type CommonRepository interface {
	GetCommons() []model.Commons
}

type commonRepository struct {
	BaseConfig config.BaseConfig
}

func (commonRepository commonRepository) GetCommons() []model.Commons {
	var commons []model.Commons
	return commons
}

func NewCommonRepository(conf config.BaseConfig) CommonRepository {
	return &commonRepository{BaseConfig: conf}
}
