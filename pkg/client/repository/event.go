package repository

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
)

type EventRepository interface {
	BootstrapEventForDB(request request.EventRequest)(response response.EventResponse) 
	GetEventForPublic(request request.EventRequest)(response response.EventResponse) 
	GetEventForInternal(request request.EventRequest)(response response.EventResponse) 
	GetEventForPrivate(request request.EventRequest)(response response.EventResponse) 
	CreateEventForPublic(request request.EventRequest)   (response response.EventResponse) 
	CreateEventForInternal(request request.EventRequest) (response response.EventResponse)
	CreateEventForPrivate(request request.EventRequest)  (response response.EventResponse)
	UpdateEventForPublic(request request.EventRequest)   (response response.EventResponse)
	UpdateEventForInternal(request request.EventRequest) (response response.EventResponse)
	UpdateEventForPrivate(request request.EventRequest)  (response response.EventResponse)
	DeleteEventForPublic(request request.EventRequest)   (response response.EventResponse)
	DeleteEventForInternal(request request.EventRequest) (response response.EventResponse)
	DeleteEventForPrivate(request request.EventRequest)  (response response.EventResponse)
}

type eventRepository struct {
	BaseConfig   config.BaseConfig
}

//Bootstrap
func (eventRepository eventRepository) BootstrapEventForDB(request request.EventRequest) (response response.EventResponse) {
	fmt.Println("BootstrapEventForDB")
	return response
}

//GET
func (eventRepository eventRepository) GetEventForPublic(request request.EventRequest) (response response.EventResponse) {
	fmt.Println("GetEventForPublic")
	return response
}

func (eventRepository eventRepository) GetEventForInternal(request request.EventRequest) (response response.EventResponse ){
	fmt.Println("GetEventForInternal")
	return response
}

func (eventRepository eventRepository) GetEventForPrivate(request request.EventRequest) (response response.EventResponse) {
	fmt.Println("GetEventForPrivate")
	return response
}

//CREATE
func (eventRepository eventRepository) CreateEventForPublic(request request.EventRequest) (response response.EventResponse ){
	fmt.Println("CreateEventForPublic")
	return response
}

func (eventRepository eventRepository) CreateEventForInternal(request request.EventRequest) (response response.EventResponse) {
	fmt.Println("CreateEventForInternal()")
	return response
}

func (eventRepository eventRepository) CreateEventForPrivate(request request.EventRequest) (response response.EventResponse){
	fmt.Println("CreateEventForPrivate()")
	return response
}

//UPDATE
func (eventRepository eventRepository) UpdateEventForPublic(request request.EventRequest) (response response.EventResponse){
	fmt.Println("UpdateEventForPublic()")
	return response
}

func (eventRepository eventRepository) UpdateEventForInternal(request request.EventRequest) (response response.EventResponse){
	fmt.Println("UpdateEventForInternal")
	return response
}

func (eventRepository eventRepository) UpdateEventForPrivate(request request.EventRequest) (response response.EventResponse){
	fmt.Println("UpdateEventForPrivate")
	return response
}

//DELETE
func (eventRepository eventRepository) DeleteEventForPublic(request request.EventRequest) (response response.EventResponse){
	fmt.Println("DeleteEventForPublic")
	return response
}

func (eventRepository eventRepository) DeleteEventForInternal(request request.EventRequest) (response response.EventResponse ){
	fmt.Println("DeleteEventForInternal")
	return response
}

func (eventRepository eventRepository) DeleteEventForPrivate(request request.EventRequest) (response response.EventResponse){
	fmt.Println("DeleteEventForPrivate")
	return response
}

func NewEventRepository(conf config.BaseConfig) EventRepository {
	return &eventRepository{BaseConfig: conf}
}