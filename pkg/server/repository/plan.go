package repository

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type PlanRepository interface {
	GetPlans() []model.Plan
	CreatePlan(plan model.Plan) model.Plan
	UpdatePlan(plan model.Plan) model.Plan
	DeletePlan(plan model.Plan) model.Plan
}

type planRepository struct {
	BaseConfig config.BaseConfig
}

func (planRepository planRepository) GetPlans() []model.Plan {
	var plan []model.Plan
	return plan
}

func (planRepository planRepository) CreatePlan(plan model.Plan) model.Plan {
	return plan
}

func (planRepository planRepository) UpdatePlan(plan model.Plan) model.Plan {
	return plan
}

func (planRepository planRepository) DeletePlan(plan model.Plan) model.Plan {
	return plan
}

func NewPlanRepository(conf config.BaseConfig) PlanRepository {
	return &planRepository{BaseConfig: conf}
}
