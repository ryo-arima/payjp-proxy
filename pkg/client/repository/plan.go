package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type PlanRepository interface {
	BootstrapPlanForDB(request request.PlanRequest)(response response.PlanResponse) 
	GetPlanForPublic(request request.PlanRequest)(response response.PlanResponse) 
	GetPlanForInternal(request request.PlanRequest)(response response.PlanResponse) 
	GetPlanForPrivate(request request.PlanRequest)(response response.PlanResponse) 
	CreatePlanForPublic(request request.PlanRequest)   (response response.PlanResponse) 
	CreatePlanForInternal(request request.PlanRequest) (response response.PlanResponse)
	CreatePlanForPrivate(request request.PlanRequest)  (response response.PlanResponse)
	UpdatePlanForPublic(request request.PlanRequest)   (response response.PlanResponse)
	UpdatePlanForInternal(request request.PlanRequest) (response response.PlanResponse)
	UpdatePlanForPrivate(request request.PlanRequest)  (response response.PlanResponse)
	DeletePlanForPublic(request request.PlanRequest)   (response response.PlanResponse)
	DeletePlanForInternal(request request.PlanRequest) (response response.PlanResponse)
	DeletePlanForPrivate(request request.PlanRequest)  (response response.PlanResponse)
}

type planRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (planRepository planRepository) BootstrapPlanForDB(request request.PlanRequest) (response response.PlanResponse) {
	fmt.Println("BootstrapPlanForDB")
	return response
}

//GET
func (planRepository planRepository) GetPlanForPublic(request request.PlanRequest) (response response.PlanResponse) {
	fmt.Println("GetPlanForPublic")
	return response
}

func (planRepository planRepository) GetPlanForInternal(request request.PlanRequest) (response response.PlanResponse ){
	fmt.Println("GetPlanForInternal")
	return response
}

func (planRepository planRepository) GetPlanForPrivate(request request.PlanRequest) (response response.PlanResponse) {
	fmt.Println("GetPlanForPrivate")
	return response
}

//CREATE
func (planRepository planRepository) CreatePlanForPublic(request request.PlanRequest) (response response.PlanResponse ){
	fmt.Println("CreatePlanForPublic")
	return response
}

func (planRepository planRepository) CreatePlanForInternal(request request.PlanRequest) (response response.PlanResponse) {
	fmt.Println("CreatePlanForInternal()")
	return response
}

func (planRepository planRepository) CreatePlanForPrivate(request request.PlanRequest) (response response.PlanResponse){
	fmt.Println("CreatePlanForPrivate()")
	return response
}

//UPDATE
func (planRepository planRepository) UpdatePlanForPublic(request request.PlanRequest) (response response.PlanResponse){
	fmt.Println("UpdatePlanForPublic()")
	return response
}

func (planRepository planRepository) UpdatePlanForInternal(request request.PlanRequest) (response response.PlanResponse){
	fmt.Println("UpdatePlanForInternal")
	return response
}

func (planRepository planRepository) UpdatePlanForPrivate(request request.PlanRequest) (response response.PlanResponse){
	fmt.Println("UpdatePlanForPrivate")
	return response
}

//DELETE
func (planRepository planRepository) DeletePlanForPublic(request request.PlanRequest) (response response.PlanResponse){
	fmt.Println("DeletePlanForPublic")
	return response
}

func (planRepository planRepository) DeletePlanForInternal(request request.PlanRequest) (response response.PlanResponse ){
	fmt.Println("DeletePlanForInternal")
	return response
}

func (planRepository planRepository) DeletePlanForPrivate(request request.PlanRequest) (response response.PlanResponse){
	fmt.Println("DeletePlanForPrivate")
	return response
}

func NewPlanRepository(conf config.BaseConfig) PlanRepository {
	return &planRepository{BaseConfig: conf}
}