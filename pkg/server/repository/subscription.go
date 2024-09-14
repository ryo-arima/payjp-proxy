package repository

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type SubscriptionRepository interface {
	GetSubscriptions() []model.Subscriptions
	CreateSubscription(subscription model.Subscriptions) model.Subscriptions
	UpdateSubscription(subscription model.Subscriptions) model.Subscriptions
	DeleteSubscription(subscription model.Subscriptions) model.Subscriptions
}

type subscriptionRepository struct {
	BaseConfig config.BaseConfig
}

func (subscriptionRepository subscriptionRepository) GetSubscriptions() []model.Subscriptions {
	var subscriptions []model.Subscriptions
	return subscriptions
}

func (subscriptionRepository subscriptionRepository) CreateSubscription(subscription model.Subscriptions) model.Subscriptions {
	return subscription
}

func (subscriptionRepository subscriptionRepository) UpdateSubscription(subscription model.Subscriptions) model.Subscriptions {
	return subscription
}

func (subscriptionRepository subscriptionRepository) DeleteSubscription(subscription model.Subscriptions) model.Subscriptions {
	return subscription
}

func NewSubscriptionRepository(conf config.BaseConfig) SubscriptionRepository {
	return &subscriptionRepository{BaseConfig: conf}
}
