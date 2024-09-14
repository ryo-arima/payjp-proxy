package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type AccountRepository interface {
	BootstrapAccountForDB(request request.AccountRequest)(response response.AccountResponse) 
	GetAccountForPublic(request request.AccountRequest)(response response.AccountResponse) 
	GetAccountForInternal(request request.AccountRequest)(response response.AccountResponse) 
	GetAccountForPrivate(request request.AccountRequest)(response response.AccountResponse) 
	CreateAccountForPublic(request request.AccountRequest)   (response response.AccountResponse) 
	CreateAccountForInternal(request request.AccountRequest) (response response.AccountResponse)
	CreateAccountForPrivate(request request.AccountRequest)  (response response.AccountResponse)
	UpdateAccountForPublic(request request.AccountRequest)   (response response.AccountResponse)
	UpdateAccountForInternal(request request.AccountRequest) (response response.AccountResponse)
	UpdateAccountForPrivate(request request.AccountRequest)  (response response.AccountResponse)
	DeleteAccountForPublic(request request.AccountRequest)   (response response.AccountResponse)
	DeleteAccountForInternal(request request.AccountRequest) (response response.AccountResponse)
	DeleteAccountForPrivate(request request.AccountRequest)  (response response.AccountResponse)
}

type accountRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (accountRepository accountRepository) BootstrapAccountForDB(request request.AccountRequest) (response response.AccountResponse) {
	fmt.Println("BootstrapAccountForDB")
	return response
}

//GET
func (accountRepository accountRepository) GetAccountForPublic(request request.AccountRequest) (response response.AccountResponse) {
	fmt.Println("GetAccountForPublic")
	return response
}

func (accountRepository accountRepository) GetAccountForInternal(request request.AccountRequest) (response response.AccountResponse ){
	fmt.Println("GetAccountForInternal")
	return response
}

func (accountRepository accountRepository) GetAccountForPrivate(request request.AccountRequest) (response response.AccountResponse) {
	fmt.Println("GetAccountForPrivate")
	return response
}

//CREATE
func (accountRepository accountRepository) CreateAccountForPublic(request request.AccountRequest) (response response.AccountResponse ){
	fmt.Println("CreateAccountForPublic")
	return response
}

func (accountRepository accountRepository) CreateAccountForInternal(request request.AccountRequest) (response response.AccountResponse) {
	fmt.Println("CreateAccountForInternal()")
	return response
}

func (accountRepository accountRepository) CreateAccountForPrivate(request request.AccountRequest) (response response.AccountResponse){
	fmt.Println("CreateAccountForPrivate()")
	return response
}

//UPDATE
func (accountRepository accountRepository) UpdateAccountForPublic(request request.AccountRequest) (response response.AccountResponse){
	fmt.Println("UpdateAccountForPublic()")
	return response
}

func (accountRepository accountRepository) UpdateAccountForInternal(request request.AccountRequest) (response response.AccountResponse){
	fmt.Println("UpdateAccountForInternal")
	return response
}

func (accountRepository accountRepository) UpdateAccountForPrivate(request request.AccountRequest) (response response.AccountResponse){
	fmt.Println("UpdateAccountForPrivate")
	return response
}

//DELETE
func (accountRepository accountRepository) DeleteAccountForPublic(request request.AccountRequest) (response response.AccountResponse){
	fmt.Println("DeleteAccountForPublic")
	return response
}

func (accountRepository accountRepository) DeleteAccountForInternal(request request.AccountRequest) (response response.AccountResponse ){
	fmt.Println("DeleteAccountForInternal")
	return response
}

func (accountRepository accountRepository) DeleteAccountForPrivate(request request.AccountRequest) (response response.AccountResponse){
	fmt.Println("DeleteAccountForPrivate")
	return response
}

func NewAccountRepository(conf config.BaseConfig) AccountRepository {
	return &accountRepository{BaseConfig: conf}
}