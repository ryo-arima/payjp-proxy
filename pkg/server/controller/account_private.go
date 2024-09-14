package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type AccountControllerForPrivate interface {
	GetAccounts(c *gin.Context)
	CreateAccount(c *gin.Context)
	UpdateAccount(c *gin.Context)
	DeleteAccount(c *gin.Context)
}

type accountControllerForPrivate struct {
	AccountRepository repository.AccountRepository
}

func (accountController accountControllerForPrivate) GetAccounts(c *gin.Context) {
	var accountRequest request.AccountRequest
	if err := c.Bind(&accountRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.AccountResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Accounts: []response.Account{}})
		return
	}
	res := accountController.AccountRepository.GetAccounts()
	c.JSON(http.StatusOK, res)
	return
}


func (accountController accountControllerForPrivate) CreateAccount(c *gin.Context) {
	var accountRequest request.AccountRequest
	if err := c.Bind(&accountRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.AccountResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Accounts: []response.Account{}})
		return
	}
	var accountModel model.Accounts
	res := accountController.AccountRepository.CreateAccount(accountModel)
	c.JSON(http.StatusOK, res)
	return
}


func (accountController accountControllerForPrivate) UpdateAccount(c *gin.Context) {
	var accountRequest request.AccountRequest
	if err := c.Bind(&accountRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.AccountResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Accounts: []response.Account{}})
		return
	}
	var accountModel model.Accounts
	res := accountController.AccountRepository.UpdateAccount(accountModel)
	c.JSON(http.StatusOK, res)
	return
}


func (accountController accountControllerForPrivate) DeleteAccount(c *gin.Context) {
	var accountRequest request.AccountRequest
	if err := c.Bind(&accountRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.AccountResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Accounts: []response.Account{}})
		return
	}
	var uuid string
	res := accountController.AccountRepository.DeleteAccount(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewAccountControllerForPrivate(accountRepository repository.AccountRepository) AccountControllerForPrivate {
	return &accountControllerForPrivate{ AccountRepository: accountRepository}
}