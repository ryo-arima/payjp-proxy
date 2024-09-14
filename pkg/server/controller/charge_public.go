package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type ChargeControllerForPublic interface {
	GetCharges(c *gin.Context)
}

type chargeControllerForPublic struct {
	ChargeRepository repository.ChargeRepository
}

func (chargeController chargeControllerForPublic) GetCharges(c *gin.Context) {
	var chargeRequest request.ChargeRequest
	if err := c.Bind(&chargeRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ChargeResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Charges: []response.Charge{}})
		return
	}
	res := chargeController.ChargeRepository.GetCharges()
	c.JSON(http.StatusOK, res)
	return
}


func NewChargeControllerForPublic(chargeRepository repository.ChargeRepository) ChargeControllerForPublic {
	return &chargeControllerForPublic{ ChargeRepository: chargeRepository}
}