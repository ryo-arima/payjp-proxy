package usecase

import (
	"fmt"

	"github.com/ryo-arima/payjp-proxy/pkg/client/repository"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
)

type EventUsecase interface {
	BootstrapEventForDB(request request.EventRequest)
	GetEventForPublic(request request.EventRequest)
	GetEventForInternal(request request.EventRequest)
	GetEventForPrivate(request request.EventRequest)
	CreateEventForPublic(request request.EventRequest)
	CreateEventForInternal(request request.EventRequest)
	CreateEventForPrivate(request request.EventRequest)
	UpdateEventForPublic(request request.EventRequest)
	UpdateEventForInternal(request request.EventRequest)
	UpdateEventForPrivate(request request.EventRequest)
	DeleteEventForPublic(request request.EventRequest)
	DeleteEventForInternal(request request.EventRequest)
	DeleteEventForPrivate(request request.EventRequest)
}

type eventUsecase struct {
	EventRepository   repository.EventRepository
}

//Bootstrap
func (eventUsecase eventUsecase) BootstrapEventForDB(request request.EventRequest) {
	events := eventUsecase.EventRepository.BootstrapEventForDB(request)
	fmt.Println(events)
}

//GET
func (eventUsecase eventUsecase) GetEventForPublic(request request.EventRequest) {
	events := eventUsecase.EventRepository.GetEventForPublic(request)
	fmt.Println(events)
}

func (eventUsecase eventUsecase) GetEventForInternal(request request.EventRequest) {
	events := eventUsecase.EventRepository.GetEventForInternal(request)
	fmt.Println(events)
}

func (eventUsecase eventUsecase) GetEventForPrivate(request request.EventRequest) {
	events := eventUsecase.EventRepository.GetEventForPrivate(request)
	fmt.Println(events)
}

//CREATE
func (eventUsecase eventUsecase) CreateEventForPublic(request request.EventRequest) {
	events := eventUsecase.EventRepository.CreateEventForPublic(request)
	fmt.Println(events)
}

func (eventUsecase eventUsecase) CreateEventForInternal(request request.EventRequest) {
	events := eventUsecase.EventRepository.CreateEventForInternal(request)
	fmt.Println(events)
}

func (eventUsecase eventUsecase) CreateEventForPrivate(request request.EventRequest) {
	events := eventUsecase.EventRepository.CreateEventForPrivate(request)
	fmt.Println(events)
}

//UPDATE
func (eventUsecase eventUsecase) UpdateEventForPublic(request request.EventRequest) {
	events := eventUsecase.EventRepository.UpdateEventForPublic(request)
	fmt.Println(events)
}

func (eventUsecase eventUsecase) UpdateEventForInternal(request request.EventRequest) {
	events := eventUsecase.EventRepository.UpdateEventForInternal(request)
	fmt.Println(events)
}

func (eventUsecase eventUsecase) UpdateEventForPrivate(request request.EventRequest) {
	events := eventUsecase.EventRepository.UpdateEventForPrivate(request)
	fmt.Println(events)
}

//DELETE
func (eventUsecase eventUsecase) DeleteEventForPublic(request request.EventRequest) {
	events := eventUsecase.EventRepository.DeleteEventForPublic(request)
	fmt.Println(events)
}

func (eventUsecase eventUsecase) DeleteEventForInternal(request request.EventRequest) {
	events := eventUsecase.EventRepository.DeleteEventForInternal(request)
	fmt.Println(events)
}

func (eventUsecase eventUsecase) DeleteEventForPrivate(request request.EventRequest) {
	events := eventUsecase.EventRepository.DeleteEventForPrivate(request)
	fmt.Println(events)
}

func NewEventUsecase(eventRepository repository.EventRepository) EventUsecase {
	return &eventUsecase{ EventRepository: eventRepository}
}