package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type SubscriptionRepository interface {
	BootstrapSubscriptionForDB(request request.SubscriptionRequest)(response response.SubscriptionResponse) 
	GetSubscriptionForPublic(request request.SubscriptionRequest)(response response.SubscriptionResponse) 
	GetSubscriptionForInternal(request request.SubscriptionRequest)(response response.SubscriptionResponse) 
	GetSubscriptionForPrivate(request request.SubscriptionRequest)(response response.SubscriptionResponse) 
	CreateSubscriptionForPublic(request request.SubscriptionRequest)   (response response.SubscriptionResponse) 
	CreateSubscriptionForInternal(request request.SubscriptionRequest) (response response.SubscriptionResponse)
	CreateSubscriptionForPrivate(request request.SubscriptionRequest)  (response response.SubscriptionResponse)
	UpdateSubscriptionForPublic(request request.SubscriptionRequest)   (response response.SubscriptionResponse)
	UpdateSubscriptionForInternal(request request.SubscriptionRequest) (response response.SubscriptionResponse)
	UpdateSubscriptionForPrivate(request request.SubscriptionRequest)  (response response.SubscriptionResponse)
	DeleteSubscriptionForPublic(request request.SubscriptionRequest)   (response response.SubscriptionResponse)
	DeleteSubscriptionForInternal(request request.SubscriptionRequest) (response response.SubscriptionResponse)
	DeleteSubscriptionForPrivate(request request.SubscriptionRequest)  (response response.SubscriptionResponse)
}

type subscriptionRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (subscriptionRepository subscriptionRepository) BootstrapSubscriptionForDB(request request.SubscriptionRequest) (response response.SubscriptionResponse) {
	fmt.Println("BootstrapSubscriptionForDB")
	return response
}

//GET
func (subscriptionRepository subscriptionRepository) GetSubscriptionForPublic(request request.SubscriptionRequest) (response response.SubscriptionResponse) {
	fmt.Println("GetSubscriptionForPublic")
	return response
}

func (subscriptionRepository subscriptionRepository) GetSubscriptionForInternal(request request.SubscriptionRequest) (response response.SubscriptionResponse ){
	fmt.Println("GetSubscriptionForInternal")
	return response
}

func (subscriptionRepository subscriptionRepository) GetSubscriptionForPrivate(request request.SubscriptionRequest) (response response.SubscriptionResponse) {
	fmt.Println("GetSubscriptionForPrivate")
	return response
}

//CREATE
func (subscriptionRepository subscriptionRepository) CreateSubscriptionForPublic(request request.SubscriptionRequest) (response response.SubscriptionResponse ){
	fmt.Println("CreateSubscriptionForPublic")
	return response
}

func (subscriptionRepository subscriptionRepository) CreateSubscriptionForInternal(request request.SubscriptionRequest) (response response.SubscriptionResponse) {
	fmt.Println("CreateSubscriptionForInternal()")
	return response
}

func (subscriptionRepository subscriptionRepository) CreateSubscriptionForPrivate(request request.SubscriptionRequest) (response response.SubscriptionResponse){
	fmt.Println("CreateSubscriptionForPrivate()")
	return response
}

//UPDATE
func (subscriptionRepository subscriptionRepository) UpdateSubscriptionForPublic(request request.SubscriptionRequest) (response response.SubscriptionResponse){
	fmt.Println("UpdateSubscriptionForPublic()")
	return response
}

func (subscriptionRepository subscriptionRepository) UpdateSubscriptionForInternal(request request.SubscriptionRequest) (response response.SubscriptionResponse){
	fmt.Println("UpdateSubscriptionForInternal")
	return response
}

func (subscriptionRepository subscriptionRepository) UpdateSubscriptionForPrivate(request request.SubscriptionRequest) (response response.SubscriptionResponse){
	fmt.Println("UpdateSubscriptionForPrivate")
	return response
}

//DELETE
func (subscriptionRepository subscriptionRepository) DeleteSubscriptionForPublic(request request.SubscriptionRequest) (response response.SubscriptionResponse){
	fmt.Println("DeleteSubscriptionForPublic")
	return response
}

func (subscriptionRepository subscriptionRepository) DeleteSubscriptionForInternal(request request.SubscriptionRequest) (response response.SubscriptionResponse ){
	fmt.Println("DeleteSubscriptionForInternal")
	return response
}

func (subscriptionRepository subscriptionRepository) DeleteSubscriptionForPrivate(request request.SubscriptionRequest) (response response.SubscriptionResponse){
	fmt.Println("DeleteSubscriptionForPrivate")
	return response
}

func NewSubscriptionRepository(conf config.BaseConfig) SubscriptionRepository {
	return &subscriptionRepository{BaseConfig: conf}
}