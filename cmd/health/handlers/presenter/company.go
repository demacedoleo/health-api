package presenter

import (
	"github.com/demacedoleo/health-api/cmd/health/handlers/entities"
	"github.com/demacedoleo/health-api/internal/app/company"
)

func Company(c *company.Company) *entities.Company {
	return &entities.Company{
		CompanyID:        c.CompanyID,
		CompanyName:      c.CompanyName,
		CompanyShortName: c.CompanyShortName,
		CompanyColor:     c.CompanyColor,
		CompanyRegister:  c.CompanyRegister,
	}
}

func Roles(r []company.Role) []entities.Role {
	data := make([]entities.Role, len(r))
	for i, role := range r {
		data[i] = entities.Role{
			RoleID: role.RoleID,
			Name:   role.Name,
			Type:   role.Type,
		}
	}
	return data
}

func Modalities(modalities []company.Modality) []entities.Modality {
	output := make([]entities.Modality, len(modalities))
	for i, modality := range modalities {
		output[i] = entities.Modality{
			ID:       modality.ID,
			Modality: modality.Modality,
		}
	}
	return output
}

func Staffs(staffs []company.Staff) []entities.StaffFlatten {
	output := make([]entities.StaffFlatten, len(staffs))
	for i, s := range staffs {
		output[i] = entities.StaffFlatten{
			ID:        s.ID,
			CompanyID: s.Job.CompanyID,
			Doc:       s.Person.Document,
			Person:    entities.Person(s.Person),
			Charge:    entities.Charge(s.Job),
			Status:    s.Status,
			Color:     s.Color,
		}
	}
	return output
}

func Customers(customers []company.Customer) []entities.Person {
	output := make([]entities.Person, len(customers))
	for i, c := range customers {
		output[i] = entities.Person{
			ID:        c.Person.ID,
			Document:  c.Person.Document,
			CompanyID: c.Person.CompanyID,
			Name:      c.Person.Name,
			LastName:  c.Person.LastName,
			Birthday:  c.Person.Birthday,
			Gender:    c.Person.Gender,
			Type:      c.Person.Type,
			CreatedAt: c.Person.CreatedAt,
			UpdatedAt: c.Person.UpdatedAt,
		}
	}
	return output
}
