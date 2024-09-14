package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type CommonControllerForInternal interface {
	GetCommons(c *gin.Context)
}

type commonControllerForInternal struct {
	CommonRepository repository.CommonRepository
}

func (commonController commonControllerForInternal) GetCommons(c *gin.Context) {
	var commonRequest request.CommonRequest
	if err := c.Bind(&commonRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.CommonResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Commons: []response.Common{}})
		return
	}
	commonController.CommonRepository.GetCommons()
	c.JSON(http.StatusOK, &response.CommonResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: "Common retrieved", Commons: []response.Common{}})
	return
}

func NewCommonControllerForInternal(commonRepository repository.CommonRepository) CommonControllerForInternal {
	return &commonControllerForInternal{CommonRepository: commonRepository}
}
