Feature: CRUD Restful API

  Scenario: Example Update by users id
    Given i set host "ENV_HOST_API" with endpoint "/api/users"
    When i send request using method PUT with params "/1" and type body json:
      """
      {
        "name": "Yogi"
      }
      """
    Then assert status code is "200"
    And assert value "$.status" equal is "OK"

  @example
  Scenario: Create new users
    Given i set connection to database
    And i set host "ENV_HOST_API" with endpoint "/api/users"
    When i send request using method POST with type body x-www-form-urlencoded:
      | name | Yogi |
    Then assert status code is "200"
    And assert value "$.status" equal is "OK"
    And match with json schema "file://schema/test.json"
    And assert value db with query "select name from users where id = 5" equal is "Yogi"

  Scenario: Get all users
    Given i set host "ENV_HOST_API" with endpoint "/api/users"
    When i send request using method GET
    Then assert status code is "200"
    And assert value "$.status" equal is "OK"

  @example
  Scenario: Get by user id
    Given i set connection to database
    And i set host "ENV_HOST_API" with endpoint "/api/users"
    When i send request using method GET with params "/1"
    Then assert status code is "200"
    And assert value "$.status" equal is "OK"
    And assert value db with query "select name from users where id = 5" equal is "Ari"

  Scenario: Delete by user id
    Given i set host "ENV_HOST_API" with endpoint "/api/users"
    When i send request using method DELETE with params "/13"
    Then assert status code is "200"
    And assert value "$.status" equal is "OK"
