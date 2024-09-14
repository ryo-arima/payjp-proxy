package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type TermRepository interface {
	BootstrapTermForDB(request request.TermRequest)(response response.TermResponse) 
	GetTermForPublic(request request.TermRequest)(response response.TermResponse) 
	GetTermForInternal(request request.TermRequest)(response response.TermResponse) 
	GetTermForPrivate(request request.TermRequest)(response response.TermResponse) 
	CreateTermForPublic(request request.TermRequest)   (response response.TermResponse) 
	CreateTermForInternal(request request.TermRequest) (response response.TermResponse)
	CreateTermForPrivate(request request.TermRequest)  (response response.TermResponse)
	UpdateTermForPublic(request request.TermRequest)   (response response.TermResponse)
	UpdateTermForInternal(request request.TermRequest) (response response.TermResponse)
	UpdateTermForPrivate(request request.TermRequest)  (response response.TermResponse)
	DeleteTermForPublic(request request.TermRequest)   (response response.TermResponse)
	DeleteTermForInternal(request request.TermRequest) (response response.TermResponse)
	DeleteTermForPrivate(request request.TermRequest)  (response response.TermResponse)
}

type termRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (termRepository termRepository) BootstrapTermForDB(request request.TermRequest) (response response.TermResponse) {
	fmt.Println("BootstrapTermForDB")
	return response
}

//GET
func (termRepository termRepository) GetTermForPublic(request request.TermRequest) (response response.TermResponse) {
	fmt.Println("GetTermForPublic")
	return response
}

func (termRepository termRepository) GetTermForInternal(request request.TermRequest) (response response.TermResponse ){
	fmt.Println("GetTermForInternal")
	return response
}

func (termRepository termRepository) GetTermForPrivate(request request.TermRequest) (response response.TermResponse) {
	fmt.Println("GetTermForPrivate")
	return response
}

//CREATE
func (termRepository termRepository) CreateTermForPublic(request request.TermRequest) (response response.TermResponse ){
	fmt.Println("CreateTermForPublic")
	return response
}

func (termRepository termRepository) CreateTermForInternal(request request.TermRequest) (response response.TermResponse) {
	fmt.Println("CreateTermForInternal()")
	return response
}

func (termRepository termRepository) CreateTermForPrivate(request request.TermRequest) (response response.TermResponse){
	fmt.Println("CreateTermForPrivate()")
	return response
}

//UPDATE
func (termRepository termRepository) UpdateTermForPublic(request request.TermRequest) (response response.TermResponse){
	fmt.Println("UpdateTermForPublic()")
	return response
}

func (termRepository termRepository) UpdateTermForInternal(request request.TermRequest) (response response.TermResponse){
	fmt.Println("UpdateTermForInternal")
	return response
}

func (termRepository termRepository) UpdateTermForPrivate(request request.TermRequest) (response response.TermResponse){
	fmt.Println("UpdateTermForPrivate")
	return response
}

//DELETE
func (termRepository termRepository) DeleteTermForPublic(request request.TermRequest) (response response.TermResponse){
	fmt.Println("DeleteTermForPublic")
	return response
}

func (termRepository termRepository) DeleteTermForInternal(request request.TermRequest) (response response.TermResponse ){
	fmt.Println("DeleteTermForInternal")
	return response
}

func (termRepository termRepository) DeleteTermForPrivate(request request.TermRequest) (response response.TermResponse){
	fmt.Println("DeleteTermForPrivate")
	return response
}

func NewTermRepository(conf config.BaseConfig) TermRepository {
	return &termRepository{BaseConfig: conf}
}