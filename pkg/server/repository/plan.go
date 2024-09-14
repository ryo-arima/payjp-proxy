package repository

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type PlanRepository interface {
	GetPlans() []model.Plans
	CreatePlan(plan model.Plans) model.Plans
	UpdatePlan(plan model.Plans) model.Plans
	DeletePlan(plan model.Plans) model.Plans
}

type planRepository struct {
	BaseConfig config.BaseConfig
}

func (planRepository planRepository) GetPlans() []model.Plans {
	var plans []model.Plans
	return plans
}

func (planRepository planRepository) CreatePlan(plan model.Plans) model.Plans {
	return plan
}

func (planRepository planRepository) UpdatePlan(plan model.Plans) model.Plans {
	return plan
}

func (planRepository planRepository) DeletePlan(plan model.Plans) model.Plans {
	return plan
}

func NewPlanRepository(conf config.BaseConfig) PlanRepository {
	return &planRepository{BaseConfig: conf}
}
