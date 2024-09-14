package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type StatementRepository interface {
	BootstrapStatementForDB(request request.StatementRequest)(response response.StatementResponse) 
	GetStatementForPublic(request request.StatementRequest)(response response.StatementResponse) 
	GetStatementForInternal(request request.StatementRequest)(response response.StatementResponse) 
	GetStatementForPrivate(request request.StatementRequest)(response response.StatementResponse) 
	CreateStatementForPublic(request request.StatementRequest)   (response response.StatementResponse) 
	CreateStatementForInternal(request request.StatementRequest) (response response.StatementResponse)
	CreateStatementForPrivate(request request.StatementRequest)  (response response.StatementResponse)
	UpdateStatementForPublic(request request.StatementRequest)   (response response.StatementResponse)
	UpdateStatementForInternal(request request.StatementRequest) (response response.StatementResponse)
	UpdateStatementForPrivate(request request.StatementRequest)  (response response.StatementResponse)
	DeleteStatementForPublic(request request.StatementRequest)   (response response.StatementResponse)
	DeleteStatementForInternal(request request.StatementRequest) (response response.StatementResponse)
	DeleteStatementForPrivate(request request.StatementRequest)  (response response.StatementResponse)
}

type statementRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (statementRepository statementRepository) BootstrapStatementForDB(request request.StatementRequest) (response response.StatementResponse) {
	fmt.Println("BootstrapStatementForDB")
	return response
}

//GET
func (statementRepository statementRepository) GetStatementForPublic(request request.StatementRequest) (response response.StatementResponse) {
	fmt.Println("GetStatementForPublic")
	return response
}

func (statementRepository statementRepository) GetStatementForInternal(request request.StatementRequest) (response response.StatementResponse ){
	fmt.Println("GetStatementForInternal")
	return response
}

func (statementRepository statementRepository) GetStatementForPrivate(request request.StatementRequest) (response response.StatementResponse) {
	fmt.Println("GetStatementForPrivate")
	return response
}

//CREATE
func (statementRepository statementRepository) CreateStatementForPublic(request request.StatementRequest) (response response.StatementResponse ){
	fmt.Println("CreateStatementForPublic")
	return response
}

func (statementRepository statementRepository) CreateStatementForInternal(request request.StatementRequest) (response response.StatementResponse) {
	fmt.Println("CreateStatementForInternal()")
	return response
}

func (statementRepository statementRepository) CreateStatementForPrivate(request request.StatementRequest) (response response.StatementResponse){
	fmt.Println("CreateStatementForPrivate()")
	return response
}

//UPDATE
func (statementRepository statementRepository) UpdateStatementForPublic(request request.StatementRequest) (response response.StatementResponse){
	fmt.Println("UpdateStatementForPublic()")
	return response
}

func (statementRepository statementRepository) UpdateStatementForInternal(request request.StatementRequest) (response response.StatementResponse){
	fmt.Println("UpdateStatementForInternal")
	return response
}

func (statementRepository statementRepository) UpdateStatementForPrivate(request request.StatementRequest) (response response.StatementResponse){
	fmt.Println("UpdateStatementForPrivate")
	return response
}

//DELETE
func (statementRepository statementRepository) DeleteStatementForPublic(request request.StatementRequest) (response response.StatementResponse){
	fmt.Println("DeleteStatementForPublic")
	return response
}

func (statementRepository statementRepository) DeleteStatementForInternal(request request.StatementRequest) (response response.StatementResponse ){
	fmt.Println("DeleteStatementForInternal")
	return response
}

func (statementRepository statementRepository) DeleteStatementForPrivate(request request.StatementRequest) (response response.StatementResponse){
	fmt.Println("DeleteStatementForPrivate")
	return response
}

func NewStatementRepository(conf config.BaseConfig) StatementRepository {
	return &statementRepository{BaseConfig: conf}
}