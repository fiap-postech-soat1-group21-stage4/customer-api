package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/entity"
	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/port"
	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/port/mocks"
	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	defaultID  = uuid.MustParse("b4dacf92-7000-4523-9fab-166212acc92d")
	ctxDefault = context.Background()

	given = &entity.Customer{
		Name:  "João",
		Email: "joao@email.com",
		CPF:   "12345678912",
	}

	wanted = &entity.Customer{
		ID:    defaultID,
		Name:  "João",
		Email: "joao@email.com",
		CPF:   "12345678912",
	}
)

type customerTest struct {
	service port.CustomerUseCase
	repo    *mocks.CustomerRepository
	err     error
	result  *entity.Customer
}

func (c *customerTest) aCustomerWithTheFollowingDetails(table *godog.Table) error {
	rows := table.Rows
	for _, row := range rows {
		id := row.Cells[0].Value
		name := row.Cells[1].Value
		email := row.Cells[2].Value
		cpf := row.Cells[3].Value

		expectedID, _ := uuid.Parse(id)
		given = &entity.Customer{
			ID:    expectedID,
			Name:  name,
			Email: email,
			CPF:   cpf,
		}
	}
	return nil
}

func (c *customerTest) theCustomerIsCreated() error {
	c.repo = &mocks.CustomerRepository{}
	c.service = usecase.NewCustomerUseCase(c.repo)

	c.repo.On("CreateCustomer", ctxDefault, given).Return(wanted, nil).Once()

	c.result, c.err = c.service.CreateCustomer(ctxDefault, given)
	return nil
}

func (c *customerTest) theCreatedCustomerShouldHaveTheFollowingDetails(table *godog.Table) error {
	rows := table.Rows
	for _, row := range rows {
		id := row.Cells[0].Value
		name := row.Cells[1].Value
		email := row.Cells[2].Value
		cpf := row.Cells[3].Value

		expectedID, _ := uuid.Parse(id)
		expected := &entity.Customer{
			ID:    expectedID,
			Name:  name,
			Email: email,
			CPF:   cpf,
		}

		assertExpectedAndActual(assert.EqualValues, expected, c.result, "Expected: %s\nReceived:%s", expected, c.result)
	}
	return nil
}

func (c *customerTest) theCustomerCreationFailsWithAnError() error {
	c.err = errors.New("error")
	return nil
}

func (c *customerTest) anErrorShouldBeReturned() error {
	var t asserter

	assert.Error(&t, c.err)
	assert.Nil(&t, c.result)
	return nil
}

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Name:                 "customer",
		ScenarioInitializer:  InitializeScenario,
		TestSuiteInitializer: InitializeTestSuite,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../../features/customer.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(s *godog.ScenarioContext) {
	customerTest := &customerTest{}

	s.Given(`^a customer with the following details:`, customerTest.aCustomerWithTheFollowingDetails)
	s.When(`^the customer is created$`, customerTest.theCustomerIsCreated)
	s.Step(`^the created customer should have the following details:$`, customerTest.theCreatedCustomerShouldHaveTheFollowingDetails)
	s.Step(`^the customer creation fails with an error$`, customerTest.theCustomerCreationFailsWithAnError)
	s.Step(`^an error should be returned$`, customerTest.anErrorShouldBeReturned)

	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		customerTest.err = nil
		return ctx, nil
	})

	s.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		return ctx, nil
	})
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {}
