package company

import (
	"encoding/json"
	"log"
)

type RolesAdapter interface {
	Roles
}

type Role struct {
	RoleID    int64
	CompanyID int64
	Name      string
	Type      string
	CreatedAt string
	UpdatedAt string
}

func (r *Role) ToString() string {
	b, err := json.Marshal(r)
	if err != nil {
		log.Println("err trying to parse role", err)
		return ""
	}

	return string(b)
}
