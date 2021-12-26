package handlers

import (
	"log"
	"net/http"

	"github.com/demacedoleo/health-api/cmd/health/handlers/entities"
	"github.com/demacedoleo/health-api/cmd/health/handlers/presenter"
	"github.com/demacedoleo/health-api/internal/app/company"
	"github.com/demacedoleo/health-api/internal/app/errors"
	"github.com/gin-gonic/gin"
)

func (ch *companyHandler) CreateStaff(c *gin.Context) {
	var s entities.Staff
	s.CompanyID = c.GetInt64("company_id")

	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
			Message:    "cannot bin json check required fields",
		})
		return
	}

	var staff company.Staff
	staff.CompanyID = s.CompanyID
	staff.Job = company.Charge(s.Job)
	staff.Person = company.Person(s.Person)
	staff.Contact = company.Contact(s.Contact)
	staff.Address = company.Address(s.Address)
	staff.Status = true
	staff.Color = s.Color

	staff.Schedule = make([]company.WorkDay, 0)
	staff.Registers = make([]company.Register, 0)

	for _, register := range s.Registers {
		staff.Registers = append(staff.Registers, company.Register{
			CompanyID: s.CompanyID,
			Document:  s.Person.Document,
			Register:  register,
		})
	}

	for _, workday := range s.Schedule {
		if workday != (entities.WorkDay{}) {
			staff.Schedule = append(staff.Schedule, company.WorkDay{
				Document:  s.Person.Document,
				CompanyID: s.CompanyID,
				WeekDay:   workday.WeekDay,
				Start:     workday.Start,
				End:       workday.End,
			})
		}
	}

	if err := ch.company.CreateStaff(c.Request.Context(), staff); err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (ch *companyHandler) FindStaffs(c *gin.Context) {
	companyID := c.GetInt64("company_id")

	staffs, err := ch.company.FindStaffs(c.Request.Context(), companyID)
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

	c.JSON(http.StatusOK, presenter.Staffs(staffs))
}

func (ch *companyHandler) FindProfessional(c *gin.Context) {
	companyID := c.GetInt64("company_id")

	staffs, err := ch.company.FindStaffs(c.Request.Context(), companyID)
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

	c.JSON(http.StatusOK, presenter.Staffs(staffs))
}
