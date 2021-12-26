package company

import (
	"context"

	"github.com/demacedoleo/health-api/internal/platform/mysql"
)

type Service interface {
	Institution
	Modalities
	Roles
	Staffs
	Customers
}

type service struct {
	Institution Institution
	Modalities  Modalities
	Roles       Roles
	Staffs      Staffs
	Customers   Customers
}

type Institution interface {
	GetCompany(ctx context.Context, id int64) (*Company, error)
	CreateCompany(ctx context.Context, c Company) error
}

type Roles interface {
	GetRoles(ctx context.Context, companyID int64) ([]Role, error)
	CreateRole(ctx context.Context, role Role) error
}

type Modalities interface {
	GetModalities(ctx context.Context, companyID int64) ([]Modality, error)
	CreateModality(ctx context.Context, modality Modality) error
}

type Staffs interface {
	FindStaffs(ctx context.Context, companyID int64) ([]Staff, error)
	CreateStaff(ctx context.Context, staff Staff) error
}

type Customers interface {
	FindCustomers(ctx context.Context, companyID int64) ([]Customer, error)
	CreateCustomers(ctx context.Context, staff CustomerHealth) error
}

func (s *service) GetCompany(ctx context.Context, id int64) (*Company, error) {
	return s.Institution.GetCompany(ctx, id)
}

func (s *service) CreateCompany(ctx context.Context, company Company) error {
	return s.Institution.CreateCompany(ctx, company)
}

func (s *service) GetRoles(ctx context.Context, companyID int64) ([]Role, error) {
	return s.Roles.GetRoles(ctx, companyID)
}

func (s *service) CreateRole(ctx context.Context, role Role) error {
	return s.Roles.CreateRole(ctx, role)
}

func (s *service) GetModalities(ctx context.Context, companyID int64) ([]Modality, error) {
	return s.Modalities.GetModalities(ctx, companyID)
}

func (s *service) CreateModality(ctx context.Context, modality Modality) error {
	return s.Modalities.CreateModality(ctx, modality)
}

func (s *service) FindStaffs(ctx context.Context, companyID int64) ([]Staff, error) {
	return s.Staffs.FindStaffs(ctx, companyID)
}

func (s *service) CreateStaff(ctx context.Context, staff Staff) error {
	return s.Staffs.CreateStaff(ctx, staff)
}

func (s *service) FindCustomers(ctx context.Context, companyID int64) ([]Customer, error) {
	return s.Customers.FindCustomers(ctx, companyID)
}

func (s *service) CreateCustomers(ctx context.Context, ch CustomerHealth) error {
	return s.Customers.CreateCustomers(ctx, ch)
}

func NewCompanyService(repository mysql.Repository) *service {
	return &service{
		Institution: NewCompanyAdapter(repository),
		Roles:       NewRolesAdapter(repository),
		Modalities:  NewModalitiesAdapter(repository),
		Staffs:      NewStaffsAdapter(repository),
		Customers:   NewCustomerAdapter(repository),
	}
}
