package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type ProductControllerForPrivate interface {
	GetProducts(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type productControllerForPrivate struct {
	ProductRepository repository.ProductRepository
}

func (productController productControllerForPrivate) GetProducts(c *gin.Context) {
	var productRequest request.ProductRequest
	if err := c.Bind(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProductResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Products: []response.Product{}})
		return
	}
	res := productController.ProductRepository.GetProducts()
	c.JSON(http.StatusOK, res)
	return
}

func (productController productControllerForPrivate) CreateProduct(c *gin.Context) {
	var productRequest request.ProductRequest
	if err := c.Bind(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProductResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Products: []response.Product{}})
		return
	}
	var productModel model.Products
	res := productController.ProductRepository.CreateProduct(productModel)
	c.JSON(http.StatusOK, res)
	return
}

func (productController productControllerForPrivate) UpdateProduct(c *gin.Context) {
	var productRequest request.ProductRequest
	if err := c.Bind(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProductResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Products: []response.Product{}})
		return
	}
	var productModel model.Products
	res := productController.ProductRepository.UpdateProduct(productModel)
	c.JSON(http.StatusOK, res)
	return
}

func (productController productControllerForPrivate) DeleteProduct(c *gin.Context) {
	var productRequest request.ProductRequest
	if err := c.Bind(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProductResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Products: []response.Product{}})
		return
	}
	var productModel model.Products
	res := productController.ProductRepository.DeleteProduct(productModel)
	c.JSON(http.StatusOK, res)
	return
}

func NewProductControllerForPrivate(productRepository repository.ProductRepository) ProductControllerForPrivate {
	return &productControllerForPrivate{ProductRepository: productRepository}
}
