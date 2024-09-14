package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type ProductRepository interface {
	BootstrapProductForDB(request request.ProductRequest)(response response.ProductResponse) 
	GetProductForPublic(request request.ProductRequest)(response response.ProductResponse) 
	GetProductForInternal(request request.ProductRequest)(response response.ProductResponse) 
	GetProductForPrivate(request request.ProductRequest)(response response.ProductResponse) 
	CreateProductForPublic(request request.ProductRequest)   (response response.ProductResponse) 
	CreateProductForInternal(request request.ProductRequest) (response response.ProductResponse)
	CreateProductForPrivate(request request.ProductRequest)  (response response.ProductResponse)
	UpdateProductForPublic(request request.ProductRequest)   (response response.ProductResponse)
	UpdateProductForInternal(request request.ProductRequest) (response response.ProductResponse)
	UpdateProductForPrivate(request request.ProductRequest)  (response response.ProductResponse)
	DeleteProductForPublic(request request.ProductRequest)   (response response.ProductResponse)
	DeleteProductForInternal(request request.ProductRequest) (response response.ProductResponse)
	DeleteProductForPrivate(request request.ProductRequest)  (response response.ProductResponse)
}

type productRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (productRepository productRepository) BootstrapProductForDB(request request.ProductRequest) (response response.ProductResponse) {
	fmt.Println("BootstrapProductForDB")
	return response
}

//GET
func (productRepository productRepository) GetProductForPublic(request request.ProductRequest) (response response.ProductResponse) {
	fmt.Println("GetProductForPublic")
	return response
}

func (productRepository productRepository) GetProductForInternal(request request.ProductRequest) (response response.ProductResponse ){
	fmt.Println("GetProductForInternal")
	return response
}

func (productRepository productRepository) GetProductForPrivate(request request.ProductRequest) (response response.ProductResponse) {
	fmt.Println("GetProductForPrivate")
	return response
}

//CREATE
func (productRepository productRepository) CreateProductForPublic(request request.ProductRequest) (response response.ProductResponse ){
	fmt.Println("CreateProductForPublic")
	return response
}

func (productRepository productRepository) CreateProductForInternal(request request.ProductRequest) (response response.ProductResponse) {
	fmt.Println("CreateProductForInternal()")
	return response
}

func (productRepository productRepository) CreateProductForPrivate(request request.ProductRequest) (response response.ProductResponse){
	fmt.Println("CreateProductForPrivate()")
	return response
}

//UPDATE
func (productRepository productRepository) UpdateProductForPublic(request request.ProductRequest) (response response.ProductResponse){
	fmt.Println("UpdateProductForPublic()")
	return response
}

func (productRepository productRepository) UpdateProductForInternal(request request.ProductRequest) (response response.ProductResponse){
	fmt.Println("UpdateProductForInternal")
	return response
}

func (productRepository productRepository) UpdateProductForPrivate(request request.ProductRequest) (response response.ProductResponse){
	fmt.Println("UpdateProductForPrivate")
	return response
}

//DELETE
func (productRepository productRepository) DeleteProductForPublic(request request.ProductRequest) (response response.ProductResponse){
	fmt.Println("DeleteProductForPublic")
	return response
}

func (productRepository productRepository) DeleteProductForInternal(request request.ProductRequest) (response response.ProductResponse ){
	fmt.Println("DeleteProductForInternal")
	return response
}

func (productRepository productRepository) DeleteProductForPrivate(request request.ProductRequest) (response response.ProductResponse){
	fmt.Println("DeleteProductForPrivate")
	return response
}

func NewProductRepository(conf config.BaseConfig) ProductRepository {
	return &productRepository{BaseConfig: conf}
}