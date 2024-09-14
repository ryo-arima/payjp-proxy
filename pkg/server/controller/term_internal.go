package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type TermControllerForInternal interface {
	GetTerms(c *gin.Context)
	CreateTerm(c *gin.Context)
	UpdateTerm(c *gin.Context)
	DeleteTerm(c *gin.Context)
}

type termControllerForInternal struct {
	TermRepository repository.TermRepository
}

func (termController termControllerForInternal) GetTerms(c *gin.Context) {
	var termRequest request.TermRequest
	if err := c.Bind(&termRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TermResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Terms: []response.Term{}})
		return
	}
	res := termController.TermRepository.GetTerms()
	c.JSON(http.StatusOK, res)
	return
}


func (termController termControllerForInternal) CreateTerm(c *gin.Context) {
	var termRequest request.TermRequest
	if err := c.Bind(&termRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TermResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Terms: []response.Term{}})
		return
	}
	var termModel model.Terms
	res := termController.TermRepository.CreateTerm(termModel)
	c.JSON(http.StatusOK, res)
	return
}


func (termController termControllerForInternal) UpdateTerm(c *gin.Context) {
	var termRequest request.TermRequest
	if err := c.Bind(&termRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TermResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Terms: []response.Term{}})
		return
	}
	var termModel model.Terms
	res := termController.TermRepository.UpdateTerm(termModel)
	c.JSON(http.StatusOK, res)
	return
}


func (termController termControllerForInternal) DeleteTerm(c *gin.Context) {
	var termRequest request.TermRequest
	if err := c.Bind(&termRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TermResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Terms: []response.Term{}})
		return
	}
	var uuid string
	res := termController.TermRepository.DeleteTerm(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewTermControllerForInternal(termRepository repository.TermRepository) TermControllerForInternal {
	return &termControllerForInternal{ TermRepository: termRepository}
}