package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type CustomerControllerForPublic interface {
	GetCustomers(c *gin.Context)
}

type customerControllerForPublic struct {
	CustomerRepository repository.CustomerRepository
}

func (customerController customerControllerForPublic) GetCustomers(c *gin.Context) {
	var customerRequest request.CustomerRequest
	if err := c.Bind(&customerRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.CustomerResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Customers: []response.Customer{}})
		return
	}
	res := customerController.CustomerRepository.GetCustomers()
	c.JSON(http.StatusOK, res)
	return
}


func NewCustomerControllerForPublic(customerRepository repository.CustomerRepository) CustomerControllerForPublic {
	return &customerControllerForPublic{ CustomerRepository: customerRepository}
}