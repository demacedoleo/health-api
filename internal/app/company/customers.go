package company

import (
	"encoding/json"
	"log"
	"time"
)

type CustomersAdapter interface {
	Customers
}

type Customer struct {
	Person Person
}

type CustomerHealth struct {
	Person Person
	Health Health
}

type Health struct {
	ID             int64
	Document       string
	CompanyID      int64
	HealthName     string
	MemberID       string
	MemberStatus   string
	MemberInitDate string
	MemberEndDate  string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (ch *CustomerHealth) ToString() string {
	b, err := json.Marshal(ch)
	if err != nil {
		log.Println("err trying to parse customer health", err)
		return ""
	}

	return string(b)
}
