package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type CustomerRepository interface {
	BootstrapCustomerForDB(request request.CustomerRequest)(response response.CustomerResponse) 
	GetCustomerForPublic(request request.CustomerRequest)(response response.CustomerResponse) 
	GetCustomerForInternal(request request.CustomerRequest)(response response.CustomerResponse) 
	GetCustomerForPrivate(request request.CustomerRequest)(response response.CustomerResponse) 
	CreateCustomerForPublic(request request.CustomerRequest)   (response response.CustomerResponse) 
	CreateCustomerForInternal(request request.CustomerRequest) (response response.CustomerResponse)
	CreateCustomerForPrivate(request request.CustomerRequest)  (response response.CustomerResponse)
	UpdateCustomerForPublic(request request.CustomerRequest)   (response response.CustomerResponse)
	UpdateCustomerForInternal(request request.CustomerRequest) (response response.CustomerResponse)
	UpdateCustomerForPrivate(request request.CustomerRequest)  (response response.CustomerResponse)
	DeleteCustomerForPublic(request request.CustomerRequest)   (response response.CustomerResponse)
	DeleteCustomerForInternal(request request.CustomerRequest) (response response.CustomerResponse)
	DeleteCustomerForPrivate(request request.CustomerRequest)  (response response.CustomerResponse)
}

type customerRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (customerRepository customerRepository) BootstrapCustomerForDB(request request.CustomerRequest) (response response.CustomerResponse) {
	fmt.Println("BootstrapCustomerForDB")
	return response
}

//GET
func (customerRepository customerRepository) GetCustomerForPublic(request request.CustomerRequest) (response response.CustomerResponse) {
	fmt.Println("GetCustomerForPublic")
	return response
}

func (customerRepository customerRepository) GetCustomerForInternal(request request.CustomerRequest) (response response.CustomerResponse ){
	fmt.Println("GetCustomerForInternal")
	return response
}

func (customerRepository customerRepository) GetCustomerForPrivate(request request.CustomerRequest) (response response.CustomerResponse) {
	fmt.Println("GetCustomerForPrivate")
	return response
}

//CREATE
func (customerRepository customerRepository) CreateCustomerForPublic(request request.CustomerRequest) (response response.CustomerResponse ){
	fmt.Println("CreateCustomerForPublic")
	return response
}

func (customerRepository customerRepository) CreateCustomerForInternal(request request.CustomerRequest) (response response.CustomerResponse) {
	fmt.Println("CreateCustomerForInternal()")
	return response
}

func (customerRepository customerRepository) CreateCustomerForPrivate(request request.CustomerRequest) (response response.CustomerResponse){
	fmt.Println("CreateCustomerForPrivate()")
	return response
}

//UPDATE
func (customerRepository customerRepository) UpdateCustomerForPublic(request request.CustomerRequest) (response response.CustomerResponse){
	fmt.Println("UpdateCustomerForPublic()")
	return response
}

func (customerRepository customerRepository) UpdateCustomerForInternal(request request.CustomerRequest) (response response.CustomerResponse){
	fmt.Println("UpdateCustomerForInternal")
	return response
}

func (customerRepository customerRepository) UpdateCustomerForPrivate(request request.CustomerRequest) (response response.CustomerResponse){
	fmt.Println("UpdateCustomerForPrivate")
	return response
}

//DELETE
func (customerRepository customerRepository) DeleteCustomerForPublic(request request.CustomerRequest) (response response.CustomerResponse){
	fmt.Println("DeleteCustomerForPublic")
	return response
}

func (customerRepository customerRepository) DeleteCustomerForInternal(request request.CustomerRequest) (response response.CustomerResponse ){
	fmt.Println("DeleteCustomerForInternal")
	return response
}

func (customerRepository customerRepository) DeleteCustomerForPrivate(request request.CustomerRequest) (response response.CustomerResponse){
	fmt.Println("DeleteCustomerForPrivate")
	return response
}

func NewCustomerRepository(conf config.BaseConfig) CustomerRepository {
	return &customerRepository{BaseConfig: conf}
}