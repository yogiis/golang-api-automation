@ping-api
Feature: Hit ping api

  @get
  Scenario: endpoint ping will response http 200
    Given I set host "ENV_HOST_API" with endpoint "/users"
    When I send request using method POST with type body json:
      """
      {
        "name": "John Doe",
        "email": "john.doe@example.com",
        "password": "p@ssw0rd",
        "address": {
          "street": "123 Main St",
          "city": "Anytown",
          "state": "CA",
          "zip": "12345"
        }
      }
      """
    Then assert status code is "201"


  @get
  Scenario: endpoint ping will response http 200
    Given I set host "ENV_HOST_API" with endpoint "/users/3"
    When I send request using method GET
    Then assert status code is "200"
    And match with json schema "file://schema/test.json"
    And assert value "$.id" equal is "1"

  @get
  Scenario: endpoint ping will response http 200
    Given I set host "ENV_HOST_API" with endpoint "/users"
    When I send request using method GET with params "/3"
    Then assert status code is "200"
    And match with json schema "file://schema/test.json"


  @get
  Scenario: endpoint ping will response http 200
    Given I set host "ENV_HOST_API" with endpoint "/api/v2/user/login"
    When I send request using method POST with type body x-www-form-urlencoded:
      | via      | email                 |
      | platform | web                   |
      | version  | 2                     |
      | email    | kangyogi047@gmail.com |
      | password | R4h4si4123            |
    Then assert status code is "201"

