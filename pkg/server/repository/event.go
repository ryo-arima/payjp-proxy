package repository

import (
	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
)

type EventRepository interface {
	GetEvents() []model.Events
	CreateEvent(event model.Events) model.Events
	UpdateEvent(event model.Events) model.Events
	DeleteEvent(event model.Events) model.Events
}

type eventRepository struct {
	BaseConfig config.BaseConfig
}

func (eventRepository eventRepository) GetEvents() []model.Events {
	var events []model.Events
	return events
}

func (eventRepository eventRepository) CreateEvent(event model.Events) model.Events {
	return event
}

func (eventRepository eventRepository) UpdateEvent(event model.Events) model.Events {
	return event
}

func (eventRepository eventRepository) DeleteEvent(event model.Events) model.Events {
	return event
}

func NewEventRepository(conf config.BaseConfig) EventRepository {
	return &eventRepository{BaseConfig: conf}
}
