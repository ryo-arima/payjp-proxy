package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/request"
	"github.com/ryo-arima/payjp-proxy/pkg/entity/response"
	"github.com/ryo-arima/payjp-proxy/pkg/server/repository"
)

type EventControllerForPublic interface {
	GetEvents(c *gin.Context)
}

type eventControllerForPublic struct {
	EventRepository repository.EventRepository
}

func (eventController eventControllerForPublic) GetEvents(c *gin.Context) {
	var eventRequest request.EventRequest
	if err := c.Bind(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, &response.EventResponse{Code: "SERVER_CONTROLLER_GET__FOR__001", Message: err.Error(), Events: []response.Event{}})
		return
	}
	res := eventController.EventRepository.GetEvents()
	c.JSON(http.StatusOK, res)
	return
}


func NewEventControllerForPublic(eventRepository repository.EventRepository) EventControllerForPublic {
	return &eventControllerForPublic{ EventRepository: eventRepository}
}