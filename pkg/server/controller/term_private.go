package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type TermControllerForPrivate interface {
	GetTerms(c *gin.Context)
	CreateTerm(c *gin.Context)
	UpdateTerm(c *gin.Context)
	DeleteTerm(c *gin.Context)
}

type termControllerForPrivate struct {
	TermRepository repository.TermRepository
}

func (termController termControllerForPrivate) GetTerms(c *gin.Context) {
	var termRequest request.TermRequest
	if err := c.Bind(&termRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TermResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Terms: []response.Term{}})
		return
	}
	res := termController.TermRepository.GetTerms()
	c.JSON(http.StatusOK, res)
	return
}

func (termController termControllerForPrivate) CreateTerm(c *gin.Context) {
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

func (termController termControllerForPrivate) UpdateTerm(c *gin.Context) {
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

func (termController termControllerForPrivate) DeleteTerm(c *gin.Context) {
	var termRequest request.TermRequest
	if err := c.Bind(&termRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TermResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Terms: []response.Term{}})
		return
	}
	var termModel model.Terms
	res := termController.TermRepository.DeleteTerm(termModel)
	c.JSON(http.StatusOK, res)
	return
}

func NewTermControllerForPrivate(termRepository repository.TermRepository) TermControllerForPrivate {
	return &termControllerForPrivate{TermRepository: termRepository}
}
