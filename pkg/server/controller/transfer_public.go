package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type TransferControllerForPublic interface {
	GetTransfers(c *gin.Context)
}

type transferControllerForPublic struct {
	TransferRepository repository.TransferRepository
}

func (transferController transferControllerForPublic) GetTransfers(c *gin.Context) {
	var transferRequest request.TransferRequest
	if err := c.Bind(&transferRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TransferResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Transfers: []response.Transfer{}})
		return
	}
	res := transferController.TransferRepository.GetTransfers()
	c.JSON(http.StatusOK, res)
	return
}


func NewTransferControllerForPublic(transferRepository repository.TransferRepository) TransferControllerForPublic {
	return &transferControllerForPublic{ TransferRepository: transferRepository}
}