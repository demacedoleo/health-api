package meetings

import (
	"context"
	"errors"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

var (
	ErrMappingMeetings  = errors.New("error mapping meetings")
	ErrNotFoundMeetings = errors.New("not found meetings")
)

type adapter struct {
	mysql.Repository
}

func (a *adapter) GetMeetings(ctx context.Context, startTime, endTime string) ([]Meeting, error) {
	result, err := a.Fetch(ctx, mysql.Statements.Selects.Meetings, startTime, endTime)
	if err != nil {
		return nil, err
	}

	var meetings []Meeting
	for result.Next() {
		var meeting Meeting

		if err := result.Scan(&meeting.ID, &meeting.EventID, &meeting.CompanyID, &meeting.Subject,
			&meeting.MeetStatus, &meeting.StartTime, &meeting.EndTime, &meeting.ModalityID, &meeting.AttendantDocument,
			&meeting.AttendantPhone, &meeting.ResourceID, &meeting.CreatedAt, &meeting.UpdatedAt); err != nil {
			return nil, ErrMappingMeetings
		}

		meetings = append(meetings, meeting)
	}

	if len(meetings) == 0 {
		return nil, ErrNotFoundMeetings
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
