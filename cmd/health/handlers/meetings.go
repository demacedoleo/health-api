package handlers

import (
	"github.com/demacedoleo/health-api/cmd/health/handlers/entities"
	"github.com/demacedoleo/health-api/cmd/health/handlers/presenter"
	"github.com/demacedoleo/health-api/internal/app/errors"
	"github.com/demacedoleo/health-api/internal/app/meetings"
	"github.com/gin-gonic/gin"
	"log"
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
	companyID := c.GetInt64("company_id")

	meets, err := m.service.GetMeetings(c.Request.Context(), companyID, "", "")
	if errors.Is(err, meetings.ErrInternalScan) {
		log.Println(errors.Format(err))

		c.JSON(http.StatusInternalServerError, presenter.Error{
			Message: err.Error(),
		})
		return
	}

	if errors.Is(err, meetings.ErrNotFound) {
		log.Println(errors.Format(err))

		c.JSON(http.StatusNotFound, presenter.Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.Meetings(meets))
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
