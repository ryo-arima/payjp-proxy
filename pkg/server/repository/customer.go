package repository

import (
	"fmt"
	"log"

	"github.com/payjp/payjp-go/v1"
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

func (customerRepository customerRepository) CreateCustomer(customerModel model.Customers) model.Customers {
	//customerParams := payjp.Customer{
	//	Email:       customerModel.Email,
	//	Description: customerModel.Description,
	//}
	customerParams := payjp.Customer{
		Email:       "test1@mail.com",
		Description: "test1",
	}

	res, err := customerRepository.BaseConfig.PayJpConfig.Service.Customer.Create(customerParams)
	if err != nil {
		log.Fatalf("顧客作成に失敗しました: %v", err)
	}

	fmt.Println(res)

	return customerModel
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
