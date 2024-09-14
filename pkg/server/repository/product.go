package repository

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type ProductRepository interface {
	GetProducts() []model.Products
	CreateProduct(product model.Products) model.Products
	UpdateProduct(product model.Products) model.Products
	DeleteProduct(product model.Products) model.Products
}

type productRepository struct {
	BaseConfig config.BaseConfig
}

func (productRepository productRepository) GetProducts() []model.Products {
	var products []model.Products
	return products
}

func (productRepository productRepository) CreateProduct(product model.Products) model.Products {
	return product
}

func (productRepository productRepository) UpdateProduct(product model.Products) model.Products {
	return product
}

func (productRepository productRepository) DeleteProduct(product model.Products) model.Products {
	return product
}

func NewProductRepository(conf config.BaseConfig) ProductRepository {
	return &productRepository{BaseConfig: conf}
}
