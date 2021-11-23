package meetings

import (
	"context"
	"errors"

	oops "github.com/demacedoleo/health-api/internal/app/errors"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

var (
	ErrInternalScan = errors.New("errors mapping meetings")
	ErrNotFound     = errors.New("not found meetings")
)

type adapter struct {
	mysql.Repository
}

func (a *adapter) GetMeetings(ctx context.Context, companyID int64, startTime, endTime string) ([]Meeting, error) {
	result, err := a.Fetch(ctx, mysql.Statements.Selects.Meetings, companyID, startTime, endTime)
	if err != nil {
		return nil, err
	}

	var meetings []Meeting
	for result.Next() {
		var meeting Meeting

		if err := result.Scan(&meeting.ID, &meeting.EventID, &meeting.CompanyID, &meeting.Subject,
			&meeting.MeetStatus, &meeting.StartTime, &meeting.EndTime, &meeting.Modality, &meeting.AttendantDocument,
			&meeting.AttendantPhone, &meeting.ResourceID, &meeting.CreatedAt, &meeting.UpdatedAt); err != nil {
			return nil, oops.NewError(ErrInternalScan).AddStack(err)
		}

		meetings = append(meetings, meeting)
	}

	if len(meetings) == 0 {
		return nil, oops.NewError(ErrNotFound)
	}

	return meetings, nil
}

func (a *adapter) CreateMeeting(ctx context.Context, meeting Meeting) error {
	_, err := a.Repository.Save(ctx, mysql.Statements.Inserts.Meeting, meeting.ToString())
	return err
}

func NewMeetingsAdapter(repository mysql.Repository) *adapter {
	return &adapter{repository}
}
