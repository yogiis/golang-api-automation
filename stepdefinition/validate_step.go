package stepdefinition

import (
	"encoding/json"
	"fmt"

	"github.com/oliveagle/jsonpath"
	"github.com/xeipuuv/gojsonschema"

	"github.com/yogiis/golang-api-automation/helper"
)

var json_data interface{}

// function code for validate response status code
func (t *Entity) ValidateStatusCode(expected string) error {
	fmt.Printf("Respon Body : %v ", string(t.ResponseBody))

	t.cases.AssertEqual(expected, t.ResponseData.StatusCode, helper.ErrorHandleEqual(expected, t.ResponseData.StatusCode))
	return nil
}

// function code for validate response body by value
func (t *Entity) ValidateResponseBody(path, expected string) error {
	jsonpath, err := jsonpath.Compile(path)
	helper.LogPanicln(err)

	json.Unmarshal(t.ResponseBody, &json_data)
	t.assertEqualByValue(jsonpath, expected)

	return nil
}

// function code for handle if an error occurs
func (t *Entity) assertEqualByValue(jsonPath *jsonpath.Compiled, expected string) {
	actual, err := jsonPath.Lookup(json_data)
	helper.LogPanicln(err)
	t.cases.AssertEqual(expected, actual, helper.ErrorHandleEqual(expected, actual))
}

// function code for validate response body by json schema
func (t *Entity) ValidateJSONSchema(jsonSchemaPath string) error {
	schemaLoader := gojsonschema.NewReferenceLoader(jsonSchemaPath)
	jsonLoader := gojsonschema.NewBytesLoader(t.ResponseBody)

	// Use Validator to check whether all JSON objects in the response match the JSON schema.
	validator, err := gojsonschema.NewSchemaLoader().Compile(schemaLoader)
	helper.LogPanicln(err)

	result, err := validator.Validate(jsonLoader)
	helper.LogPanicln(err)

	t.cases.AssertTrue(result.Valid(), helper.ErrorHandleBool(result.Valid()))

	return nil
}
