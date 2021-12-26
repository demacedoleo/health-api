package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"net/http"
	"strings"
	"time"
)

const (
	GenderFemale = "FEMENINO"
	Male         = "MALE"
	Female       = "FEMALE"
)

var (
	collector = colly.NewCollector()
)

type ScrapData struct {
	CompanyID   int    `json:"company_id,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	Document    string `json:"document"`
}

type UserAFJP struct {
	ScrapperID     int    `json:"scrapper_id,omitempty"`
	CompanyID      int    `json:"company_id,omitempty"`
	Document       string `json:"document,omitempty"`
	Member         string `json:"member"`
	MemberName     string `json:"member_name"`
	MemberLastName string `json:"member_last_name"`
	MemberNumber   string `json:"member_number"`
	GradeParent    string `json:"grade_parent"`
	CreatedAt      string `json:"created_at"`
	RemovedAt      string `json:"removed_at"`

	// Full Data
	MemberType        string   `json:"member_type"`
	Birthday          string   `json:"birthday"`
	Nationality       string   `json:"nationality"`
	Country           string   `json:"country"`
	Ugl               string   `json:"ugl"`
	Gender            string   `json:"gender"`
	MaritalStatus     string   `json:"marital_status"`
	MemberExpiration  string   `json:"member_expiration"`
	MemberSince       string   `json:"member_since"`
	MemberStatus      bool     `json:"member_status"`
	OperativeUnit     string   `json:"operative_unit"`
	AnotherHealthCare string   `json:"another_health_care"`
	HeadDoctor        string   `json:"head_doctor"`
	HeadDoctorNetwork string   `json:"head_doctor_network"`
	IsCompanyProvider bool     `json:"is_company_provider"`
	LastUpdated       string   `json:"last_updated"`
	Modalities        []string `json:"modalities"`
}

// I.N.S.S.J.P
func GetMemberDataAFJP(data ScrapData) (UserAFJP, error) {
	var member UserAFJP

	collector.AllowURLRevisit = true

	basicDataSelector := "p.whitetxt"
	fullDataSelector := "tbody:nth-child(1) > tr > td:nth-child(2) > p"
	providersSelector := "table:nth-child(9) > tbody > tr > td > p"
	notFoundSelector := "table:nth-child(4) > tbody > tr:nth-child(1) > td > p"

	var rowData []string

	code := 0

	collector.OnHTML(notFoundSelector, func(row *colly.HTMLElement) {
		fmt.Println(row.Text)
		if strings.ToLower(strings.TrimSpace(row.Text)) == "0 registro/s encontrado/s" {
			code = 404
		}
	})

	// Rules Basic Data
	collector.OnHTML(basicDataSelector, func(row *colly.HTMLElement) {
		rowData = append(rowData, row.Text)
	})

	// Rules Full Data
	collector.OnHTML(fullDataSelector, func(row *colly.HTMLElement) {
		rowData = append(rowData, row.Text)
	})

	var isCompanyProvider bool
	// Rules Providers Data
	modalities := make([]string, 0)
	history := make([]string, 0)

	collector.OnHTML(providersSelector, func(row *colly.HTMLElement) {
		history = append(history, strings.TrimSpace(row.Text))

		if strings.ToLower(strings.TrimSpace(row.Text)) == strings.ToLower(data.CompanyName) {
			modalities = append(modalities, history[len(history)-3])
			isCompanyProvider = true
		}
	})

	endpoint := "https://prestadores.pami.org.ar/result.php?c=6-2-2"
	body := map[string]string{"tipoDocumento": "DNI", "nroDocumento": data.Document, "submit2": "Buscar",
		"g-recaptcha-response": "03AGdBq25R1C8qYi5ecbpcDqGAOWelko53d5BhuOrjpL7jiXIpBwmfdtpM_0g0dcWPiAgRtq4DfQAQ6WXLSidT0d0dtQUWygdoYkeMPOI5scLgHxQ1KsHk9NLk2MV76jAzBei5dbxUDDrR3mRaVkRMMO2xOVXXit1yWlGI10CCtmfvdhFxDmzfGU32yHCC5AcU4GzUIhF6ScFYaXK0SAQzPkgn2rrKI6a4C6pBMkVN1jTtOFr7v1_BLDDufp-MuPfgCvbLwRx2HQmQUFAZ7mYEOKtpBe2fgMFzI2FeyR0KO3oMHJKlrYyELKrqTuZPkgY3JUbzipeSd1RUTIGh6W1r2JCzaoAQCQ6ehdGxHk5AdRM8fRacwV46HsEFMA0E8aI5J7uoP6xWGTFMTY3Yxv3s6A7yapnTn6V0xMTcyOLzREMFXfi4C5qKUIyAeNwH-y1Kui8Kzp9RwAQ1j5O2a1A90a7oW7bKAeTEEg",
	}

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	// Request Web Data
	if err := collector.Post(endpoint, body); err != nil {
		return member, errors.New("cannot retrieve data from inssjp")
	}

	if code == http.StatusNotFound {
		return member, errors.New("not found registry in inssjp")
	}

	if len(rowData) < 5 {
		return member, errors.New("cannot retrieve data from inssjp out of range")
	}

	memberNumber := strings.TrimSpace(rowData[1])
	parent := strings.TrimSpace(rowData[2])
	endpoint = fmt.Sprintf("https://prestadores.pami.org.ar/result.php?c=6-2-1-1&beneficio=%s&parent=%s&vm=2", memberNumber, parent)

	if err := collector.Visit(endpoint); err != nil {
		return member, errors.New("cannot retrieve full data from inssjp")
	}

	b := strings.Split(strings.TrimSpace(rowData[10]), "/")
	member.Birthday = fmt.Sprintf("%s-%s-%s", b[2], b[1], b[0])

	b = strings.Split(strings.TrimSpace(rowData[18]), "/")
	member.MemberSince = fmt.Sprintf("%s-%s-%s", b[2], b[1], b[0])

	if disaffiliated := strings.TrimSpace(rowData[19]); len(disaffiliated) == 0 {
		member.MemberStatus = true
	}

	member.CompanyID = data.CompanyID
	member.Document = data.Document
	member.Member = strings.TrimSpace(rowData[0])
	member.MemberNumber = memberNumber + "/" + parent
	member.GradeParent = strings.TrimSpace(rowData[2])
	member.CreatedAt = strings.TrimSpace(rowData[3])
	member.RemovedAt = strings.TrimSpace(rowData[4])
	member.MemberType = strings.TrimSpace(rowData[8])
	member.Nationality = strings.TrimSpace(rowData[11])
	member.Country = strings.TrimSpace(rowData[12])
	member.Ugl = strings.TrimSpace(rowData[13])

	if gender := strings.ToUpper(strings.TrimSpace(rowData[14])); gender == GenderFemale {
		member.Gender = Female
	} else {
		member.Gender = Male
	}


	member.MaritalStatus = strings.TrimSpace(rowData[15])
	member.MemberExpiration = strings.TrimSpace(rowData[16])
	member.OperativeUnit = strings.TrimSpace(rowData[17])
	member.OperativeUnit = strings.TrimSpace(rowData[17])
	member.AnotherHealthCare = strings.TrimSpace(rowData[20])
	member.HeadDoctor = strings.TrimSpace(rowData[21])
	member.HeadDoctorNetwork = strings.TrimSpace(rowData[22])
	member.IsCompanyProvider = isCompanyProvider
	member.LastUpdated = time.Now().Format(time.RFC3339)

	names := strings.Split(member.Member, " ")
	member.MemberName = strings.Join(names[1:], " ")
	member.MemberLastName = names[0]
	member.Modalities = modalities

	d, _ := json.Marshal(member)
	fmt.Println(string(d))

	return member, nil
}
