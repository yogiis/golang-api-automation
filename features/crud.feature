Feature: CRUD Restful API

  @post
  Scenario: Create new users
    Given i set host "ENV_HOST_API" with endpoint "/api/users"
    When i send request using method POST with type body x-www-form-urlencoded:
      | name | Yogi |
    Then assert status code is "200"
    And assert value "$.status" equal is "OK"
    And match with json schema "file://schema/test.json"

  @put
  Scenario: Update by users id
    Given i set host "ENV_HOST_API" with endpoint "/api/users"
    When i send request using method PUT with params "/1" and type body json:
      """
      {
        "name": "Ari"
      }
      """
    Then assert status code is "200"
    And assert value "$.status" equal is "OK"
    And match with json schema "file://schema/test.json"

  @get
  Scenario: Get all users
    Given i set host "ENV_HOST_API" with endpoint "/api/users"
    When i send request using method GET
    Then assert status code is "200"
    And assert value "$.status" equal is "OK"

  @get
  Scenario: Get by user id
    Given i set host "ENV_HOST_API" with endpoint "/api/users"
    When i send request using method GET with params "/1"
    Then assert status code is "200"
    And assert value "$.status" equal is "OK"

  @delete
  Scenario: Delete by user id
    Given i set host "ENV_HOST_API" with endpoint "/api/users"
    When i send request using method DELETE with params "/10"
    Then assert status code is "200"
    And assert value "$.status" equal is "OK"
