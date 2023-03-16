package stepdefinition

import (
	"net/http"

	"github.com/yogiis/golang-api-automation/helper"
)

type Entity struct {
	UrlEndpoint  string
	ResponseData *http.Response
	Cases        helper.Case
	ResponseBody []byte
}
