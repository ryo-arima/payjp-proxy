package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type TransferUsecase interface {
	BootstrapTransferForDB(request request.TransferRequest)
	GetTransferForPublic(request request.TransferRequest)
	GetTransferForInternal(request request.TransferRequest)
	GetTransferForPrivate(request request.TransferRequest)
	CreateTransferForPublic(request request.TransferRequest)
	CreateTransferForInternal(request request.TransferRequest)
	CreateTransferForPrivate(request request.TransferRequest)
	UpdateTransferForPublic(request request.TransferRequest)
	UpdateTransferForInternal(request request.TransferRequest)
	UpdateTransferForPrivate(request request.TransferRequest)
	DeleteTransferForPublic(request request.TransferRequest)
	DeleteTransferForInternal(request request.TransferRequest)
	DeleteTransferForPrivate(request request.TransferRequest)
}

type transferUsecase struct {
	TransferRepository   repository.TransferRepository
}

//Bootstrap
func (transferUsecase transferUsecase) BootstrapTransferForDB(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.BootstrapTransferForDB(request)
	fmt.Println(transfers)
}

//GET
func (transferUsecase transferUsecase) GetTransferForPublic(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.GetTransferForPublic(request)
	fmt.Println(transfers)
}

func (transferUsecase transferUsecase) GetTransferForInternal(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.GetTransferForInternal(request)
	fmt.Println(transfers)
}

func (transferUsecase transferUsecase) GetTransferForPrivate(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.GetTransferForPrivate(request)
	fmt.Println(transfers)
}

//CREATE
func (transferUsecase transferUsecase) CreateTransferForPublic(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.CreateTransferForPublic(request)
	fmt.Println(transfers)
}

func (transferUsecase transferUsecase) CreateTransferForInternal(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.CreateTransferForInternal(request)
	fmt.Println(transfers)
}

func (transferUsecase transferUsecase) CreateTransferForPrivate(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.CreateTransferForPrivate(request)
	fmt.Println(transfers)
}

//UPDATE
func (transferUsecase transferUsecase) UpdateTransferForPublic(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.UpdateTransferForPublic(request)
	fmt.Println(transfers)
}

func (transferUsecase transferUsecase) UpdateTransferForInternal(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.UpdateTransferForInternal(request)
	fmt.Println(transfers)
}

func (transferUsecase transferUsecase) UpdateTransferForPrivate(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.UpdateTransferForPrivate(request)
	fmt.Println(transfers)
}

//DELETE
func (transferUsecase transferUsecase) DeleteTransferForPublic(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.DeleteTransferForPublic(request)
	fmt.Println(transfers)
}

func (transferUsecase transferUsecase) DeleteTransferForInternal(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.DeleteTransferForInternal(request)
	fmt.Println(transfers)
}

func (transferUsecase transferUsecase) DeleteTransferForPrivate(request request.TransferRequest) {
	transfers := transferUsecase.TransferRepository.DeleteTransferForPrivate(request)
	fmt.Println(transfers)
}

func NewTransferUsecase(transferRepository repository.TransferRepository) TransferUsecase {
	return &transferUsecase{ TransferRepository: transferRepository}
}