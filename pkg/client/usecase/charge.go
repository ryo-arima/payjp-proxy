package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type ChargeUsecase interface {
	BootstrapChargeForDB(request request.ChargeRequest)
	GetChargeForPublic(request request.ChargeRequest)
	GetChargeForInternal(request request.ChargeRequest)
	GetChargeForPrivate(request request.ChargeRequest)
	CreateChargeForPublic(request request.ChargeRequest)
	CreateChargeForInternal(request request.ChargeRequest)
	CreateChargeForPrivate(request request.ChargeRequest)
	UpdateChargeForPublic(request request.ChargeRequest)
	UpdateChargeForInternal(request request.ChargeRequest)
	UpdateChargeForPrivate(request request.ChargeRequest)
	DeleteChargeForPublic(request request.ChargeRequest)
	DeleteChargeForInternal(request request.ChargeRequest)
	DeleteChargeForPrivate(request request.ChargeRequest)
}

type chargeUsecase struct {
	ChargeRepository   repository.ChargeRepository
}

//Bootstrap
func (chargeUsecase chargeUsecase) BootstrapChargeForDB(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.BootstrapChargeForDB(request)
	fmt.Println(charges)
}

//GET
func (chargeUsecase chargeUsecase) GetChargeForPublic(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.GetChargeForPublic(request)
	fmt.Println(charges)
}

func (chargeUsecase chargeUsecase) GetChargeForInternal(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.GetChargeForInternal(request)
	fmt.Println(charges)
}

func (chargeUsecase chargeUsecase) GetChargeForPrivate(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.GetChargeForPrivate(request)
	fmt.Println(charges)
}

//CREATE
func (chargeUsecase chargeUsecase) CreateChargeForPublic(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.CreateChargeForPublic(request)
	fmt.Println(charges)
}

func (chargeUsecase chargeUsecase) CreateChargeForInternal(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.CreateChargeForInternal(request)
	fmt.Println(charges)
}

func (chargeUsecase chargeUsecase) CreateChargeForPrivate(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.CreateChargeForPrivate(request)
	fmt.Println(charges)
}

//UPDATE
func (chargeUsecase chargeUsecase) UpdateChargeForPublic(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.UpdateChargeForPublic(request)
	fmt.Println(charges)
}

func (chargeUsecase chargeUsecase) UpdateChargeForInternal(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.UpdateChargeForInternal(request)
	fmt.Println(charges)
}

func (chargeUsecase chargeUsecase) UpdateChargeForPrivate(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.UpdateChargeForPrivate(request)
	fmt.Println(charges)
}

//DELETE
func (chargeUsecase chargeUsecase) DeleteChargeForPublic(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.DeleteChargeForPublic(request)
	fmt.Println(charges)
}

func (chargeUsecase chargeUsecase) DeleteChargeForInternal(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.DeleteChargeForInternal(request)
	fmt.Println(charges)
}

func (chargeUsecase chargeUsecase) DeleteChargeForPrivate(request request.ChargeRequest) {
	charges := chargeUsecase.ChargeRepository.DeleteChargeForPrivate(request)
	fmt.Println(charges)
}

func NewChargeUsecase(chargeRepository repository.ChargeRepository) ChargeUsecase {
	return &chargeUsecase{ ChargeRepository: chargeRepository}
}