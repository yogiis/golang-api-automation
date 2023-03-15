package stepdefinition

import (
	"bytes"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/cucumber/godog"
	"github.com/joho/godotenv"
	"github.com/yogiis/golang-api-automation/helper"
)

func (e *Entity) GivenEndpoint(host string, endpoint string) error {
	env := godotenv.Load()
	helper.LogPanicln(env)

	host = os.Getenv(host)
	e.UrlEndpoint = host + endpoint

	return nil
}

func (e *Entity) SendPOSTEndpointWithBodyFormUrlEncoded(table *godog.Table) error {
	formData := url.Values{}

	for _, row := range table.Rows {
		for i := 0; i < len(row.Cells); i += 2 {
			formData.Set(row.Cells[i].Value, row.Cells[i+1].Value)
		}
	}

	hitEndpoint, err := http.NewRequest(http.MethodPost, e.UrlEndpoint, strings.NewReader(formData.Encode()))
	helper.LogPanicln(err)
	hitEndpoint.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	AddAPIKeyHeaderIfHas(hitEndpoint)

	e.ResponseData, err = SendHTTPRequest(hitEndpoint)
	helper.LogPanicln(err)

	e.ResponseBody, err = ReaderResponseBody(e.ResponseData)
	helper.LogPanicln(err)

	defer CloseResponseBody(e.ResponseData)

	return nil
}

func (e *Entity) SendGETEndpointWithoutBody() error {
	hitEndpoint, err := http.NewRequest(http.MethodGet, e.UrlEndpoint, nil)
	helper.LogPanicln(err)
	hitEndpoint.Header.Add("Content-Type", "application/json")
	AddAPIKeyHeaderIfHas(hitEndpoint)

	e.ResponseData, err = SendHTTPRequest(hitEndpoint)
	helper.LogPanicln(err)

	e.ResponseBody, err = ReaderResponseBody(e.ResponseData)
	helper.LogPanicln(err)

	defer CloseResponseBody(e.ResponseData)

	return nil
}

func (e *Entity) SendPOSTEndpointWithBodyJSON(requestBody *godog.DocString) error {
	hitEndpoint, err := http.NewRequest(http.MethodPost, e.UrlEndpoint, bytes.NewBuffer([]byte(requestBody.Content)))
	helper.LogPanicln(err)
	hitEndpoint.Header.Add("Content-Type", "application/json")
	AddAPIKeyHeaderIfHas(hitEndpoint)

	e.ResponseData, err = SendHTTPRequest(hitEndpoint)
	helper.LogPanicln(err)

	e.ResponseBody, err = ReaderResponseBody(e.ResponseData)
	helper.LogPanicln(err)

	defer CloseResponseBody(e.ResponseData)

	return nil
}

func (e *Entity) SendGETEndpointWithParams(params string) error {
	e.UrlEndpoint = e.UrlEndpoint + params
	hitEndpoint, err := http.NewRequest(http.MethodGet, e.UrlEndpoint, nil)
	helper.LogPanicln(err)
	hitEndpoint.Header.Add("Content-Type", "application/json")
	AddAPIKeyHeaderIfHas(hitEndpoint)

	e.ResponseData, err = SendHTTPRequest(hitEndpoint)
	helper.LogPanicln(err)

	e.ResponseBody, err = ReaderResponseBody(e.ResponseData)
	helper.LogPanicln(err)

	defer CloseResponseBody(e.ResponseData)

	return nil
}

func (e *Entity) SendPUTEndpointWithBodyJSON(params string, requestBody *godog.DocString) error {
	e.UrlEndpoint = e.UrlEndpoint + params
	hitEndpoint, err := http.NewRequest(http.MethodPut, e.UrlEndpoint, bytes.NewBuffer([]byte(requestBody.Content)))
	helper.LogPanicln(err)
	hitEndpoint.Header.Add("Content-Type", "application/json")
	AddAPIKeyHeaderIfHas(hitEndpoint)

	e.ResponseData, err = SendHTTPRequest(hitEndpoint)
	helper.LogPanicln(err)

	e.ResponseBody, err = ReaderResponseBody(e.ResponseData)
	helper.LogPanicln(err)

	defer CloseResponseBody(e.ResponseData)

	return nil
}

func (e *Entity) SendDELETEEndpointWithParams(params string) error {
	e.UrlEndpoint = e.UrlEndpoint + params
	hitEndpoint, err := http.NewRequest(http.MethodDelete, e.UrlEndpoint, nil)
	helper.LogPanicln(err)
	hitEndpoint.Header.Add("Content-Type", "application/json")
	AddAPIKeyHeaderIfHas(hitEndpoint)

	e.ResponseData, err = SendHTTPRequest(hitEndpoint)
	helper.LogPanicln(err)

	e.ResponseBody, err = ReaderResponseBody(e.ResponseData)
	helper.LogPanicln(err)

	defer CloseResponseBody(e.ResponseData)

	return nil
}
