package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type ChargeControllerForInternal interface {
	GetCharges(c *gin.Context)
	CreateCharge(c *gin.Context)
	UpdateCharge(c *gin.Context)
	DeleteCharge(c *gin.Context)
}

type chargeControllerForInternal struct {
	ChargeRepository repository.ChargeRepository
}

func (chargeController chargeControllerForInternal) GetCharges(c *gin.Context) {
	var chargeRequest request.ChargeRequest
	if err := c.Bind(&chargeRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ChargeResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Charges: []response.Charge{}})
		return
	}
	res := chargeController.ChargeRepository.GetCharges()
	c.JSON(http.StatusOK, res)
	return
}


func (chargeController chargeControllerForInternal) CreateCharge(c *gin.Context) {
	var chargeRequest request.ChargeRequest
	if err := c.Bind(&chargeRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ChargeResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Charges: []response.Charge{}})
		return
	}
	var chargeModel model.Charges
	res := chargeController.ChargeRepository.CreateCharge(chargeModel)
	c.JSON(http.StatusOK, res)
	return
}


func (chargeController chargeControllerForInternal) UpdateCharge(c *gin.Context) {
	var chargeRequest request.ChargeRequest
	if err := c.Bind(&chargeRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ChargeResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Charges: []response.Charge{}})
		return
	}
	var chargeModel model.Charges
	res := chargeController.ChargeRepository.UpdateCharge(chargeModel)
	c.JSON(http.StatusOK, res)
	return
}


func (chargeController chargeControllerForInternal) DeleteCharge(c *gin.Context) {
	var chargeRequest request.ChargeRequest
	if err := c.Bind(&chargeRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ChargeResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Charges: []response.Charge{}})
		return
	}
	var uuid string
	res := chargeController.ChargeRepository.DeleteCharge(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewChargeControllerForInternal(chargeRepository repository.ChargeRepository) ChargeControllerForInternal {
	return &chargeControllerForInternal{ ChargeRepository: chargeRepository}
}