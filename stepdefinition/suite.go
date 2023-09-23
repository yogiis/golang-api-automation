package stepdefinition

import (
	"github.com/cucumber/godog"
)

func InitializeScenario(sc *godog.ScenarioContext) {
	step := &Entity{}

	sc.Step(`i set host "([^"]*)" with endpoint "([^"]*)"$`, step.GivenEndpoint)
	sc.Step(`i send request using method GET with params "([^"]*)"`, step.SendGETEndpointWithParams)
	sc.Step(`i send request using method GET`, step.SendGETEndpointWithoutBody)
	sc.Step(`i send request using method POST with type body json:`, step.SendPOSTEndpointWithBodyJSON)
	sc.Step(`i send request using method POST with type body x-www-form-urlencoded:`, step.SendPOSTEndpointWithBodyFormUrlEncoded)
	sc.Step(`i send request using method PUT with params "([^"]*)" and type body json:`, step.SendPUTEndpointWithBodyJSON)
	sc.Step(`i send request using method DELETE with params "([^"]*)"`, step.SendDELETEEndpointWithParams)
	sc.Step(`assert status code is "([^"]*)"`, step.ValidateStatusCode)
	sc.Step(`assert value "([^"]*)" equal is "([^"]*)"`, step.ValidateResponseBody)
	sc.Step(`match with json schema "([^"]*)"`, step.ValidateJSONSchema)
	sc.Step(`i set connection to database`, step.InitializeDB)
	sc.Step(`assert value db with query "([^"]*)" equal is "([^"]*)"`, step.ValidateValueDB)
}
