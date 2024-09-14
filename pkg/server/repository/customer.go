package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomers() []model.Customers
	CreateCustomer(customer model.Customers) *gorm.DB
	UpdateCustomer(customer model.Customers) *gorm.DB
	DeleteCustomer(uuid string) *gorm.DB
}

type customerRepository struct {
	BaseConfig config.BaseConfig
}


func (customerRepository customerRepository) GetCustomers() []model.Customers {
	var customers []model.Customers
	customerRepository.BaseConfig.DBConnection.Find(&customers)
	return customers
}

func (customerRepository customerRepository) CreateCustomer(customer model.Customers) *gorm.DB {
	results := customerRepository.BaseConfig.DBConnection.Create(&customer)
	return results
}

func (customerRepository customerRepository) UpdateCustomer(customer model.Customers) *gorm.DB {
	results := customerRepository.BaseConfig.DBConnection.Model(model.Customers{}).Where("uuid = ?", customer.UUID).Updates(customer)
	return results
}

func (customerRepository customerRepository) DeleteCustomer(uuid string) *gorm.DB {
	results := customerRepository.BaseConfig.DBConnection.Model(model.Customers{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewCustomerRepository(conf config.BaseConfig) CustomerRepository {
	return &customerRepository{BaseConfig: conf}
}