package handlers

import (
	"github.com/demacedoleo/health-api/cmd/health/handlers/entities"
	"github.com/demacedoleo/health-api/cmd/health/handlers/presenter"
	"github.com/demacedoleo/health-api/internal/app/meetings"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MeetingsHandler interface {
	GetMeetings(c *gin.Context)
	CreateMeeting(c *gin.Context)
}

type meetingsHandler struct {
	service meetings.Service
}

func (m *meetingsHandler) GetMeetings(c *gin.Context) {
	//TODO: hard coded dates
	meetings, err := m.service.GetMeetings(c.Request.Context(), "", "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.Meetings(meetings))
}

func (m *meetingsHandler) CreateMeeting(c *gin.Context) {
	var meeting entities.Meeting
	if err := c.BindJSON(&meeting); err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
			Message:    "cannot bin json check required fields",
		})
		return
	}

	if err := m.service.CreateMeeting(c.Request.Context(), meetings.Meeting(meeting)); err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func NewMeetingsHandler(meetingsService meetings.Service) *meetingsHandler {
	return &meetingsHandler{service: meetingsService}
}
