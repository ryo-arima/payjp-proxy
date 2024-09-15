package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type BalanceUsecase interface {
	GetBalanceForInternal(request request.BalanceRequest)
	GetBalanceForPrivate(request request.BalanceRequest)
	CreateBalanceForInternal(request request.BalanceRequest)
	CreateBalanceForPrivate(request request.BalanceRequest)
	UpdateBalanceForInternal(request request.BalanceRequest)
	UpdateBalanceForPrivate(request request.BalanceRequest)
	DeleteBalanceForInternal(request request.BalanceRequest)
	DeleteBalanceForPrivate(request request.BalanceRequest)
}

type balanceUsecase struct {
	BalanceRepository repository.BalanceRepository
}

// GET
func (balanceUsecase balanceUsecase) GetBalanceForInternal(request request.BalanceRequest) {
	balances := balanceUsecase.BalanceRepository.GetBalanceForInternal(request)
	fmt.Println(balances)
}

func (balanceUsecase balanceUsecase) GetBalanceForPrivate(request request.BalanceRequest) {
	balances := balanceUsecase.BalanceRepository.GetBalanceForPrivate(request)
	fmt.Println(balances)
}

// CREATE
func (balanceUsecase balanceUsecase) CreateBalanceForInternal(request request.BalanceRequest) {
	balances := balanceUsecase.BalanceRepository.CreateBalanceForInternal(request)
	fmt.Println(balances)
}

func (balanceUsecase balanceUsecase) CreateBalanceForPrivate(request request.BalanceRequest) {
	balances := balanceUsecase.BalanceRepository.CreateBalanceForPrivate(request)
	fmt.Println(balances)
}

// UPDATE
func (balanceUsecase balanceUsecase) UpdateBalanceForInternal(request request.BalanceRequest) {
	balances := balanceUsecase.BalanceRepository.UpdateBalanceForInternal(request)
	fmt.Println(balances)
}

func (balanceUsecase balanceUsecase) UpdateBalanceForPrivate(request request.BalanceRequest) {
	balances := balanceUsecase.BalanceRepository.UpdateBalanceForPrivate(request)
	fmt.Println(balances)
}

// DELETE
func (balanceUsecase balanceUsecase) DeleteBalanceForInternal(request request.BalanceRequest) {
	balances := balanceUsecase.BalanceRepository.DeleteBalanceForInternal(request)
	fmt.Println(balances)
}

func (balanceUsecase balanceUsecase) DeleteBalanceForPrivate(request request.BalanceRequest) {
	balances := balanceUsecase.BalanceRepository.DeleteBalanceForPrivate(request)
	fmt.Println(balances)
}

func NewBalanceUsecase(balanceRepository repository.BalanceRepository) BalanceUsecase {
	return &balanceUsecase{BalanceRepository: balanceRepository}
}
