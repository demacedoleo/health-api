package entities

type Company struct {
	CompanyID        int64  `json:"company_id,omitempty"`
	CompanyName      string `json:"company_name"       binding:"required"`
	CompanyShortName string `json:"company_short_name" binding:"required"`
	CompanyColor     string `json:"company_color"      binding:"required"`
	CompanyRegister  string `json:"company_register"   binding:"required"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
}

type Role struct {
	RoleID    int64  `json:"id,omitempty"`
	CompanyID int64  `json:"company_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type Modality struct {
	ID        int64  `json:"id"`
	CompanyID int64  `json:"company_id,omitempty"`
	Modality  string `json:"modality"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type Person struct {
	ID        int64  `json:"id,omitempty"`
	Document  string `json:"document,omitempty"`
	CompanyID int64  `json:"company_id,omitempty"`
	Parent    string `json:"relative_grade,omitempty"`
	Name      string `json:"person_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Birthday  string `json:"birthday,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Type      string `json:"person_type,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type Contact struct {
	ID        int64  `json:"id,omitempty"`
	Document  string `json:"document,omitempty"`
	CompanyID int64  `json:"company_id,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Mail      string `json:"mail,omitempty"`
}

type Address struct {
	ID               int64  `json:"id,omitempty"`
	Document         string `json:"document,omitempty"`
	CompanyID        int64  `json:"company_id,omitempty"`
	StreetName       string `json:"street_name,omitempty"`
	StreetNumber     string `json:"street_number,omitempty"`
	StreetComplement string `json:"street_complement,omitempty"`
	ZipCode          string `json:"zip_code,omitempty"`
	State            string `json:"state,omitempty"`
	City             string `json:"city,omitempty"`
}
