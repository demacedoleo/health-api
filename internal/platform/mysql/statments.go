package mysql

type (
	inserts struct {
		Company        string
		Role           string
		Modality       string
		HealthProvider string
		Meeting        string
		Staff          string
	}

	selects struct {
		Company         string
		Roles           string
		States          string
		Cities          string
		Modalities      string
		HealthProviders string
		Meetings        string
		Staffs          string
	}

	updates struct{}

	deletes struct{}

	statements struct {
		Inserts inserts
		Updates updates
		Deletes deletes
		Selects selects
	}
)

var (
	Statements = statements{
		Inserts: inserts{
			Company:        "CALL InsertCompany(?);",
			Role:           "CALL InsertCompanyRole(?);",
			Modality:       "CALL InsertHealthModality(?);",
			HealthProvider: "CALL InsertHealthProvider(?);",
			Meeting:        "CALL InsertMeeting(?);",
			Staff:          "CALL InsertStaff(?);",
		},
		Selects: selects{
			Company:         "select * from company where company_id = ?;",
			Roles:           "select * from company_roles where company_id = ?;",
			States:          "select * from country_states;",
			Cities:          "select * from country_cities where state_id = ?;",
			Modalities:      "select * from company_modalities where company_id = ?;",
			HealthProviders: "select * from company_health_providers where company_id = ?;",
			Meetings:        "select * from health.scheduler_meetings where company_id = ? and start_time >= ? and end_time <= ?;",
			Staffs:          "select cs.id, cs.company_id, cs.document, cs.charge_type, cs.job_position, cs.status, cs.color, p.document, p.person_name, p.last_name, p.birthday, p.gender from company_staffs as cs join person as p on cs.document = p.document where cs.company_id = ?;",
		},
	}
)
