package rest

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	oops "github.com/demacedoleo/health-api/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Actions interface {
	Get(uri string) (*Response, error)
	Post(uri string) (*Response, error)
	Put(endpoint string) (*Response, error)
}

type Body []byte

func (b Body) ToString() string {
	return string(b)
}

type Response struct {
	StatusCode int
	Body       Body
}

type Request struct {
	c       http.Client
	body    io.Reader
	headers Headers
}

func (r *Request) Get(endpoint string) (*Response, error) {
	return execute(http.MethodGet, r.c, endpoint, r.body, r.headers)
}

func (r *Request) Post(endpoint string) (*Response, error) {
	return execute(http.MethodPost, r.c, endpoint, r.body, r.headers)
}

func (r *Request) Put(endpoint string) (*Response, error) {
	return execute(http.MethodPut, r.c, endpoint, r.body, r.headers)
}

func NewRequest(options ...func(request *Request)) *Request {
	defaultHeader := make(map[string]string)
	defaultHeader["Content-Type"] = "application/json"

	r := &Request{
		headers: Headers{headers: defaultHeader},
	}
	for _, option := range options {
		option(r)
	}
	return r
}

func WithClient(c http.Client) func(*Request) {
	return func(request *Request) {
		request.c = c
	}
}

func WithBody(body interface{}) func(*Request) {
	if b, ok := body.(string); ok {
		return func(request *Request) {
			request.body = strings.NewReader(b)
		}
	}

	if b, ok := body.([]byte); ok {
		return func(request *Request) {
			request.body = bytes.NewReader(b)
		}
	}

	return func(request *Request) {
		b, _ := json.Marshal(body)
		request.body = bytes.NewReader(b)
	}
}

func WithHeaders(h Headers) func(*Request) {
	return func(request *Request) {
		request.headers = h
	}
}

type Headers struct {
	headers map[string]string
}

func (h *Headers) Add(k, v string) *Headers {
	if h.headers == nil {
		h.headers = make(map[string]string)
		h.headers["Content-Type"] = "application/json"
	}
	h.headers[k] = v
	return h
}

func BindResponse(response *Response, err error, in interface{}) error {
	if err != nil {
		return err
	}

	if in == nil {
		return nil
	}

	if err = json.Unmarshal(response.Body, &in); err != nil {
		return oops.Errorf(oops.E5xxINTERNAL, err.Error())
	}

	return nil
}

func BindXmlResponse(response *Response, err error, in interface{}) error {
	if err != nil {
		return err
	}

	if in == nil {
		return nil
	}

	if err = xml.Unmarshal(response.Body, &in); err != nil {
		return oops.Errorf(oops.E5xxINTERNAL, err.Error())
	}

	return nil
}

func toString(body io.ReadCloser) string {
	r, err := ioutil.ReadAll(body)
	if err != nil {
		return ""
	}
	return string(r)
}

func execute(method string, client http.Client, endpoint string, payload io.Reader, h Headers) (*Response, error) {
	req, err := http.NewRequest(method, endpoint, payload)
	if err != nil {
		return nil, oops.Errorf(oops.E5xxINTERNAL, "error building request")
	}

	for k, v := range h.headers {
		req.Header.Add(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, oops.Errorf(oops.E5xxINTERNAL, fmt.Sprintf("[execute] [endpoint: %s] [error: %s]", endpoint, err.Error()))
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return makeResponse(response)
	case http.StatusBadRequest:
		return nil, oops.Errorf(oops.E4xxCLIENTSIDE, toString(response.Body))
	case http.StatusUnauthorized:
		return nil, oops.Errorf(oops.E4xxUNAUTHORIZED, toString(response.Body))
	case http.StatusNotFound:
		return nil, oops.Errorf(oops.E4xxNOTFOUND, toString(response.Body))
	case http.StatusUnprocessableEntity:
		return nil, oops.Errorf(oops.E4xxUNPROCESSABLE, toString(response.Body))
	default:
		return nil, oops.Errorf(oops.E5xxINTERNAL, toString(response.Body))
	}
}

func makeResponse(response *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, oops.Errorf(oops.E5xxINTERNAL, "error reading response")
	}

	fmt.Println(string(body))

	return &Response{
		StatusCode: response.StatusCode,
		Body:       body,
	}, nil
}
