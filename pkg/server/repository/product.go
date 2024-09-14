package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() []model.Products
	CreateProduct(product model.Products) *gorm.DB
	UpdateProduct(product model.Products) *gorm.DB
	DeleteProduct(uuid string) *gorm.DB
}

type productRepository struct {
	BaseConfig config.BaseConfig
}


func (productRepository productRepository) GetProducts() []model.Products {
	var products []model.Products
	productRepository.BaseConfig.DBConnection.Find(&products)
	return products
}

func (productRepository productRepository) CreateProduct(product model.Products) *gorm.DB {
	results := productRepository.BaseConfig.DBConnection.Create(&product)
	return results
}

func (productRepository productRepository) UpdateProduct(product model.Products) *gorm.DB {
	results := productRepository.BaseConfig.DBConnection.Model(model.Products{}).Where("uuid = ?", product.UUID).Updates(product)
	return results
}

func (productRepository productRepository) DeleteProduct(uuid string) *gorm.DB {
	results := productRepository.BaseConfig.DBConnection.Model(model.Products{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewProductRepository(conf config.BaseConfig) ProductRepository {
	return &productRepository{BaseConfig: conf}
}