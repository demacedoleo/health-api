package health

import (
	"github.com/demacedoleo/health-api/internal/app/company"
	"github.com/demacedoleo/health-api/internal/app/health"
	location "github.com/demacedoleo/health-api/internal/app/locations"
	"github.com/demacedoleo/health-api/internal/app/meetings"
	"github.com/demacedoleo/health-api/internal/platform/environment"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

type Dependencies struct {
	companyAdapter  mysql.Repository
	locationAdapter mysql.Repository
	healthAdapter   mysql.Repository
	meetingAdapter  mysql.Repository
	staffsAdapter   mysql.Repository
}

func BuildDependencies(env environment.Environment) (*Dependencies, error) {
	switch env {
	case environment.Production, environment.Beta, environment.Development:
		db := mysql.NewRepository(nil)

		return &Dependencies{
			companyAdapter:  company.NewCompanyAdapter(db),
			locationAdapter: location.NewLocationsAdapter(db),
			healthAdapter:   health.NewProvidersAdapter(db),
			meetingAdapter:  meetings.NewMeetingsAdapter(db),
			staffsAdapter:   company.NewStaffsAdapter(db),
		}, nil
	}

	return nil, nil
}
