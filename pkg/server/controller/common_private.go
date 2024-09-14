package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type CommonControllerForPrivate interface {
	GetCommons(c *gin.Context)
	CreateCommon(c *gin.Context)
	UpdateCommon(c *gin.Context)
	DeleteCommon(c *gin.Context)
}

type commonControllerForPrivate struct {
	CommonRepository repository.CommonRepository
}

func (commonController commonControllerForPrivate) GetCommons(c *gin.Context) {
	var commonRequest request.CommonRequest
	if err := c.Bind(&commonRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.CommonResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Commons: []response.Common{}})
		return
	}
	res := commonController.CommonRepository.GetCommons()
	c.JSON(http.StatusOK, res)
	return
}


func (commonController commonControllerForPrivate) CreateCommon(c *gin.Context) {
	var commonRequest request.CommonRequest
	if err := c.Bind(&commonRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.CommonResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Commons: []response.Common{}})
		return
	}
	var commonModel model.Commons
	res := commonController.CommonRepository.CreateCommon(commonModel)
	c.JSON(http.StatusOK, res)
	return
}


func (commonController commonControllerForPrivate) UpdateCommon(c *gin.Context) {
	var commonRequest request.CommonRequest
	if err := c.Bind(&commonRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.CommonResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Commons: []response.Common{}})
		return
	}
	var commonModel model.Commons
	res := commonController.CommonRepository.UpdateCommon(commonModel)
	c.JSON(http.StatusOK, res)
	return
}


func (commonController commonControllerForPrivate) DeleteCommon(c *gin.Context) {
	var commonRequest request.CommonRequest
	if err := c.Bind(&commonRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.CommonResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Commons: []response.Common{}})
		return
	}
	var uuid string
	res := commonController.CommonRepository.DeleteCommon(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewCommonControllerForPrivate(commonRepository repository.CommonRepository) CommonControllerForPrivate {
	return &commonControllerForPrivate{ CommonRepository: commonRepository}
}