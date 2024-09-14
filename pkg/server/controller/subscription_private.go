package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type SubscriptionControllerForPrivate interface {
	GetSubscriptions(c *gin.Context)
	CreateSubscription(c *gin.Context)
	UpdateSubscription(c *gin.Context)
	DeleteSubscription(c *gin.Context)
}

type subscriptionControllerForPrivate struct {
	SubscriptionRepository repository.SubscriptionRepository
}

func (subscriptionController subscriptionControllerForPrivate) GetSubscriptions(c *gin.Context) {
	var subscriptionRequest request.SubscriptionRequest
	if err := c.Bind(&subscriptionRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.SubscriptionResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Subscriptions: []response.Subscription{}})
		return
	}
	res := subscriptionController.SubscriptionRepository.GetSubscriptions()
	c.JSON(http.StatusOK, res)
	return
}


func (subscriptionController subscriptionControllerForPrivate) CreateSubscription(c *gin.Context) {
	var subscriptionRequest request.SubscriptionRequest
	if err := c.Bind(&subscriptionRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.SubscriptionResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Subscriptions: []response.Subscription{}})
		return
	}
	var subscriptionModel model.Subscriptions
	res := subscriptionController.SubscriptionRepository.CreateSubscription(subscriptionModel)
	c.JSON(http.StatusOK, res)
	return
}


func (subscriptionController subscriptionControllerForPrivate) UpdateSubscription(c *gin.Context) {
	var subscriptionRequest request.SubscriptionRequest
	if err := c.Bind(&subscriptionRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.SubscriptionResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Subscriptions: []response.Subscription{}})
		return
	}
	var subscriptionModel model.Subscriptions
	res := subscriptionController.SubscriptionRepository.UpdateSubscription(subscriptionModel)
	c.JSON(http.StatusOK, res)
	return
}


func (subscriptionController subscriptionControllerForPrivate) DeleteSubscription(c *gin.Context) {
	var subscriptionRequest request.SubscriptionRequest
	if err := c.Bind(&subscriptionRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.SubscriptionResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Subscriptions: []response.Subscription{}})
		return
	}
	var uuid string
	res := subscriptionController.SubscriptionRepository.DeleteSubscription(uuid)
	c.JSON(http.StatusOK, res)
	return
}


func NewSubscriptionControllerForPrivate(subscriptionRepository repository.SubscriptionRepository) SubscriptionControllerForPrivate {
	return &subscriptionControllerForPrivate{ SubscriptionRepository: subscriptionRepository}
}