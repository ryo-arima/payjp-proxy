package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type CommonRepository interface {
	BootstrapCommonForDB(request request.CommonRequest)(response response.CommonResponse) 
	GetCommonForPublic(request request.CommonRequest)(response response.CommonResponse) 
	GetCommonForInternal(request request.CommonRequest)(response response.CommonResponse) 
	GetCommonForPrivate(request request.CommonRequest)(response response.CommonResponse) 
	CreateCommonForPublic(request request.CommonRequest)   (response response.CommonResponse) 
	CreateCommonForInternal(request request.CommonRequest) (response response.CommonResponse)
	CreateCommonForPrivate(request request.CommonRequest)  (response response.CommonResponse)
	UpdateCommonForPublic(request request.CommonRequest)   (response response.CommonResponse)
	UpdateCommonForInternal(request request.CommonRequest) (response response.CommonResponse)
	UpdateCommonForPrivate(request request.CommonRequest)  (response response.CommonResponse)
	DeleteCommonForPublic(request request.CommonRequest)   (response response.CommonResponse)
	DeleteCommonForInternal(request request.CommonRequest) (response response.CommonResponse)
	DeleteCommonForPrivate(request request.CommonRequest)  (response response.CommonResponse)
}

type commonRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (commonRepository commonRepository) BootstrapCommonForDB(request request.CommonRequest) (response response.CommonResponse) {
	fmt.Println("BootstrapCommonForDB")
	return response
}

//GET
func (commonRepository commonRepository) GetCommonForPublic(request request.CommonRequest) (response response.CommonResponse) {
	fmt.Println("GetCommonForPublic")
	return response
}

func (commonRepository commonRepository) GetCommonForInternal(request request.CommonRequest) (response response.CommonResponse ){
	fmt.Println("GetCommonForInternal")
	return response
}

func (commonRepository commonRepository) GetCommonForPrivate(request request.CommonRequest) (response response.CommonResponse) {
	fmt.Println("GetCommonForPrivate")
	return response
}

//CREATE
func (commonRepository commonRepository) CreateCommonForPublic(request request.CommonRequest) (response response.CommonResponse ){
	fmt.Println("CreateCommonForPublic")
	return response
}

func (commonRepository commonRepository) CreateCommonForInternal(request request.CommonRequest) (response response.CommonResponse) {
	fmt.Println("CreateCommonForInternal()")
	return response
}

func (commonRepository commonRepository) CreateCommonForPrivate(request request.CommonRequest) (response response.CommonResponse){
	fmt.Println("CreateCommonForPrivate()")
	return response
}

//UPDATE
func (commonRepository commonRepository) UpdateCommonForPublic(request request.CommonRequest) (response response.CommonResponse){
	fmt.Println("UpdateCommonForPublic()")
	return response
}

func (commonRepository commonRepository) UpdateCommonForInternal(request request.CommonRequest) (response response.CommonResponse){
	fmt.Println("UpdateCommonForInternal")
	return response
}

func (commonRepository commonRepository) UpdateCommonForPrivate(request request.CommonRequest) (response response.CommonResponse){
	fmt.Println("UpdateCommonForPrivate")
	return response
}

//DELETE
func (commonRepository commonRepository) DeleteCommonForPublic(request request.CommonRequest) (response response.CommonResponse){
	fmt.Println("DeleteCommonForPublic")
	return response
}

func (commonRepository commonRepository) DeleteCommonForInternal(request request.CommonRequest) (response response.CommonResponse ){
	fmt.Println("DeleteCommonForInternal")
	return response
}

func (commonRepository commonRepository) DeleteCommonForPrivate(request request.CommonRequest) (response response.CommonResponse){
	fmt.Println("DeleteCommonForPrivate")
	return response
}

func NewCommonRepository(conf config.BaseConfig) CommonRepository {
	return &commonRepository{BaseConfig: conf}
}