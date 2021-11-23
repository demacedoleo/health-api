package handlers

import (
	"github.com/demacedoleo/health-api/cmd/health/handlers/entities"
	"github.com/demacedoleo/health-api/cmd/health/handlers/presenter"
	"github.com/demacedoleo/health-api/internal/app/company"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CompanyHandler interface {
	ReadCompany(c *gin.Context)
	CreateCompany(c *gin.Context)
	RolesHandler
	StaffsHandler
}

type StaffsHandler interface {
	CreateStaffs(c *gin.Context)
	FindStaffs(c *gin.Context)
}

type RolesHandler interface {
	ReadRoles(c *gin.Context)
	CreateRole(c *gin.Context)
}

type ModalitiesHandler interface {
	GetModalities(c *gin.Context)
	CreateModality(c *gin.Context)
}

type companyHandler struct {
	company company.Service
}

func (ch *companyHandler) ReadRoles(c *gin.Context) {
	companyID := c.GetInt64("company_id")
	roles, err := ch.company.GetRoles(c.Request.Context(), companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.Roles(roles))
}

func (ch *companyHandler) CreateRole(c *gin.Context) {
	var role entities.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
			Message:    "cannot bin json check required fields",
		})
		return
	}

	if err := ch.company.CreateRole(c.Request.Context(), company.Role(role)); err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (ch *companyHandler) ReadCompany(c *gin.Context) {

	companyID := c.Param("company_id")
	tokenCompanyID := c.GetInt64("company_id")

	id, err := strconv.ParseInt(companyID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
			Message:    err.Error(),
		})
		return
	}

	if id != tokenCompanyID {
		c.JSON(http.StatusForbidden, presenter.Error{
			StatusCode: http.StatusForbidden,
			Code:       "forbidden",
			Message:    "mismatch company token",
		})
		return
	}

	cia, err := ch.company.GetCompany(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.Company(cia))
}

func (ch *companyHandler) CreateCompany(c *gin.Context) {
	var co entities.Company
	if err := c.ShouldBindJSON(&co); err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
			Message:    "cannot bin json check required fields",
		})
		return
	}

	if err := ch.company.CreateCompany(c.Request.Context(), company.Company(co)); err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (ch *companyHandler) GetModalities(c *gin.Context) {
	companyID := c.GetInt64("company_id")
	modalities, err := ch.company.GetModalities(c.Request.Context(), companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.Modalities(modalities))
}

func (ch *companyHandler) CreateModality(c *gin.Context) {
	var modality entities.Modality
	if err := c.ShouldBindJSON(&modality); err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
			Message:    "cannot bin json check required fields",
		})
		return
	}

	if err := ch.company.CreateModality(c.Request.Context(), company.Modality(modality)); err != nil {
		c.JSON(http.StatusInternalServerError, presenter.Error{
			StatusCode: http.StatusInternalServerError,
			Code:       "internal_error",
			Message:    err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func NewCompanyHandler(companyService company.Service) *companyHandler {
	return &companyHandler{company: companyService}
}
