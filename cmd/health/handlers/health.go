package handlers

import (
	"github.com/demacedoleo/health-api/cmd/health/handlers/entities"
	"github.com/demacedoleo/health-api/cmd/health/handlers/presenter"
	"github.com/demacedoleo/health-api/internal/app/health"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler interface {
	GetProviders(c *gin.Context)
	CreateProviders(c *gin.Context)
}

type healthHandler struct {
	health health.Service
}

func (h *healthHandler) GetProviders(c *gin.Context) {
	companyID := c.GetInt64("company_id")
	providers, err := h.health.GetProviders(c.Request.Context(), companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.Providers(providers))
}

func (h *healthHandler) CreateProviders(c *gin.Context) {
	var provider entities.Provider
	if err := c.ShouldBindJSON(&provider); err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{StatusCode: http.StatusBadRequest, Code: "bad_request", Message: "cannot bin json check required fields"})
		return
	}

	if err := h.health.CreateProvider(c.Request.Context(), health.Provider(provider)); err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func NewHealthHandler(healthService health.Service) *healthHandler {
	return &healthHandler{health: healthService}
}
