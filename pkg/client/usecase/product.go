package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type ProductUsecase interface {
	BootstrapProductForDB(request request.ProductRequest)
	GetProductForPublic(request request.ProductRequest)
	GetProductForInternal(request request.ProductRequest)
	GetProductForPrivate(request request.ProductRequest)
	CreateProductForPublic(request request.ProductRequest)
	CreateProductForInternal(request request.ProductRequest)
	CreateProductForPrivate(request request.ProductRequest)
	UpdateProductForPublic(request request.ProductRequest)
	UpdateProductForInternal(request request.ProductRequest)
	UpdateProductForPrivate(request request.ProductRequest)
	DeleteProductForPublic(request request.ProductRequest)
	DeleteProductForInternal(request request.ProductRequest)
	DeleteProductForPrivate(request request.ProductRequest)
}

type productUsecase struct {
	ProductRepository   repository.ProductRepository
}

//Bootstrap
func (productUsecase productUsecase) BootstrapProductForDB(request request.ProductRequest) {
	products := productUsecase.ProductRepository.BootstrapProductForDB(request)
	fmt.Println(products)
}

//GET
func (productUsecase productUsecase) GetProductForPublic(request request.ProductRequest) {
	products := productUsecase.ProductRepository.GetProductForPublic(request)
	fmt.Println(products)
}

func (productUsecase productUsecase) GetProductForInternal(request request.ProductRequest) {
	products := productUsecase.ProductRepository.GetProductForInternal(request)
	fmt.Println(products)
}

func (productUsecase productUsecase) GetProductForPrivate(request request.ProductRequest) {
	products := productUsecase.ProductRepository.GetProductForPrivate(request)
	fmt.Println(products)
}

//CREATE
func (productUsecase productUsecase) CreateProductForPublic(request request.ProductRequest) {
	products := productUsecase.ProductRepository.CreateProductForPublic(request)
	fmt.Println(products)
}

func (productUsecase productUsecase) CreateProductForInternal(request request.ProductRequest) {
	products := productUsecase.ProductRepository.CreateProductForInternal(request)
	fmt.Println(products)
}

func (productUsecase productUsecase) CreateProductForPrivate(request request.ProductRequest) {
	products := productUsecase.ProductRepository.CreateProductForPrivate(request)
	fmt.Println(products)
}

//UPDATE
func (productUsecase productUsecase) UpdateProductForPublic(request request.ProductRequest) {
	products := productUsecase.ProductRepository.UpdateProductForPublic(request)
	fmt.Println(products)
}

func (productUsecase productUsecase) UpdateProductForInternal(request request.ProductRequest) {
	products := productUsecase.ProductRepository.UpdateProductForInternal(request)
	fmt.Println(products)
}

func (productUsecase productUsecase) UpdateProductForPrivate(request request.ProductRequest) {
	products := productUsecase.ProductRepository.UpdateProductForPrivate(request)
	fmt.Println(products)
}

//DELETE
func (productUsecase productUsecase) DeleteProductForPublic(request request.ProductRequest) {
	products := productUsecase.ProductRepository.DeleteProductForPublic(request)
	fmt.Println(products)
}

func (productUsecase productUsecase) DeleteProductForInternal(request request.ProductRequest) {
	products := productUsecase.ProductRepository.DeleteProductForInternal(request)
	fmt.Println(products)
}

func (productUsecase productUsecase) DeleteProductForPrivate(request request.ProductRequest) {
	products := productUsecase.ProductRepository.DeleteProductForPrivate(request)
	fmt.Println(products)
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{ ProductRepository: productRepository}
}