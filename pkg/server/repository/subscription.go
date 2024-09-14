package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type SubscriptionRepository interface {
	GetSubscriptions() []model.Subscriptions
	CreateSubscription(subscription model.Subscriptions) *gorm.DB
	UpdateSubscription(subscription model.Subscriptions) *gorm.DB
	DeleteSubscription(uuid string) *gorm.DB
}

type subscriptionRepository struct {
	BaseConfig config.BaseConfig
}


func (subscriptionRepository subscriptionRepository) GetSubscriptions() []model.Subscriptions {
	var subscriptions []model.Subscriptions
	subscriptionRepository.BaseConfig.DBConnection.Find(&subscriptions)
	return subscriptions
}

func (subscriptionRepository subscriptionRepository) CreateSubscription(subscription model.Subscriptions) *gorm.DB {
	results := subscriptionRepository.BaseConfig.DBConnection.Create(&subscription)
	return results
}

func (subscriptionRepository subscriptionRepository) UpdateSubscription(subscription model.Subscriptions) *gorm.DB {
	results := subscriptionRepository.BaseConfig.DBConnection.Model(model.Subscriptions{}).Where("uuid = ?", subscription.UUID).Updates(subscription)
	return results
}

func (subscriptionRepository subscriptionRepository) DeleteSubscription(uuid string) *gorm.DB {
	results := subscriptionRepository.BaseConfig.DBConnection.Model(model.Subscriptions{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewSubscriptionRepository(conf config.BaseConfig) SubscriptionRepository {
	return &subscriptionRepository{BaseConfig: conf}
}