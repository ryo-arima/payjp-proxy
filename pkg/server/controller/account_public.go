package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type AccountControllerForPublic interface {
	GetAccounts(c *gin.Context)
}

type accountControllerForPublic struct {
	AccountRepository repository.AccountRepository
}

func (accountController accountControllerForPublic) GetAccounts(c *gin.Context) {
	var accountRequest request.AccountRequest
	if err := c.Bind(&accountRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.AccountResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Accounts: []response.Account{}})
		return
	}
	res := accountController.AccountRepository.GetAccounts()
	c.JSON(http.StatusOK, res)
	return
}


func NewAccountControllerForPublic(accountRepository repository.AccountRepository) AccountControllerForPublic {
	return &accountControllerForPublic{ AccountRepository: accountRepository}
}