package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type PlanRepository interface {
	GetPlans() []model.Plans
	CreatePlan(plan model.Plans) *gorm.DB
	UpdatePlan(plan model.Plans) *gorm.DB
	DeletePlan(uuid string) *gorm.DB
}

type planRepository struct {
	BaseConfig config.BaseConfig
}


func (planRepository planRepository) GetPlans() []model.Plans {
	var plans []model.Plans
	planRepository.BaseConfig.DBConnection.Find(&plans)
	return plans
}

func (planRepository planRepository) CreatePlan(plan model.Plans) *gorm.DB {
	results := planRepository.BaseConfig.DBConnection.Create(&plan)
	return results
}

func (planRepository planRepository) UpdatePlan(plan model.Plans) *gorm.DB {
	results := planRepository.BaseConfig.DBConnection.Model(model.Plans{}).Where("uuid = ?", plan.UUID).Updates(plan)
	return results
}

func (planRepository planRepository) DeletePlan(uuid string) *gorm.DB {
	results := planRepository.BaseConfig.DBConnection.Model(model.Plans{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewPlanRepository(conf config.BaseConfig) PlanRepository {
	return &planRepository{BaseConfig: conf}
}