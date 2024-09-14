package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type SubscriptionUsecase interface {
	BootstrapSubscriptionForDB(request request.SubscriptionRequest)
	GetSubscriptionForPublic(request request.SubscriptionRequest)
	GetSubscriptionForInternal(request request.SubscriptionRequest)
	GetSubscriptionForPrivate(request request.SubscriptionRequest)
	CreateSubscriptionForPublic(request request.SubscriptionRequest)
	CreateSubscriptionForInternal(request request.SubscriptionRequest)
	CreateSubscriptionForPrivate(request request.SubscriptionRequest)
	UpdateSubscriptionForPublic(request request.SubscriptionRequest)
	UpdateSubscriptionForInternal(request request.SubscriptionRequest)
	UpdateSubscriptionForPrivate(request request.SubscriptionRequest)
	DeleteSubscriptionForPublic(request request.SubscriptionRequest)
	DeleteSubscriptionForInternal(request request.SubscriptionRequest)
	DeleteSubscriptionForPrivate(request request.SubscriptionRequest)
}

type subscriptionUsecase struct {
	SubscriptionRepository   repository.SubscriptionRepository
}

//Bootstrap
func (subscriptionUsecase subscriptionUsecase) BootstrapSubscriptionForDB(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.BootstrapSubscriptionForDB(request)
	fmt.Println(subscriptions)
}

//GET
func (subscriptionUsecase subscriptionUsecase) GetSubscriptionForPublic(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.GetSubscriptionForPublic(request)
	fmt.Println(subscriptions)
}

func (subscriptionUsecase subscriptionUsecase) GetSubscriptionForInternal(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.GetSubscriptionForInternal(request)
	fmt.Println(subscriptions)
}

func (subscriptionUsecase subscriptionUsecase) GetSubscriptionForPrivate(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.GetSubscriptionForPrivate(request)
	fmt.Println(subscriptions)
}

//CREATE
func (subscriptionUsecase subscriptionUsecase) CreateSubscriptionForPublic(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.CreateSubscriptionForPublic(request)
	fmt.Println(subscriptions)
}

func (subscriptionUsecase subscriptionUsecase) CreateSubscriptionForInternal(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.CreateSubscriptionForInternal(request)
	fmt.Println(subscriptions)
}

func (subscriptionUsecase subscriptionUsecase) CreateSubscriptionForPrivate(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.CreateSubscriptionForPrivate(request)
	fmt.Println(subscriptions)
}

//UPDATE
func (subscriptionUsecase subscriptionUsecase) UpdateSubscriptionForPublic(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.UpdateSubscriptionForPublic(request)
	fmt.Println(subscriptions)
}

func (subscriptionUsecase subscriptionUsecase) UpdateSubscriptionForInternal(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.UpdateSubscriptionForInternal(request)
	fmt.Println(subscriptions)
}

func (subscriptionUsecase subscriptionUsecase) UpdateSubscriptionForPrivate(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.UpdateSubscriptionForPrivate(request)
	fmt.Println(subscriptions)
}

//DELETE
func (subscriptionUsecase subscriptionUsecase) DeleteSubscriptionForPublic(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.DeleteSubscriptionForPublic(request)
	fmt.Println(subscriptions)
}

func (subscriptionUsecase subscriptionUsecase) DeleteSubscriptionForInternal(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.DeleteSubscriptionForInternal(request)
	fmt.Println(subscriptions)
}

func (subscriptionUsecase subscriptionUsecase) DeleteSubscriptionForPrivate(request request.SubscriptionRequest) {
	subscriptions := subscriptionUsecase.SubscriptionRepository.DeleteSubscriptionForPrivate(request)
	fmt.Println(subscriptions)
}

func NewSubscriptionUsecase(subscriptionRepository repository.SubscriptionRepository) SubscriptionUsecase {
	return &subscriptionUsecase{ SubscriptionRepository: subscriptionRepository}
}