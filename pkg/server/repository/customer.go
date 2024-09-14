package repository

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type CustomerRepository interface {
	GetCustomers() []model.Customers
	CreateCustomer(customer model.Customers) model.Customers
	UpdateCustomer(customer model.Customers) model.Customers
	DeleteCustomer(customer model.Customers) model.Customers
}

type customerRepository struct {
	BaseConfig config.BaseConfig
}

func (customerRepository customerRepository) GetCustomers() []model.Customers {
	var customers []model.Customers
	return customers
}

func (customerRepository customerRepository) CreateCustomer(customer model.Customers) model.Customers {
	return customer
}

func (customerRepository customerRepository) UpdateCustomer(customer model.Customers) model.Customers {
	return customer
}

func (customerRepository customerRepository) DeleteCustomer(customer model.Customers) model.Customers {
	return customer
}

func NewCustomerRepository(conf config.BaseConfig) CustomerRepository {
	return &customerRepository{BaseConfig: conf}
}
