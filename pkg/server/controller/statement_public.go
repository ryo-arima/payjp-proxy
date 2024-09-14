package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type StatementControllerForPublic interface {
	GetStatements(c *gin.Context)
}

type statementControllerForPublic struct {
	StatementRepository repository.StatementRepository
}

func (statementController statementControllerForPublic) GetStatements(c *gin.Context) {
	var statementRequest request.StatementRequest
	if err := c.Bind(&statementRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.StatementResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Statements: []response.Statement{}})
		return
	}
	res := statementController.StatementRepository.GetStatements()
	c.JSON(http.StatusOK, res)
	return
}


func NewStatementControllerForPublic(statementRepository repository.StatementRepository) StatementControllerForPublic {
	return &statementControllerForPublic{ StatementRepository: statementRepository}
}