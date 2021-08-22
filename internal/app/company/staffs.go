package company

import (
	"encoding/json"
	"log"
)

type StaffsAdapter interface {
	Staffs
}

type Staff struct {
	ID        int64
	CompanyID int64
	Registers []Register
	Job       Charge
	Person    Person
	Contact   Contact
	Address   Address
	Schedule  []WorkDay
	Status    bool
}

type Register struct {
	CompanyID int64
	Document  string
	Register  string
}

type Charge struct {
	Document    string
	CompanyID   int64
	ChargeType  string
	JobPosition string
	Status      bool
}

type WorkDay struct {
	Document  string
	CompanyID int64
	WeekDay   string
	Start     string
	End       string
}

func (s *Staff) ToString() string {
	staff := make(map[string]interface{})
	staff["CompanyID"] = s.CompanyID
	staff["Job"] = s.Job
	staff["Person"] = s.Person
	staff["Status"] = s.Status

	if len(s.Registers) > 0 {
		staff["Registers"] = s.Registers
	}

	if s.Contact != (Contact{}) {
		staff["Contact"] = s.Contact
	}

	if s.Address != (Address{}) {
		staff["Address"] = s.Address
	}

	if len(s.Schedule) > 0 {
		staff["Schedule"] = s.Schedule
	}

	b, err := json.Marshal(staff)
	if err != nil {
		log.Println("err trying to parse staff", err)
		return ""
	}

	return string(b)
}
