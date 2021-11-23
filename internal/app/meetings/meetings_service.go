package meetings

import (
	"context"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

type Service interface {
	Scheduler
}

type Scheduler interface {
	GetMeetings(ctx context.Context, companyID int64, startTime, endTime string) ([]Meeting, error)
	CreateMeeting(ctx context.Context, meeting Meeting) error
}

type service struct {
	adapter
}

func (s *service) GetMeetings(ctx context.Context, companyID int64, startTime, endTime string) ([]Meeting, error) {
	startTime, endTime = CalculateWeek()
	return s.adapter.GetMeetings(ctx, companyID, startTime, endTime)
}

func (s *service) CreateMeeting(ctx context.Context, meeting Meeting) error {
	return s.adapter.CreateMeeting(ctx, meeting)
}

func NewMeetingService(repository mysql.Repository) *service {
	return &service{adapter{repository}}
}
