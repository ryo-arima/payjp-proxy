package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type PlanUsecase interface {
	BootstrapPlanForDB(request request.PlanRequest)
	GetPlanForPublic(request request.PlanRequest)
	GetPlanForInternal(request request.PlanRequest)
	GetPlanForPrivate(request request.PlanRequest)
	CreatePlanForInternal(request request.PlanRequest)
	CreatePlanForPrivate(request request.PlanRequest)
	UpdatePlanForInternal(request request.PlanRequest)
	UpdatePlanForPrivate(request request.PlanRequest)
	DeletePlanForInternal(request request.PlanRequest)
	DeletePlanForPrivate(request request.PlanRequest)
}

type planUsecase struct {
	PlanRepository repository.PlanRepository
}

// Bootstrap
func (planUsecase planUsecase) BootstrapPlanForDB(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.BootstrapPlanForDB(request)
	fmt.Println(plans)
}

// GET
func (planUsecase planUsecase) GetPlanForPublic(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.GetPlanForPublic(request)
	fmt.Println(plans)
}

func (planUsecase planUsecase) GetPlanForInternal(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.GetPlanForInternal(request)
	fmt.Println(plans)
}

func (planUsecase planUsecase) GetPlanForPrivate(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.GetPlanForPrivate(request)
	fmt.Println(plans)
}

// CREATE
func (planUsecase planUsecase) CreatePlanForPublic(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.CreatePlanForPublic(request)
	fmt.Println(plans)
}

func (planUsecase planUsecase) CreatePlanForInternal(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.CreatePlanForInternal(request)
	fmt.Println(plans)
}

func (planUsecase planUsecase) CreatePlanForPrivate(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.CreatePlanForPrivate(request)
	fmt.Println(plans)
}

// UPDATE
func (planUsecase planUsecase) UpdatePlanForPublic(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.UpdatePlanForPublic(request)
	fmt.Println(plans)
}

func (planUsecase planUsecase) UpdatePlanForInternal(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.UpdatePlanForInternal(request)
	fmt.Println(plans)
}

func (planUsecase planUsecase) UpdatePlanForPrivate(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.UpdatePlanForPrivate(request)
	fmt.Println(plans)
}

// DELETE
func (planUsecase planUsecase) DeletePlanForPublic(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.DeletePlanForPublic(request)
	fmt.Println(plans)
}

func (planUsecase planUsecase) DeletePlanForInternal(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.DeletePlanForInternal(request)
	fmt.Println(plans)
}

func (planUsecase planUsecase) DeletePlanForPrivate(request request.PlanRequest) {
	plans := planUsecase.PlanRepository.DeletePlanForPrivate(request)
	fmt.Println(plans)
}

func NewPlanUsecase(planRepository repository.PlanRepository) PlanUsecase {
	return &planUsecase{PlanRepository: planRepository}
}
