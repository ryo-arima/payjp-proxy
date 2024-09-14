package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type EventControllerForPrivate interface {
	GetEvents(c *gin.Context)
	CreateEvent(c *gin.Context)
	UpdateEvent(c *gin.Context)
	DeleteEvent(c *gin.Context)
}

type eventControllerForPrivate struct {
	EventRepository repository.EventRepository
}

func (eventController eventControllerForPrivate) GetEvents(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	res := eventController.EventRepository.GetEvents()
	c.JSON(http.StatusOK, res)
	return
}

func (eventController eventControllerForPrivate) CreateEvent(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	var eventModel model.Events
	res := eventController.EventRepository.CreateEvent(eventModel)
	c.JSON(http.StatusOK, res)
	return
}

func (eventController eventControllerForPrivate) UpdateEvent(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	var eventModel model.Events
	res := eventController.EventRepository.UpdateEvent(eventModel)
	c.JSON(http.StatusOK, res)
	return
}

func (eventController eventControllerForPrivate) DeleteEvent(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	var eventModel model.Events
	res := eventController.EventRepository.DeleteEvent(eventModel)
	c.JSON(http.StatusOK, res)
	return
}

func NewEventControllerForPrivate(eventRepository repository.EventRepository) EventControllerForPrivate {
	return &eventControllerForPrivate{EventRepository: eventRepository}
}
