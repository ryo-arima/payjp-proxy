package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type CommonControllerForPublic interface {
	GetCommons(c *gin.Context)
}

type commonControllerForPublic struct {
	CommonRepository repository.CommonRepository
}

func (commonController commonControllerForPublic) GetCommons(c *gin.Context) {
	var commonRequest request.CommonRequest
	if err := c.Bind(&commonRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.CommonResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Commons: []response.Common{}})
		return
	}
	res := commonController.CommonRepository.GetCommons()
	c.JSON(http.StatusOK, res)
	return
}


func NewCommonControllerForPublic(commonRepository repository.CommonRepository) CommonControllerForPublic {
	return &commonControllerForPublic{ CommonRepository: commonRepository}
}