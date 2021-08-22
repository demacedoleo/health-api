package handlers

import (
	"github.com/demacedoleo/health-api/cmd/health/handlers/presenter"
	"github.com/demacedoleo/health-api/internal/app/locations"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LocationsHandler interface {
	GetCities(c *gin.Context)
	GetStates(c *gin.Context)
}

type locationsHandler struct {
	service location.LocationsService
}

func (l *locationsHandler) GetCities(c *gin.Context) {
	stateID := c.Param("state_id")
	id, err := strconv.Atoi(stateID)

	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
			Message:    err.Error()})
		return
	}

	if id <= 0 {
		c.JSON(http.StatusBadRequest, presenter.Error{
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
			Message:    "invalid state id"})
		return
	}

	cities, err := l.service.GetCities(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    "cannot retrieve cities"})
		return
	}

	c.JSON(http.StatusOK, presenter.Cities(cities))
}

func (l *locationsHandler) GetStates(c *gin.Context) {
	states, err := l.service.GetStates(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    "cannot retrieve states"})
		return
	}

	c.JSON(http.StatusOK, presenter.States(states))
}

func NewLocationsHandler(service location.LocationsService) *locationsHandler {
	return &locationsHandler{service: service}
}
