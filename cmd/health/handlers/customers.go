package handlers

import (
	"fmt"
	"github.com/demacedoleo/health-api/cmd/health/handlers/entities"
	"github.com/demacedoleo/health-api/cmd/health/handlers/presenter"
	"github.com/demacedoleo/health-api/internal/app/company"
	"github.com/demacedoleo/health-api/internal/app/errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func (ch *companyHandler) CreateCustomers(c *gin.Context) {
	var s entities.Customer
	s.CompanyID = c.GetInt64("company_id")

	if err := c.ShouldBindJSON(&s); err != nil {
		b, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println(string(b))

		c.JSON(http.StatusBadRequest, presenter.Error{
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
			Message:    "cannot bin json check required fields",
		})
		return
	}

	customerHealth := company.CustomerHealth{
		Person: company.Person(s.Person),
		Health: company.Health(s.Insurance),
	}

	customerHealth.Person.CompanyID = s.CompanyID
	customerHealth.Health.CompanyID = s.CompanyID
	customerHealth.Health.Document = s.Person.Document

	if err := ch.company.CreateCustomers(c.Request.Context(), customerHealth); err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (ch *companyHandler) FindCustomers(c *gin.Context) {
	companyID := c.GetInt64("company_id")

	customers, err := ch.company.FindCustomers(c.Request.Context(), companyID)
	if errors.Is(err, company.ErrInternalScan) {
		log.Println(errors.Format(err))

		c.JSON(http.StatusInternalServerError, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    err.Error(),
		})
		return
	}

	if errors.Is(err, company.ErrNotFound) {
		err.Error()

		c.JSON(http.StatusNotFound, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "not_found",
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.Customers(customers))
}
