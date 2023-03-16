package main

import (
	"testing"

	"github.com/cucumber/godog"
	suiteTest "github.com/yogiis/golang-api-automation/stepdefinition"
)

func TestFeatures(t *testing.T) {

	suite := godog.TestSuite{
		ScenarioInitializer: suiteTest.InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty,cucumber:./report/cucumber.json",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
			Tags:     "@example",
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
