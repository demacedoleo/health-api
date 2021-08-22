package handlers

import (
	"github.com/demacedoleo/health-api/cmd/health/handlers/entities"
	"github.com/demacedoleo/health-api/cmd/health/handlers/presenter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication(c *gin.Context) {
	var auth entities.HeaderAuth
	err := c.ShouldBindHeader(&auth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, presenter.Error{
			StatusCode: http.StatusUnauthorized,
			Code:       "unauthorized",
			Message:    "you must be authenticated",
		})
		c.Abort()
		return
	}

	// TODO: check in repository the authentication
	// the user only operates over some company 1

	c.Set("company_id", int64(1))
	c.Next()
}
