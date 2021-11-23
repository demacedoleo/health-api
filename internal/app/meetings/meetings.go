package meetings

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type Adapter interface {
	Scheduler
}

type Meeting struct {
	ID                int64
	EventID           int64
	CompanyID         int64
	Subject           string
	MeetStatus        string
	StartTime         time.Time
	EndTime           time.Time
	AttendantDocument string
	AttendantPhone    string
	ResourceID        int64
	Modality          string
	CreatedAt         string
	UpdatedAt         string
}

func (m *Meeting) ToString() string {
	b, err := json.Marshal(m)
	if err != nil {
		log.Println("err trying to parse meeting", err)
		return ""
	}

	return strings.Replace(string(b), "00Z", "00", 2)
}

func CalculateWeek() (string, string) {
	var monday time.Time
	var sunday time.Time

	t := time.Now()
	weekday := t.Weekday()

	zero := 0
	if weekday == time.Sunday {
		monday = t.AddDate(zero, zero, -6)
		sunday = t.AddDate(zero, zero, 0)
		return FormatDate(monday), FormatDate(sunday)
	}

	if weekday == time.Saturday {
		monday = t.AddDate(zero, zero, -5)
		sunday = t.AddDate(zero, zero, 1)
		return FormatDate(monday), FormatDate(sunday)
	}

	if weekday == time.Friday {
		monday = t.AddDate(zero, zero, -4)
		sunday = t.AddDate(zero, zero, 2)
		return FormatDate(monday), FormatDate(sunday)
	}

	if weekday == time.Thursday {
		monday = t.AddDate(zero, zero, -3)
		sunday = t.AddDate(zero, zero, 3)
		return FormatDate(monday), FormatDate(sunday)
	}

	if weekday == time.Wednesday {
		monday = t.AddDate(zero, zero, -2)
		sunday = t.AddDate(zero, zero, 4)
		return FormatDate(monday), FormatDate(sunday)
	}

	if weekday == time.Tuesday {
		monday = t.AddDate(zero, zero, -1)
		sunday = t.AddDate(zero, zero, 5)
		return FormatDate(monday), FormatDate(sunday)
	}

	monday = t.AddDate(zero, zero, zero)
	sunday = t.AddDate(zero, zero, 6)
	return FormatDate(monday), FormatDate(sunday)
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}
