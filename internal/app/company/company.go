package company

import (
	"encoding/json"
	"log"
)

type InstitutionAdapter interface {
	Institution
}

type Company struct {
	CompanyID        int64
	CompanyName      string
	CompanyShortName string
	CompanyColor     string
	CompanyRegister  string
	CreatedAt        string
	UpdatedAt        string
}

type Person struct {
	ID        int64
	Document  string
	CompanyID int64
	Parent    string
	Name      string
	LastName  string
	Birthday  string
	Gender    string
	Type      string
	UpdatedAt string
	CreatedAt string
}

type Contact struct {
	ID        int64
	Document  string
	CompanyID int64
	Phone     string
	Mail      string
}

type Address struct {
	ID               int64
	Document         string
	CompanyID        int64
	StreetName       string
	StreetNumber     string
	StreetComplement string
	ZipCode          string
	State            string
	City             string
}

type PersonType int

const (
	TypeStaff PersonType = iota
	TypeRelative
)

func (p PersonType) ToString() string {
	return []string{"STAFF", "RELATIVE"}[p]
}

func (a *Address) IsValid() bool {
	return a != nil && len(a.StreetName) > 0 && len(a.StreetNumber) > 0
}

func (c *Company) ToString() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Println("err trying to parse company", err)
		return ""
	}

	return string(b)
}
