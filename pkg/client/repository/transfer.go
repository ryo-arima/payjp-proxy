package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type TransferRepository interface {
	BootstrapTransferForDB(request request.TransferRequest)(response response.TransferResponse) 
	GetTransferForPublic(request request.TransferRequest)(response response.TransferResponse) 
	GetTransferForInternal(request request.TransferRequest)(response response.TransferResponse) 
	GetTransferForPrivate(request request.TransferRequest)(response response.TransferResponse) 
	CreateTransferForPublic(request request.TransferRequest)   (response response.TransferResponse) 
	CreateTransferForInternal(request request.TransferRequest) (response response.TransferResponse)
	CreateTransferForPrivate(request request.TransferRequest)  (response response.TransferResponse)
	UpdateTransferForPublic(request request.TransferRequest)   (response response.TransferResponse)
	UpdateTransferForInternal(request request.TransferRequest) (response response.TransferResponse)
	UpdateTransferForPrivate(request request.TransferRequest)  (response response.TransferResponse)
	DeleteTransferForPublic(request request.TransferRequest)   (response response.TransferResponse)
	DeleteTransferForInternal(request request.TransferRequest) (response response.TransferResponse)
	DeleteTransferForPrivate(request request.TransferRequest)  (response response.TransferResponse)
}

type transferRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (transferRepository transferRepository) BootstrapTransferForDB(request request.TransferRequest) (response response.TransferResponse) {
	fmt.Println("BootstrapTransferForDB")
	return response
}

//GET
func (transferRepository transferRepository) GetTransferForPublic(request request.TransferRequest) (response response.TransferResponse) {
	fmt.Println("GetTransferForPublic")
	return response
}

func (transferRepository transferRepository) GetTransferForInternal(request request.TransferRequest) (response response.TransferResponse ){
	fmt.Println("GetTransferForInternal")
	return response
}

func (transferRepository transferRepository) GetTransferForPrivate(request request.TransferRequest) (response response.TransferResponse) {
	fmt.Println("GetTransferForPrivate")
	return response
}

//CREATE
func (transferRepository transferRepository) CreateTransferForPublic(request request.TransferRequest) (response response.TransferResponse ){
	fmt.Println("CreateTransferForPublic")
	return response
}

func (transferRepository transferRepository) CreateTransferForInternal(request request.TransferRequest) (response response.TransferResponse) {
	fmt.Println("CreateTransferForInternal()")
	return response
}

func (transferRepository transferRepository) CreateTransferForPrivate(request request.TransferRequest) (response response.TransferResponse){
	fmt.Println("CreateTransferForPrivate()")
	return response
}

//UPDATE
func (transferRepository transferRepository) UpdateTransferForPublic(request request.TransferRequest) (response response.TransferResponse){
	fmt.Println("UpdateTransferForPublic()")
	return response
}

func (transferRepository transferRepository) UpdateTransferForInternal(request request.TransferRequest) (response response.TransferResponse){
	fmt.Println("UpdateTransferForInternal")
	return response
}

func (transferRepository transferRepository) UpdateTransferForPrivate(request request.TransferRequest) (response response.TransferResponse){
	fmt.Println("UpdateTransferForPrivate")
	return response
}

//DELETE
func (transferRepository transferRepository) DeleteTransferForPublic(request request.TransferRequest) (response response.TransferResponse){
	fmt.Println("DeleteTransferForPublic")
	return response
}

func (transferRepository transferRepository) DeleteTransferForInternal(request request.TransferRequest) (response response.TransferResponse ){
	fmt.Println("DeleteTransferForInternal")
	return response
}

func (transferRepository transferRepository) DeleteTransferForPrivate(request request.TransferRequest) (response response.TransferResponse){
	fmt.Println("DeleteTransferForPrivate")
	return response
}

func NewTransferRepository(conf config.BaseConfig) TransferRepository {
	return &transferRepository{BaseConfig: conf}
}