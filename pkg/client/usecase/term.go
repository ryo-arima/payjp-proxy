package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type TermUsecase interface {
	BootstrapTermForDB(request request.TermRequest)
	GetTermForInternal(request request.TermRequest)
	GetTermForPrivate(request request.TermRequest)
	CreateTermForInternal(request request.TermRequest)
	CreateTermForPrivate(request request.TermRequest)
	UpdateTermForInternal(request request.TermRequest)
	UpdateTermForPrivate(request request.TermRequest)
	DeleteTermForInternal(request request.TermRequest)
	DeleteTermForPrivate(request request.TermRequest)
}

type termUsecase struct {
	TermRepository repository.TermRepository
}

// Bootstrap
func (termUsecase termUsecase) BootstrapTermForDB(request request.TermRequest) {
	terms := termUsecase.TermRepository.BootstrapTermForDB(request)
	fmt.Println(terms)
}

// GET
func (termUsecase termUsecase) GetTermForInternal(request request.TermRequest) {
	terms := termUsecase.TermRepository.GetTermForInternal(request)
	fmt.Println(terms)
}

func (termUsecase termUsecase) GetTermForPrivate(request request.TermRequest) {
	terms := termUsecase.TermRepository.GetTermForPrivate(request)
	fmt.Println(terms)
}

// CREATE

func (termUsecase termUsecase) CreateTermForInternal(request request.TermRequest) {
	terms := termUsecase.TermRepository.CreateTermForInternal(request)
	fmt.Println(terms)
}

func (termUsecase termUsecase) CreateTermForPrivate(request request.TermRequest) {
	terms := termUsecase.TermRepository.CreateTermForPrivate(request)
	fmt.Println(terms)
}

//UPDATE

func (termUsecase termUsecase) UpdateTermForInternal(request request.TermRequest) {
	terms := termUsecase.TermRepository.UpdateTermForInternal(request)
	fmt.Println(terms)
}

func (termUsecase termUsecase) UpdateTermForPrivate(request request.TermRequest) {
	terms := termUsecase.TermRepository.UpdateTermForPrivate(request)
	fmt.Println(terms)
}

//DELETE

func (termUsecase termUsecase) DeleteTermForInternal(request request.TermRequest) {
	terms := termUsecase.TermRepository.DeleteTermForInternal(request)
	fmt.Println(terms)
}

func (termUsecase termUsecase) DeleteTermForPrivate(request request.TermRequest) {
	terms := termUsecase.TermRepository.DeleteTermForPrivate(request)
	fmt.Println(terms)
}

func NewTermUsecase(termRepository repository.TermRepository) TermUsecase {
	return &termUsecase{TermRepository: termRepository}
}
