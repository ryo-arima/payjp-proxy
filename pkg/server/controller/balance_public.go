package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type BalanceControllerForPublic interface {
	GetBalances(c *gin.Context)
}

type balanceControllerForPublic struct {
	BalanceRepository repository.BalanceRepository
}

func (balanceController balanceControllerForPublic) GetBalances(c *gin.Context) {
	var balanceRequest request.BalanceRequest
	if err := c.Bind(&balanceRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.BalanceResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Balances: []response.Balance{}})
		return
	}
	res := balanceController.BalanceRepository.GetBalances()
	c.JSON(http.StatusOK, res)
	return
}


func NewBalanceControllerForPublic(balanceRepository repository.BalanceRepository) BalanceControllerForPublic {
	return &balanceControllerForPublic{ BalanceRepository: balanceRepository}
}