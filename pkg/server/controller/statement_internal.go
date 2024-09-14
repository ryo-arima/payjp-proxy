package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type StatementControllerForInternal interface {
	GetStatements(c *gin.Context)
	CreateStatement(c *gin.Context)
	UpdateStatement(c *gin.Context)
	DeleteStatement(c *gin.Context)
}

type statementControllerForInternal struct {
	StatementRepository repository.StatementRepository
}

func (statementController statementControllerForInternal) GetStatements(c *gin.Context) {
	var statementRequest request.StatementRequest
	if err := c.Bind(&statementRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.StatementResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Statements: []response.Statement{}})
		return
	}
	res := statementController.StatementRepository.GetStatements()
	c.JSON(http.StatusOK, res)
	return
}

func (statementController statementControllerForInternal) CreateStatement(c *gin.Context) {
	var statementRequest request.StatementRequest
	if err := c.Bind(&statementRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.StatementResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Statements: []response.Statement{}})
		return
	}
	var statementModel model.Statements
	res := statementController.StatementRepository.CreateStatement(statementModel)
	c.JSON(http.StatusOK, res)
	return
}

func (statementController statementControllerForInternal) UpdateStatement(c *gin.Context) {
	var statementRequest request.StatementRequest
	if err := c.Bind(&statementRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.StatementResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Statements: []response.Statement{}})
		return
	}
	var statementModel model.Statements
	res := statementController.StatementRepository.UpdateStatement(statementModel)
	c.JSON(http.StatusOK, res)
	return
}

func (statementController statementControllerForInternal) DeleteStatement(c *gin.Context) {
	var statementRequest request.StatementRequest
	if err := c.Bind(&statementRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.StatementResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Statements: []response.Statement{}})
		return
	}
	var statementModel model.Statements
	res := statementController.StatementRepository.DeleteStatement(statementModel)
	c.JSON(http.StatusOK, res)
	return
}

func NewStatementControllerForInternal(statementRepository repository.StatementRepository) StatementControllerForInternal {
	return &statementControllerForInternal{StatementRepository: statementRepository}
}
