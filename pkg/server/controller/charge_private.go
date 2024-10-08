package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type ChargeControllerForPrivate interface {
	GetCharges(c *gin.Context)
	CreateCharge(c *gin.Context)
	UpdateCharge(c *gin.Context)
	DeleteCharge(c *gin.Context)
}

type chargeControllerForPrivate struct {
	ChargeRepository repository.ChargeRepository
}

func (chargeController chargeControllerForPrivate) GetCharges(c *gin.Context) {
	var chargeRequest request.ChargeRequest
	if err := c.Bind(&chargeRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ChargeResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Charges: []response.Charge{}})
		return
	}
	chargeController.ChargeRepository.GetCharges()
	c.JSON(http.StatusOK, &response.ChargeResponse{Code: "SERVER_CONTROLLER_GET__FOR__002", Message: "Charge retrieved", Charges: []response.Charge{}})
	return
}

func (chargeController chargeControllerForPrivate) CreateCharge(c *gin.Context) {
	var chargeRequest request.ChargeRequest
	if err := c.Bind(&chargeRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ChargeResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Charges: []response.Charge{}})
		return
	}
	var chargeModel model.Charges
	chargeController.ChargeRepository.CreateCharge(chargeModel)
	c.JSON(http.StatusOK, &response.ChargeResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: "Charge created successfully", Charges: []response.Charge{}})
	return
}

func (chargeController chargeControllerForPrivate) UpdateCharge(c *gin.Context) {
	var chargeRequest request.ChargeRequest
	if err := c.Bind(&chargeRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ChargeResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Charges: []response.Charge{}})
		return
	}
	var chargeModel model.Charges
	chargeController.ChargeRepository.UpdateCharge(chargeModel)
	c.JSON(http.StatusOK, &response.ChargeResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: "Charge updated successfully", Charges: []response.Charge{}})
	return
}

func (chargeController chargeControllerForPrivate) DeleteCharge(c *gin.Context) {
	var chargeRequest request.ChargeRequest
	if err := c.Bind(&chargeRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ChargeResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Charges: []response.Charge{}})
		return
	}
	var chargeModel model.Charges
	chargeController.ChargeRepository.DeleteCharge(chargeModel)
	c.JSON(http.StatusOK, &response.ChargeResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: "Charge deleted successfully", Charges: []response.Charge{}})
	return
}

func NewChargeControllerForPrivate(chargeRepository repository.ChargeRepository) ChargeControllerForPrivate {
	return &chargeControllerForPrivate{ChargeRepository: chargeRepository}
}
