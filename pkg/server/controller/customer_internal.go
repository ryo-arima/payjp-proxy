package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type CustomerControllerForInternal interface {
	GetCustomers(c *gin.Context)
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

type customerControllerForInternal struct {
	CustomerRepository repository.CustomerRepository
}

func (customerController customerControllerForInternal) GetCustomers(c *gin.Context) {
	var customerRequest request.CustomerRequest
	if err := c.Bind(&customerRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.CustomerResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Customers: []response.Customer{}})
		return
	}
	res := customerController.CustomerRepository.GetCustomers()
	c.JSON(http.StatusOK, res)
	return
}


func (customerController customerControllerForInternal) CreateCustomer(c *gin.Context) {
	var customerRequest request.CustomerRequest
	if err := c.Bind(&customerRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.CustomerResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Customers: []response.Customer{}})
		return
	}
	var customerModel model.Customers
	res := customerController.CustomerRepository.CreateCustomer(customerModel)
	c.JSON(http.StatusOK, res)
	return
}


func (customerController customerControllerForInternal) UpdateCustomer(c *gin.Context) {
	var customerRequest request.CustomerRequest
	if err := c.Bind(&customerRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.CustomerResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Customers: []response.Customer{}})
		return
	}
	var customerModel model.Customers
	res := customerController.CustomerRepository.UpdateCustomer(customerModel)
	c.JSON(http.StatusOK, res)
	return
}


func (customerController customerControllerForInternal) DeleteCustomer(c *gin.Context) {
	var customerRequest request.CustomerRequest
	if err := c.Bind(&customerRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.CustomerResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Customers: []response.Customer{}})
		return
	}
	var uuid string
	res := customerController.CustomerRepository.DeleteCustomer(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewCustomerControllerForInternal(customerRepository repository.CustomerRepository) CustomerControllerForInternal {
	return &customerControllerForInternal{ CustomerRepository: customerRepository}
}