package handlers

import (
	"errors"
	"github.com/demacedoleo/health-api/internal/app/scrapper"
	"github.com/demacedoleo/health-api/internal/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ScrapperAFJP(c *gin.Context) {
	var data client.ScrapData
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("invalid scrap data"))
		return
	}

	userAFJP, err := scrapper.ScrapAFJP(data)
	if  err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, userAFJP)
}
