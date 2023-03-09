package stepdefinition

import "github.com/cucumber/godog"

func InitializeScenario(sc *godog.ScenarioContext) {
	step := Entity{}

	sc.Step(`I set host "([^"]*)" with endpoint "([^"]*)"$`, step.GivenEndpoint)
	sc.Step(`I send request using method GET with params "([^"]*)"`, step.SendGETEndpointWithParams)
	sc.Step(`I send request using method GET`, step.SendGETEndpointWithoutBody)
	sc.Step(`I send request using method POST with type body json:`, step.SendPOSTEndpointWithBodyJSON)
	sc.Step(`I send request using method POST with type body x-www-form-urlencoded:`, step.SendPOSTEndpointWithBodyFormUrlEncoded)
	sc.Step(`assert status code is "([^"]*)"`, step.ValidateStatusCode)
	sc.Step(`assert value "([^"]*)" equal is "([^"]*)"`, step.ValidateResponseBody)
	sc.Step(`match with json schema "([^"]*)"`, step.ValidateJSONSchema)
}
