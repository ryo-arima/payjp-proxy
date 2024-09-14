package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type TransferControllerForInternal interface {
	GetTransfers(c *gin.Context)
	CreateTransfer(c *gin.Context)
	UpdateTransfer(c *gin.Context)
	DeleteTransfer(c *gin.Context)
}

type transferControllerForInternal struct {
	TransferRepository repository.TransferRepository
}

func (transferController transferControllerForInternal) GetTransfers(c *gin.Context) {
	var transferRequest request.TransferRequest
	if err := c.Bind(&transferRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TransferResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Transfers: []response.Transfer{}})
		return
	}
	res := transferController.TransferRepository.GetTransfers()
	c.JSON(http.StatusOK, res)
	return
}

func (transferController transferControllerForInternal) CreateTransfer(c *gin.Context) {
	var transferRequest request.TransferRequest
	if err := c.Bind(&transferRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TransferResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Transfers: []response.Transfer{}})
		return
	}
	var transferModel model.Transfers
	res := transferController.TransferRepository.CreateTransfer(transferModel)
	c.JSON(http.StatusOK, res)
	return
}

func (transferController transferControllerForInternal) UpdateTransfer(c *gin.Context) {
	var transferRequest request.TransferRequest
	if err := c.Bind(&transferRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TransferResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Transfers: []response.Transfer{}})
		return
	}
	var transferModel model.Transfers
	res := transferController.TransferRepository.UpdateTransfer(transferModel)
	c.JSON(http.StatusOK, res)
	return
}

func (transferController transferControllerForInternal) DeleteTransfer(c *gin.Context) {
	var transferRequest request.TransferRequest
	if err := c.Bind(&transferRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TransferResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Transfers: []response.Transfer{}})
		return
	}
	var transferModel model.Transfers
	res := transferController.TransferRepository.DeleteTransfer(transferModel)
	c.JSON(http.StatusOK, res)
	return
}

func NewTransferControllerForInternal(transferRepository repository.TransferRepository) TransferControllerForInternal {
	return &transferControllerForInternal{TransferRepository: transferRepository}
}
