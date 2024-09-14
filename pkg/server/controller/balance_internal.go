package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type BalanceControllerForInternal interface {
	GetBalances(c *gin.Context)
	CreateBalance(c *gin.Context)
	UpdateBalance(c *gin.Context)
	DeleteBalance(c *gin.Context)
}

type balanceControllerForInternal struct {
	BalanceRepository repository.BalanceRepository
}

func (balanceController balanceControllerForInternal) GetBalances(c *gin.Context) {
	var balanceRequest request.BalanceRequest
	if err := c.Bind(&balanceRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.BalanceResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Balances: []response.Balance{}})
		return
	}
	res := balanceController.BalanceRepository.GetBalances()
	c.JSON(http.StatusOK, res)
	return
}


func (balanceController balanceControllerForInternal) CreateBalance(c *gin.Context) {
	var balanceRequest request.BalanceRequest
	if err := c.Bind(&balanceRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.BalanceResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Balances: []response.Balance{}})
		return
	}
	var balanceModel model.Balances
	res := balanceController.BalanceRepository.CreateBalance(balanceModel)
	c.JSON(http.StatusOK, res)
	return
}


func (balanceController balanceControllerForInternal) UpdateBalance(c *gin.Context) {
	var balanceRequest request.BalanceRequest
	if err := c.Bind(&balanceRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.BalanceResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Balances: []response.Balance{}})
		return
	}
	var balanceModel model.Balances
	res := balanceController.BalanceRepository.UpdateBalance(balanceModel)
	c.JSON(http.StatusOK, res)
	return
}


func (balanceController balanceControllerForInternal) DeleteBalance(c *gin.Context) {
	var balanceRequest request.BalanceRequest
	if err := c.Bind(&balanceRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.BalanceResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Balances: []response.Balance{}})
		return
	}
	var uuid string
	res := balanceController.BalanceRepository.DeleteBalance(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewBalanceControllerForInternal(balanceRepository repository.BalanceRepository) BalanceControllerForInternal {
	return &balanceControllerForInternal{ BalanceRepository: balanceRepository}
}