package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type AccountUsecase interface {
	BootstrapAccountForDB(request request.AccountRequest)
	GetAccountForPublic(request request.AccountRequest)
	GetAccountForInternal(request request.AccountRequest)
	GetAccountForPrivate(request request.AccountRequest)
	CreateAccountForPublic(request request.AccountRequest)
	CreateAccountForInternal(request request.AccountRequest)
	CreateAccountForPrivate(request request.AccountRequest)
	UpdateAccountForPublic(request request.AccountRequest)
	UpdateAccountForInternal(request request.AccountRequest)
	UpdateAccountForPrivate(request request.AccountRequest)
	DeleteAccountForPublic(request request.AccountRequest)
	DeleteAccountForInternal(request request.AccountRequest)
	DeleteAccountForPrivate(request request.AccountRequest)
}

type accountUsecase struct {
	AccountRepository   repository.AccountRepository
}

//Bootstrap
func (accountUsecase accountUsecase) BootstrapAccountForDB(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.BootstrapAccountForDB(request)
	fmt.Println(accounts)
}

//GET
func (accountUsecase accountUsecase) GetAccountForPublic(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.GetAccountForPublic(request)
	fmt.Println(accounts)
}

func (accountUsecase accountUsecase) GetAccountForInternal(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.GetAccountForInternal(request)
	fmt.Println(accounts)
}

func (accountUsecase accountUsecase) GetAccountForPrivate(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.GetAccountForPrivate(request)
	fmt.Println(accounts)
}

//CREATE
func (accountUsecase accountUsecase) CreateAccountForPublic(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.CreateAccountForPublic(request)
	fmt.Println(accounts)
}

func (accountUsecase accountUsecase) CreateAccountForInternal(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.CreateAccountForInternal(request)
	fmt.Println(accounts)
}

func (accountUsecase accountUsecase) CreateAccountForPrivate(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.CreateAccountForPrivate(request)
	fmt.Println(accounts)
}

//UPDATE
func (accountUsecase accountUsecase) UpdateAccountForPublic(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.UpdateAccountForPublic(request)
	fmt.Println(accounts)
}

func (accountUsecase accountUsecase) UpdateAccountForInternal(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.UpdateAccountForInternal(request)
	fmt.Println(accounts)
}

func (accountUsecase accountUsecase) UpdateAccountForPrivate(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.UpdateAccountForPrivate(request)
	fmt.Println(accounts)
}

//DELETE
func (accountUsecase accountUsecase) DeleteAccountForPublic(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.DeleteAccountForPublic(request)
	fmt.Println(accounts)
}

func (accountUsecase accountUsecase) DeleteAccountForInternal(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.DeleteAccountForInternal(request)
	fmt.Println(accounts)
}

func (accountUsecase accountUsecase) DeleteAccountForPrivate(request request.AccountRequest) {
	accounts := accountUsecase.AccountRepository.DeleteAccountForPrivate(request)
	fmt.Println(accounts)
}

func NewAccountUsecase(accountRepository repository.AccountRepository) AccountUsecase {
	return &accountUsecase{ AccountRepository: accountRepository}
}