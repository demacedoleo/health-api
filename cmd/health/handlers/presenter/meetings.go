package presenter

import (
	"github.com/demacedoleo/health-api/cmd/health/handlers/entities"
	"github.com/demacedoleo/health-api/internal/app/meetings"
)

func Meetings(meetings []meetings.Meeting) []entities.Meeting {
	data := make([]entities.Meeting, len(meetings))
	for i, meeting := range meetings {
		data[i] = entities.Meeting{
			ID:                meeting.ID,
			EventID:           meeting.EventID,
			CompanyID:         meeting.CompanyID,
			Subject:           meeting.Subject,
			MeetStatus:        meeting.MeetStatus,
			StartTime:         meeting.StartTime,
			EndTime:           meeting.EndTime,
			Modality:          meeting.Modality,
			AttendantDocument: meeting.AttendantDocument,
			AttendantPhone:    meeting.AttendantPhone,
			ResourceID:        meeting.ResourceID,
			CreatedAt:         meeting.CreatedAt,
			UpdatedAt:         meeting.UpdatedAt,
		}
	}
	return data
}
