package stepdefinition

import (
	"net/http"
	"testing"

	"github.com/yogiis/golang-api-automation/helper"
)

type Entity struct {
	Token        string `json:"token"`
	UrlEndpoint  string
	ResponseData *http.Response
	Testing      testing.T
	cases        helper.Case
	ResponseBody []byte
}
