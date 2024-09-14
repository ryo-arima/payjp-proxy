package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type TermControllerForPublic interface {
	GetTerms(c *gin.Context)
}

type termControllerForPublic struct {
	TermRepository repository.TermRepository
}

func (termController termControllerForPublic) GetTerms(c *gin.Context) {
	var termRequest request.TermRequest
	if err := c.Bind(&termRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.TermResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Terms: []response.Term{}})
		return
	}
	res := termController.TermRepository.GetTerms()
	c.JSON(http.StatusOK, res)
	return
}


func NewTermControllerForPublic(termRepository repository.TermRepository) TermControllerForPublic {
	return &termControllerForPublic{ TermRepository: termRepository}
}