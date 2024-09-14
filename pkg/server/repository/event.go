package repository

import (
	"time"

	"github.com/ryo-arima/payjp-proxy/pkg/config"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"gorm.io/gorm"
)

type EventRepository interface {
	GetEvents() []model.Events
	CreateEvent(event model.Events) *gorm.DB
	UpdateEvent(event model.Events) *gorm.DB
	DeleteEvent(uuid string) *gorm.DB
}

type eventRepository struct {
	BaseConfig config.BaseConfig
}


func (eventRepository eventRepository) GetEvents() []model.Events {
	var events []model.Events
	eventRepository.BaseConfig.DBConnection.Find(&events)
	return events
}

func (eventRepository eventRepository) CreateEvent(event model.Events) *gorm.DB {
	results := eventRepository.BaseConfig.DBConnection.Create(&event)
	return results
}

func (eventRepository eventRepository) UpdateEvent(event model.Events) *gorm.DB {
	results := eventRepository.BaseConfig.DBConnection.Model(model.Events{}).Where("uuid = ?", event.UUID).Updates(event)
	return results
}

func (eventRepository eventRepository) DeleteEvent(uuid string) *gorm.DB {
	results := eventRepository.BaseConfig.DBConnection.Model(model.Events{}).Where("uuid = ?", uuid).Update("deleted_at", time.Now())
	return results
}

func NewEventRepository(conf config.BaseConfig) EventRepository {
	return &eventRepository{BaseConfig: conf}
}