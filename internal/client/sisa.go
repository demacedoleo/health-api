package client

import (
	"context"
	"encoding/xml"
	"errors"
	oops "github.com/demacedoleo/health-api/pkg/errors"
	"github.com/demacedoleo/health-api/pkg/rest"
	"net/http"
	"time"
)

const (
	host     = "https://sisa.msal.gov.ar"
	endpoint = host + "/sisa/services/rest/profesional/obtener?usuario=lfpereira&clave=EQEOJ6707P&nrodoc="
	user     = ""
	pass     = ""
)

var (
	client = http.Client{
		Timeout: time.Duration(2) * time.Second,
	}
)

type Professional struct {
	XMLName           xml.Name `xml:"Professional"`
	Text              string   `xml:",chardata"`
	Resultado         string   `xml:"resultado"`
	Apellido          string   `xml:"apellido"`
	Codigo            string   `xml:"codigo"`
	FechaModificacion string   `xml:"fechaModificacion"`
	FechaRegistro     string   `xml:"fechaRegistro"`
	Matriculas        struct {
		Text      string `xml:",chardata"`
		Matricula struct {
			Text           string `xml:",chardata"`
			Especialidades struct {
				Text         string `xml:",chardata"`
				Especialidad []struct {
					Text               string `xml:",chardata"`
					Especialidad       string `xml:"especialidad"`
					FechaCertificacion string `xml:"fechaCertificacion"`
					Ministerio         string `xml:"ministerio"`
					TipoCertificacion  string `xml:"tipoCertificacion"`
				} `xml:"especialidad"`
			} `xml:"especialidades"`
			Estado            string `xml:"estado"`
			FechaMatricula    string `xml:"fechaMatricula"`
			FechaModificacion string `xml:"fechaModificacion"`
			FechaRegistro     string `xml:"fechaRegistro"`
			Jurisdiccion      string `xml:"jurisdiccion"`
			Matricula         string `xml:"matricula"`
			Profesion         string `xml:"profesion"`
			Provincia         string `xml:"provincia"`
		} `xml:"matricula"`
	} `xml:"matriculas"`
	Nombre          string `xml:"nombre"`
	NumeroDocumento string `xml:"numeroDocumento"`
	TipoDocumento   string `xml:"tipoDocumento"`
}

func GetProfessional(ctx context.Context, document string) (map[string]interface{}, error) {
	h := rest.Headers{}
	h.Add("Content-Type", "application/xml")

	response, err := rest.NewRequest(rest.WithClient(client), rest.WithHeaders(h)).Get(endpoint + document)
	if err != nil {
		return nil, errors.New("error getting professional")
	}

	var p map[string]interface{}
	if err = xml.Unmarshal(response.Body, &p); err != nil {
		return nil, oops.Errorf(oops.E5xxINTERNAL, err.Error())
	}

	return p, nil
}
