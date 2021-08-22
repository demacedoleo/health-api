package company

import (
	"context"
	"errors"
	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

var (
	ErrStaffsScan     = errors.New("cannot iterate over rows")
	ErrNotFoundStaffs = errors.New("not found staffs")
)

type staffsAdapter struct {
	mysql.Repository
}

func (s *staffsAdapter) FindStaffs(ctx context.Context, companyID int64) ([]Staff, error) {
	result, err := s.Fetch(ctx, mysql.Statements.Selects.Staffs, companyID)
	if err != nil {
		return nil, err
	}
	var staffs []Staff
	for result.Next() {
		var staff Staff
		if err := result.
			Scan(&staff.ID, &staff.Job.CompanyID, &staff.Job.Document, &staff.Job.ChargeType,
				&staff.Job.JobPosition, &staff.Status, &staff.Person.Document, &staff.Person.Name,
				&staff.Person.LastName, &staff.Person.Birthday, &staff.Person.Gender,
			); err != nil {
			return nil, ErrStaffsScan
		}

		staffs = append(staffs, staff)
	}

	if len(staffs) == 0 {
		return nil, ErrNotFoundStaffs
	}

	return staffs, nil
}

func (s *staffsAdapter) CreateStaff(ctx context.Context, staff Staff) error {
	staff.Person.Type = TypeStaff.ToString()
	staff.Person.CompanyID = staff.CompanyID

	staff.Contact.CompanyID = staff.CompanyID
	staff.Contact.Document = staff.Person.Document

	staff.Address.CompanyID = staff.CompanyID
	staff.Address.Document = staff.Person.Document

	staff.Job.CompanyID = staff.CompanyID
	staff.Job.Document = staff.Person.Document

	_, err := s.Save(ctx, mysql.Statements.Inserts.Staff, staff.ToString())
	return err
}

func NewStaffsAdapter(repository mysql.Repository) *staffsAdapter {
	return &staffsAdapter{repository}
}
