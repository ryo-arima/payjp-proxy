package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type BalanceRepository interface {
	BootstrapBalanceForDB(request request.BalanceRequest)(response response.BalanceResponse) 
	GetBalanceForPublic(request request.BalanceRequest)(response response.BalanceResponse) 
	GetBalanceForInternal(request request.BalanceRequest)(response response.BalanceResponse) 
	GetBalanceForPrivate(request request.BalanceRequest)(response response.BalanceResponse) 
	CreateBalanceForPublic(request request.BalanceRequest)   (response response.BalanceResponse) 
	CreateBalanceForInternal(request request.BalanceRequest) (response response.BalanceResponse)
	CreateBalanceForPrivate(request request.BalanceRequest)  (response response.BalanceResponse)
	UpdateBalanceForPublic(request request.BalanceRequest)   (response response.BalanceResponse)
	UpdateBalanceForInternal(request request.BalanceRequest) (response response.BalanceResponse)
	UpdateBalanceForPrivate(request request.BalanceRequest)  (response response.BalanceResponse)
	DeleteBalanceForPublic(request request.BalanceRequest)   (response response.BalanceResponse)
	DeleteBalanceForInternal(request request.BalanceRequest) (response response.BalanceResponse)
	DeleteBalanceForPrivate(request request.BalanceRequest)  (response response.BalanceResponse)
}

type balanceRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (balanceRepository balanceRepository) BootstrapBalanceForDB(request request.BalanceRequest) (response response.BalanceResponse) {
	fmt.Println("BootstrapBalanceForDB")
	return response
}

//GET
func (balanceRepository balanceRepository) GetBalanceForPublic(request request.BalanceRequest) (response response.BalanceResponse) {
	fmt.Println("GetBalanceForPublic")
	return response
}

func (balanceRepository balanceRepository) GetBalanceForInternal(request request.BalanceRequest) (response response.BalanceResponse ){
	fmt.Println("GetBalanceForInternal")
	return response
}

func (balanceRepository balanceRepository) GetBalanceForPrivate(request request.BalanceRequest) (response response.BalanceResponse) {
	fmt.Println("GetBalanceForPrivate")
	return response
}

//CREATE
func (balanceRepository balanceRepository) CreateBalanceForPublic(request request.BalanceRequest) (response response.BalanceResponse ){
	fmt.Println("CreateBalanceForPublic")
	return response
}

func (balanceRepository balanceRepository) CreateBalanceForInternal(request request.BalanceRequest) (response response.BalanceResponse) {
	fmt.Println("CreateBalanceForInternal()")
	return response
}

func (balanceRepository balanceRepository) CreateBalanceForPrivate(request request.BalanceRequest) (response response.BalanceResponse){
	fmt.Println("CreateBalanceForPrivate()")
	return response
}

//UPDATE
func (balanceRepository balanceRepository) UpdateBalanceForPublic(request request.BalanceRequest) (response response.BalanceResponse){
	fmt.Println("UpdateBalanceForPublic()")
	return response
}

func (balanceRepository balanceRepository) UpdateBalanceForInternal(request request.BalanceRequest) (response response.BalanceResponse){
	fmt.Println("UpdateBalanceForInternal")
	return response
}

func (balanceRepository balanceRepository) UpdateBalanceForPrivate(request request.BalanceRequest) (response response.BalanceResponse){
	fmt.Println("UpdateBalanceForPrivate")
	return response
}

//DELETE
func (balanceRepository balanceRepository) DeleteBalanceForPublic(request request.BalanceRequest) (response response.BalanceResponse){
	fmt.Println("DeleteBalanceForPublic")
	return response
}

func (balanceRepository balanceRepository) DeleteBalanceForInternal(request request.BalanceRequest) (response response.BalanceResponse ){
	fmt.Println("DeleteBalanceForInternal")
	return response
}

func (balanceRepository balanceRepository) DeleteBalanceForPrivate(request request.BalanceRequest) (response response.BalanceResponse){
	fmt.Println("DeleteBalanceForPrivate")
	return response
}

func NewBalanceRepository(conf config.BaseConfig) BalanceRepository {
	return &balanceRepository{BaseConfig: conf}
}