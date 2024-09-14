package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/model"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type EventControllerForInternal interface {
	GetEvents(c *gin.Context)
	CreateEvent(c *gin.Context)
	UpdateEvent(c *gin.Context)
	DeleteEvent(c *gin.Context)
}

type eventControllerForInternal struct {
	EventRepository repository.EventRepository
}

func (eventController eventControllerForInternal) GetEvents(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	eventController.EventRepository.GetEvents()
	c.JSON(http.StatusOK, &response.EventResponse{Code: "SERVER_CONTROLLER_GET__FOR__002", Message: "Success", Events: []response.Event{}})
	return
}

func (eventController eventControllerForInternal) CreateEvent(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	var eventModel model.Events
	eventController.EventRepository.CreateEvent(eventModel)
	c.JSON(http.StatusOK, &response.EventResponse{Code: "SERVER_CONTROLLER_CREATE__FOR__002", Message: "Success", Events: []response.Event{}})
	return
}

func (eventController eventControllerForInternal) UpdateEvent(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	var eventModel model.Events
	eventController.EventRepository.UpdateEvent(eventModel)
	c.JSON(http.StatusOK, &response.EventResponse{Code: "SERVER_CONTROLLER_UPDATE__FOR__002", Message: "Success", Events: []response.Event{}})
	return
}

func (eventController eventControllerForInternal) DeleteEvent(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	var eventModel model.Events
	eventController.EventRepository.DeleteEvent(eventModel)
	c.JSON(http.StatusOK, &response.EventResponse{Code: "SERVER_CONTROLLER_DELETE__FOR__002", Message: "Success", Events: []response.Event{}})
	return
}

func NewEventControllerForInternal(eventRepository repository.EventRepository) EventControllerForInternal {
	return &eventControllerForInternal{EventRepository: eventRepository}
}
