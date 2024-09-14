package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type SubscriptionControllerForPublic interface {
	GetSubscriptions(c *gin.Context)
}

type subscriptionControllerForPublic struct {
	SubscriptionRepository repository.SubscriptionRepository
}

func (subscriptionController subscriptionControllerForPublic) GetSubscriptions(c *gin.Context) {
	var subscriptionRequest request.SubscriptionRequest
	if err := c.Bind(&subscriptionRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.SubscriptionResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Subscriptions: []response.Subscription{}})
		return
	}
	res := subscriptionController.SubscriptionRepository.GetSubscriptions()
	c.JSON(http.StatusOK, res)
	return
}


func NewSubscriptionControllerForPublic(subscriptionRepository repository.SubscriptionRepository) SubscriptionControllerForPublic {
	return &subscriptionControllerForPublic{ SubscriptionRepository: subscriptionRepository}
}