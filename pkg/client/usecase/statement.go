package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type StatementUsecase interface {
	BootstrapStatementForDB(request request.StatementRequest)
	GetStatementForPublic(request request.StatementRequest)
	GetStatementForInternal(request request.StatementRequest)
	GetStatementForPrivate(request request.StatementRequest)
	CreateStatementForPublic(request request.StatementRequest)
	CreateStatementForInternal(request request.StatementRequest)
	CreateStatementForPrivate(request request.StatementRequest)
	UpdateStatementForPublic(request request.StatementRequest)
	UpdateStatementForInternal(request request.StatementRequest)
	UpdateStatementForPrivate(request request.StatementRequest)
	DeleteStatementForPublic(request request.StatementRequest)
	DeleteStatementForInternal(request request.StatementRequest)
	DeleteStatementForPrivate(request request.StatementRequest)
}

type statementUsecase struct {
	StatementRepository   repository.StatementRepository
}

//Bootstrap
func (statementUsecase statementUsecase) BootstrapStatementForDB(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.BootstrapStatementForDB(request)
	fmt.Println(statements)
}

//GET
func (statementUsecase statementUsecase) GetStatementForPublic(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.GetStatementForPublic(request)
	fmt.Println(statements)
}

func (statementUsecase statementUsecase) GetStatementForInternal(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.GetStatementForInternal(request)
	fmt.Println(statements)
}

func (statementUsecase statementUsecase) GetStatementForPrivate(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.GetStatementForPrivate(request)
	fmt.Println(statements)
}

//CREATE
func (statementUsecase statementUsecase) CreateStatementForPublic(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.CreateStatementForPublic(request)
	fmt.Println(statements)
}

func (statementUsecase statementUsecase) CreateStatementForInternal(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.CreateStatementForInternal(request)
	fmt.Println(statements)
}

func (statementUsecase statementUsecase) CreateStatementForPrivate(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.CreateStatementForPrivate(request)
	fmt.Println(statements)
}

//UPDATE
func (statementUsecase statementUsecase) UpdateStatementForPublic(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.UpdateStatementForPublic(request)
	fmt.Println(statements)
}

func (statementUsecase statementUsecase) UpdateStatementForInternal(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.UpdateStatementForInternal(request)
	fmt.Println(statements)
}

func (statementUsecase statementUsecase) UpdateStatementForPrivate(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.UpdateStatementForPrivate(request)
	fmt.Println(statements)
}

//DELETE
func (statementUsecase statementUsecase) DeleteStatementForPublic(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.DeleteStatementForPublic(request)
	fmt.Println(statements)
}

func (statementUsecase statementUsecase) DeleteStatementForInternal(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.DeleteStatementForInternal(request)
	fmt.Println(statements)
}

func (statementUsecase statementUsecase) DeleteStatementForPrivate(request request.StatementRequest) {
	statements := statementUsecase.StatementRepository.DeleteStatementForPrivate(request)
	fmt.Println(statements)
}

func NewStatementUsecase(statementRepository repository.StatementRepository) StatementUsecase {
	return &statementUsecase{ StatementRepository: statementRepository}
}