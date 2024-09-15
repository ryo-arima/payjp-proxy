package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type PlanControllerForInternal interface {
	GetPlans(c *gin.Context)
	CreatePlan(c *gin.Context)
	UpdatePlan(c *gin.Context)
	DeletePlan(c *gin.Context)
}

type planControllerForInternal struct {
	PlanRepository repository.PlanRepository
}

func (planController planControllerForInternal) GetPlans(c *gin.Context) {
	var planRequest request.PlanRequest
	if err := c.Bind(&planRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PlanResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Plans: []response.Plan{}})
		return
	}
	res := planController.PlanRepository.GetPlans()
	c.JSON(http.StatusOK, res)
	return
}

func (planController planControllerForInternal) CreatePlan(c *gin.Context) {
	var planRequest request.PlanRequest
	if err := c.Bind(&planRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PlanResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Plans: []response.Plan{}})
		return
	}
	var planModel model.Plan
	res := planController.PlanRepository.CreatePlan(planModel)
	c.JSON(http.StatusOK, res)
	return
}

func (planController planControllerForInternal) UpdatePlan(c *gin.Context) {
	var planRequest request.PlanRequest
	if err := c.Bind(&planRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PlanResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Plans: []response.Plan{}})
		return
	}
	var planModel model.Plan
	res := planController.PlanRepository.UpdatePlan(planModel)
	c.JSON(http.StatusOK, res)
	return
}

func (planController planControllerForInternal) DeletePlan(c *gin.Context) {
	var planRequest request.PlanRequest
	if err := c.Bind(&planRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PlanResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Plans: []response.Plan{}})
		return
	}
	var planModel model.Plan
	res := planController.PlanRepository.DeletePlan(planModel)
	c.JSON(http.StatusOK, res)
	return
}

func NewPlanControllerForInternal(planRepository repository.PlanRepository) PlanControllerForInternal {
	return &planControllerForInternal{PlanRepository: planRepository}
}
