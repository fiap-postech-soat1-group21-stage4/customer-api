Feature: Customer API

  Scenario: Creating a new customer
    Given the following customer details
      | Name | CPF         | Email            |
      | João | 12312312312 | joao@email.com   |
    When a request is made to create the customer
    Then the response should have status code 201
    And the response body should match the expected customer details

  Scenario: Retrieving a customer
    Given a customer with CPF "12312312312" exists
    When a request is made to retrieve the customer
    Then the response should have status code 200
    And the response body should match the expected customer details
