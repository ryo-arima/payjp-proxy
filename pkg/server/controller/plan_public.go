package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type PlanControllerForPublic interface {
	GetPlans(c *gin.Context)
}

type planControllerForPublic struct {
	PlanRepository repository.PlanRepository
}

func (planController planControllerForPublic) GetPlans(c *gin.Context) {
	var planRequest request.PlanRequest
	if err := c.Bind(&planRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.PlanResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Plans: []response.Plan{}})
		return
	}
	res := planController.PlanRepository.GetPlans()
	c.JSON(http.StatusOK, res)
	return
}


func NewPlanControllerForPublic(planRepository repository.PlanRepository) PlanControllerForPublic {
	return &planControllerForPublic{ PlanRepository: planRepository}
}