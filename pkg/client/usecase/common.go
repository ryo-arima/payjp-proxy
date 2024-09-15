package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type CommonUsecase interface {
	BootstrapCommonForDB(request request.CommonRequest)
	GetCommonForPublic(request request.CommonRequest)
	GetCommonForInternal(request request.CommonRequest)
	GetCommonForPrivate(request request.CommonRequest)
	CreateCommonForPublic(request request.CommonRequest)
	CreateCommonForInternal(request request.CommonRequest)
	CreateCommonForPrivate(request request.CommonRequest)
	UpdateCommonForPublic(request request.CommonRequest)
	UpdateCommonForInternal(request request.CommonRequest)
	UpdateCommonForPrivate(request request.CommonRequest)
	DeleteCommonForPublic(request request.CommonRequest)
	DeleteCommonForInternal(request request.CommonRequest)
	DeleteCommonForPrivate(request request.CommonRequest)
}

type commonUsecase struct {
	CommonRepository repository.CommonRepository
}

// Bootstrap
func (commonUsecase commonUsecase) BootstrapCommonForDB(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.BootstrapCommonForDB(request)
	fmt.Println(commons)
}

// GET
func (commonUsecase commonUsecase) GetCommonForPublic(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.GetCommonForPublic(request)
	fmt.Println(commons)
}

func (commonUsecase commonUsecase) GetCommonForInternal(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.GetCommonForInternal(request)
	fmt.Println(commons)
}

func (commonUsecase commonUsecase) GetCommonForPrivate(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.GetCommonForPrivate(request)
	fmt.Println(commons)
}

// CREATE
func (commonUsecase commonUsecase) CreateCommonForPublic(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.CreateCommonForPublic(request)
	fmt.Println(commons)
}

func (commonUsecase commonUsecase) CreateCommonForInternal(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.CreateCommonForInternal(request)
	fmt.Println(commons)
}

func (commonUsecase commonUsecase) CreateCommonForPrivate(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.CreateCommonForPrivate(request)
	fmt.Println(commons)
}

// UPDATE
func (commonUsecase commonUsecase) UpdateCommonForPublic(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.UpdateCommonForPublic(request)
	fmt.Println(commons)
}

func (commonUsecase commonUsecase) UpdateCommonForInternal(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.UpdateCommonForInternal(request)
	fmt.Println(commons)
}

func (commonUsecase commonUsecase) UpdateCommonForPrivate(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.UpdateCommonForPrivate(request)
	fmt.Println(commons)
}

// DELETE
func (commonUsecase commonUsecase) DeleteCommonForPublic(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.DeleteCommonForPublic(request)
	fmt.Println(commons)
}

func (commonUsecase commonUsecase) DeleteCommonForInternal(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.DeleteCommonForInternal(request)
	fmt.Println(commons)
}

func (commonUsecase commonUsecase) DeleteCommonForPrivate(request request.CommonRequest) {
	commons := commonUsecase.CommonRepository.DeleteCommonForPrivate(request)
	fmt.Println(commons)
}

func NewCommonUsecase(commonRepository repository.CommonRepository) CommonUsecase {
	return &commonUsecase{CommonRepository: commonRepository}
}
