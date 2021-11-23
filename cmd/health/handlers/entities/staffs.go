package entities

type Staff struct {
	ID        int64     `json:"id,omitempty"`
	CompanyID int64     `json:"company_id,omitempty"`
	Registers []string  `json:"registers,omitempty"`
	Job       Charge    `json:"job"`
	Person    Person    `json:"person"`
	Contact   Contact   `json:"contact,omitempty"`
	Address   Address   `json:"address,omitempty"`
	Schedule  []WorkDay `json:"schedule,omitempty"`
	Status    bool      `json:"status,omitempty"`
}

type StaffFlatten struct {
	ID        int64  `json:"id,omitempty"`
	CompanyID int64  `json:"company_id,omitempty"`
	Doc       string `json:"document,omitempty"`
	Person
	Charge
	Status bool   `json:"status,omitempty"`
	Color  string `json:"color,omitempty"`
}

type Register struct {
	CompanyID int64  `json:"company_id,omitempty"`
	Document  string `json:"document,omitempty"`
	Register  string `json:"register,omitempty"`
}

type Charge struct {
	Document    string `json:"document,omitempty"`
	CompanyID   int64  `json:"company_id,omitempty"`
	ChargeType  string `json:"charge_type,omitempty"`
	JobPosition string `json:"job_position,omitempty"`
	Color       string `json:"color"`
	Status      bool   `json:"status,omitempty"`
}

type WorkDay struct {
	Document  string
	CompanyID int64  `json:"company_id,omitempty"`
	WeekDay   string `json:"weekday,omitempty"`
	Start     string `json:"start,omitempty"`
	End       string `json:"end,omitempty"`
}
