var reporter = require("cucumber-html-reporter");

var options = {
  theme: "bootstrap",
  jsonFile: "./report/cucumber.json",
  output: "./report/report.html",
  reportSuiteAsScenarios: true,
  scenarioTimestamp: true,
  launchReport: true,
  metadata: {
    "App Version": "0.0.2",
    "Test Environment": "STAGING",
    Browser: "Chrome",
    Platform: "Macbook Pro M1",
    Parallel: "Scenario",
    Executed: "Remote",
  },
};

reporter.generate(options);
