package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type ProductControllerForPublic interface {
	GetProducts(c *gin.Context)
}

type productControllerForPublic struct {
	ProductRepository repository.ProductRepository
}

func (productController productControllerForPublic) GetProducts(c *gin.Context) {
	var productRequest request.ProductRequest
	if err := c.Bind(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.ProductResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Products: []response.Product{}})
		return
	}
	res := productController.ProductRepository.GetProducts()
	c.JSON(http.StatusOK, res)
	return
}


func NewProductControllerForPublic(productRepository repository.ProductRepository) ProductControllerForPublic {
	return &productControllerForPublic{ ProductRepository: productRepository}
}