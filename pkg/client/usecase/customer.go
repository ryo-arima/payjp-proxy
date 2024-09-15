package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type CustomerUsecase interface {
	BootstrapCustomerForDB(request request.CustomerRequest)
	GetCustomerForPublic(request request.CustomerRequest)
	GetCustomerForInternal(request request.CustomerRequest)
	GetCustomerForPrivate(request request.CustomerRequest)
	CreateCustomerForPublic(request request.CustomerRequest)
	CreateCustomerForInternal(request request.CustomerRequest)
	CreateCustomerForPrivate(request request.CustomerRequest)
	UpdateCustomerForPublic(request request.CustomerRequest)
	UpdateCustomerForInternal(request request.CustomerRequest)
	UpdateCustomerForPrivate(request request.CustomerRequest)
	DeleteCustomerForPublic(request request.CustomerRequest)
	DeleteCustomerForInternal(request request.CustomerRequest)
	DeleteCustomerForPrivate(request request.CustomerRequest)
}

type customerUsecase struct {
	CustomerRepository repository.CustomerRepository
}

// Bootstrap
func (customerUsecase customerUsecase) BootstrapCustomerForDB(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.BootstrapCustomerForDB(request)
	fmt.Println(customers)
}

// GET
func (customerUsecase customerUsecase) GetCustomerForPublic(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.GetCustomerForPublic(request)
	fmt.Println(customers)
}

func (customerUsecase customerUsecase) GetCustomerForInternal(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.GetCustomerForInternal(request)
	fmt.Println(customers)
}

func (customerUsecase customerUsecase) GetCustomerForPrivate(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.GetCustomerForPrivate(request)
	fmt.Println(customers)
}

// CREATE
func (customerUsecase customerUsecase) CreateCustomerForPublic(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.CreateCustomerForPublic(request)
	fmt.Println(customers)
}

func (customerUsecase customerUsecase) CreateCustomerForInternal(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.CreateCustomerForInternal(request)
	fmt.Println(customers)
}

func (customerUsecase customerUsecase) CreateCustomerForPrivate(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.CreateCustomerForPrivate(request)
	fmt.Println(customers)
}

// UPDATE
func (customerUsecase customerUsecase) UpdateCustomerForPublic(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.UpdateCustomerForPublic(request)
	fmt.Println(customers)
}

func (customerUsecase customerUsecase) UpdateCustomerForInternal(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.UpdateCustomerForInternal(request)
	fmt.Println(customers)
}

func (customerUsecase customerUsecase) UpdateCustomerForPrivate(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.UpdateCustomerForPrivate(request)
	fmt.Println(customers)
}

// DELETE
func (customerUsecase customerUsecase) DeleteCustomerForPublic(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.DeleteCustomerForPublic(request)
	fmt.Println(customers)
}

func (customerUsecase customerUsecase) DeleteCustomerForInternal(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.DeleteCustomerForInternal(request)
	fmt.Println(customers)
}

func (customerUsecase customerUsecase) DeleteCustomerForPrivate(request request.CustomerRequest) {
	customers := customerUsecase.CustomerRepository.DeleteCustomerForPrivate(request)
	fmt.Println(customers)
}

func NewCustomerUsecase(customerRepository repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{CustomerRepository: customerRepository}
}
