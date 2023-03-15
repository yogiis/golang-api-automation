package stepdefinition

import (
	"net/http"
	"testing"

	"github.com/yogiis/golang-api-automation/helper"
)

type Entity struct {
	UrlEndpoint  string
	ResponseData *http.Response
	Testing      testing.T
	Cases        helper.Case
	ResponseBody []byte
}
