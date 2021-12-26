package health

import (
	"github.com/demacedoleo/health-api/cmd/health/handlers"
	"github.com/demacedoleo/health-api/internal/app/company"
	health_service "github.com/demacedoleo/health-api/internal/app/health"
	"github.com/demacedoleo/health-api/internal/app/locations"
	"github.com/demacedoleo/health-api/internal/app/meetings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Build(dep *Dependencies) *gin.Engine {
	r := gin.Default()
	r.Use(cors())

	// services
	companyService := company.NewCompanyService(dep.companyAdapter)
	locationsService := location.NewLocationsService(dep.locationAdapter)
	healthService := health_service.NewHealthService(dep.healthAdapter)
	meetingsService := meetings.NewMeetingService(dep.meetingAdapter)

	// controller adapters
	locationsHandler := handlers.NewLocationsHandler(locationsService)
	companyHandler := handlers.NewCompanyHandler(companyService)
	healthHandler := handlers.NewHealthHandler(healthService)
	meetingsHandler := handlers.NewMeetingsHandler(meetingsService)

	// app health endpoint
	r.GET("/ping", handlers.Ping)

	health := r.Group("/health")
	health.Use(handlers.Authentication)
	{
		health.POST("/scrapper", handlers.ScrapperAFJP)

		// locations endpoints
		health.GET("/locations/states", locationsHandler.GetStates)
		health.GET("/locations/state/:state_id/cities", locationsHandler.GetCities)

		// company information
		health.GET("/companies/:company_id", companyHandler.ReadCompany)
		health.POST("/companies", companyHandler.CreateCompany)

		// company modalities
		health.GET("/company/modalities", companyHandler.GetModalities)
		health.POST("/company/modalities", companyHandler.CreateModality)

		// company staff's roles endpoints
		health.GET("/company/roles", companyHandler.ReadRoles)
		health.POST("/company/roles", companyHandler.CreateRole)

		health.GET("/company/staffs", companyHandler.FindStaffs)
		health.POST("/company/staffs", companyHandler.CreateStaff)

		health.GET("/company/customers", companyHandler.FindCustomers)
		health.POST("/company/customers", companyHandler.CreateCustomers)

		// health insurances providers
		health.GET("/insurances", healthHandler.GetProviders)
		health.POST("/insurances", healthHandler.CreateProviders)

		// meetings
		health.GET("/meetings", meetingsHandler.GetMeetings)
		health.POST("/meetings", meetingsHandler.CreateMeeting)
		health.PUT("/meetings", meetingsHandler.CreateMeeting)
	}

	return r
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Set("content-type", "application/json")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
