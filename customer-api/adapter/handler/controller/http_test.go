package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/adapter/handler/controller"
	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/adapter/model"
	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/entity"
	mocks "github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/port/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	customerModelInput = &model.CustomerRequestDTO{
		Name:  "João",
		CPF:   "12312312312",
		Email: "joao@email.com",
	}

	customerEntityInput = &entity.Customer{
		Name:  "João",
		CPF:   "12312312312",
		Email: "joao@email.com",
	}

	customerEntityOutput = &entity.Customer{
		ID:    uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
		Name:  "João",
		CPF:   "12312312312",
		Email: "joao@email.com",
	}

	customerModelOutput = &model.CustomerResponseDTO{
		ID:    uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
		Name:  "João",
		CPF:   "12312312312",
		Email: "joao@email.com",
	}

	retrievePath = "/customer/12312312312"

	cpf = &entity.Customer{
		CPF: "12312312312",
	}
)

type handlerContext struct {
	handler *controller.Handler
	w       *httptest.ResponseRecorder
	req     *http.Request
	err     error
	body    []byte
}

func (h *handlerContext) theFollowingCustomerDetails(table *godog.Table) error {
	customerModelInput = &model.CustomerRequestDTO{
		Name:  table.Rows[1].Cells[0].Value,
		CPF:   table.Rows[1].Cells[1].Value,
		Email: table.Rows[1].Cells[2].Value,
	}
	return nil
}

func (h *handlerContext) aRequestIsMadeToCreateTheCustomer() error {
	jsonBytes, err := json.Marshal(customerModelInput)
	if err != nil {
		return err
	}

	h.req = httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonBytes))
	h.w = httptest.NewRecorder()

	ctxGin, _ := gin.CreateTestContext(h.w)
	ctxGin.Request = h.req

	usecaseMock := &mocks.CustomerUseCase{}
	usecaseMock.On("CreateCustomer", ctxGin, customerEntityInput).Return(customerEntityOutput, h.err).Once()

	h.handler = controller.NewHandler(usecaseMock)
	h.handler.CreateCustomer(ctxGin)

	res := h.w.Result()
	defer res.Body.Close()
	h.body = h.w.Body.Bytes()

	return nil
}

func (h *handlerContext) theResponseShouldHaveStatusCode(statusCode int) error {
	return assertExpectedAndActual(assert.Equal, statusCode, h.w.Code, "status code")
}

func (h *handlerContext) theResponseBodyShouldMatchTheExpectedCustomerDetails() error {
	wantGot, err := json.Marshal(customerModelOutput)
	if err != nil {
		return err
	}

	return assertExpectedAndActual(assert.Equal, string(wantGot), strings.TrimSuffix(h.w.Body.String(), "\n"), "response body")
}

func (h *handlerContext) aCustomerWithCPFExists(cpf string) error {
	return nil
}

func (h *handlerContext) aRequestIsMadeToRetrieveTheCustomer() error {
	req := httptest.NewRequest(http.MethodGet, retrievePath, nil)
	h.w = httptest.NewRecorder()

	_, engine := gin.CreateTestContext(h.w)

	usecaseMock := &mocks.CustomerUseCase{}
	usecaseMock.
		On("RetrieveCustomer", mock.AnythingOfType("*gin.Context"), cpf).Return(customerEntityOutput, h.err).Once()

	h.handler = controller.NewHandler(usecaseMock)

	engine.GET("/customer/:cpf", h.handler.RetrieveCustomer)
	engine.ServeHTTP(h.w, req)

	res := h.w.Result()
	defer res.Body.Close()
	h.body = h.w.Body.Bytes()

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
		Name:                 "http",
		ScenarioInitializer:  InitializeScenario,
		TestSuiteInitializer: InitializeTestSuite,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../../features/http.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(s *godog.ScenarioContext) {
	h := &handlerContext{}
	s.Given(`^the following customer details`, h.theFollowingCustomerDetails)
	s.When(`^a request is made to create the customer`, h.aRequestIsMadeToCreateTheCustomer)
	s.Then(`^the response should have status code (\d+)`, h.theResponseShouldHaveStatusCode)
	s.Step(`^the response body should match the expected customer details`, h.theResponseBodyShouldMatchTheExpectedCustomerDetails)
	s.Given(`^a customer with CPF "([^"]*)" exists`, h.aCustomerWithCPFExists)
	s.When(`^a request is made to retrieve the customer`, h.aRequestIsMadeToRetrieveTheCustomer)

	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		h.err = nil
		return ctx, nil
	})

	s.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		return ctx, nil
	})
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {}
