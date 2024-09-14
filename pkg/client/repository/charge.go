package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type ChargeRepository interface {
	BootstrapChargeForDB(request request.ChargeRequest)(response response.ChargeResponse) 
	GetChargeForPublic(request request.ChargeRequest)(response response.ChargeResponse) 
	GetChargeForInternal(request request.ChargeRequest)(response response.ChargeResponse) 
	GetChargeForPrivate(request request.ChargeRequest)(response response.ChargeResponse) 
	CreateChargeForPublic(request request.ChargeRequest)   (response response.ChargeResponse) 
	CreateChargeForInternal(request request.ChargeRequest) (response response.ChargeResponse)
	CreateChargeForPrivate(request request.ChargeRequest)  (response response.ChargeResponse)
	UpdateChargeForPublic(request request.ChargeRequest)   (response response.ChargeResponse)
	UpdateChargeForInternal(request request.ChargeRequest) (response response.ChargeResponse)
	UpdateChargeForPrivate(request request.ChargeRequest)  (response response.ChargeResponse)
	DeleteChargeForPublic(request request.ChargeRequest)   (response response.ChargeResponse)
	DeleteChargeForInternal(request request.ChargeRequest) (response response.ChargeResponse)
	DeleteChargeForPrivate(request request.ChargeRequest)  (response response.ChargeResponse)
}

type chargeRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (chargeRepository chargeRepository) BootstrapChargeForDB(request request.ChargeRequest) (response response.ChargeResponse) {
	fmt.Println("BootstrapChargeForDB")
	return response
}

//GET
func (chargeRepository chargeRepository) GetChargeForPublic(request request.ChargeRequest) (response response.ChargeResponse) {
	fmt.Println("GetChargeForPublic")
	return response
}

func (chargeRepository chargeRepository) GetChargeForInternal(request request.ChargeRequest) (response response.ChargeResponse ){
	fmt.Println("GetChargeForInternal")
	return response
}

func (chargeRepository chargeRepository) GetChargeForPrivate(request request.ChargeRequest) (response response.ChargeResponse) {
	fmt.Println("GetChargeForPrivate")
	return response
}

//CREATE
func (chargeRepository chargeRepository) CreateChargeForPublic(request request.ChargeRequest) (response response.ChargeResponse ){
	fmt.Println("CreateChargeForPublic")
	return response
}

func (chargeRepository chargeRepository) CreateChargeForInternal(request request.ChargeRequest) (response response.ChargeResponse) {
	fmt.Println("CreateChargeForInternal()")
	return response
}

func (chargeRepository chargeRepository) CreateChargeForPrivate(request request.ChargeRequest) (response response.ChargeResponse){
	fmt.Println("CreateChargeForPrivate()")
	return response
}

//UPDATE
func (chargeRepository chargeRepository) UpdateChargeForPublic(request request.ChargeRequest) (response response.ChargeResponse){
	fmt.Println("UpdateChargeForPublic()")
	return response
}

func (chargeRepository chargeRepository) UpdateChargeForInternal(request request.ChargeRequest) (response response.ChargeResponse){
	fmt.Println("UpdateChargeForInternal")
	return response
}

func (chargeRepository chargeRepository) UpdateChargeForPrivate(request request.ChargeRequest) (response response.ChargeResponse){
	fmt.Println("UpdateChargeForPrivate")
	return response
}

//DELETE
func (chargeRepository chargeRepository) DeleteChargeForPublic(request request.ChargeRequest) (response response.ChargeResponse){
	fmt.Println("DeleteChargeForPublic")
	return response
}

func (chargeRepository chargeRepository) DeleteChargeForInternal(request request.ChargeRequest) (response response.ChargeResponse ){
	fmt.Println("DeleteChargeForInternal")
	return response
}

func (chargeRepository chargeRepository) DeleteChargeForPrivate(request request.ChargeRequest) (response response.ChargeResponse){
	fmt.Println("DeleteChargeForPrivate")
	return response
}

func NewChargeRepository(conf config.BaseConfig) ChargeRepository {
	return &chargeRepository{BaseConfig: conf}
}