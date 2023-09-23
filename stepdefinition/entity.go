package stepdefinition

import (
	"net/http"

	"github.com/yogiis/golang-api-automation/helper"
	"gorm.io/gorm"
)

type Entity struct {
	UrlEndpoint  string
	ResponseData *http.Response
	Cases        helper.Case
	ResponseBody []byte
	db           *gorm.DB
}
