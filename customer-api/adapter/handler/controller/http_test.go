// package controller_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/adapter/handler/controller"
// 	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/adapter/model"
// 	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/entity"
// 	mocks "github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/port/mocks"
// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// var (
// 	customerModelInput = &model.CustomerRequestDTO{
// 		Name:  "João",
// 		CPF:   "12312312312",
// 		Email: "joao@email.com",
// 	}

// 	customerEntityInput = &entity.Customer{
// 		Name:  "João",
// 		CPF:   "12312312312",
// 		Email: "joao@email.com",
// 	}

// 	customerEntityOutput = &entity.Customer{
// 		ID:    uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
// 		Name:  "João",
// 		CPF:   "12312312312",
// 		Email: "joao@email.com",
// 	}

// 	customerModelOutput = &model.CustomerResponseDTO{
// 		ID:    uuid.MustParse("8c2b51bf-7b4c-4a4b-a024-f283576cf190"),
// 		Name:  "João",
// 		CPF:   "12312312312",
// 		Email: "joao@email.com",
// 	}

// 	retrievePath = "/customer/12312312312"

// 	cpf = &entity.Customer{
// 		CPF: "12312312312",
// 	}
// )

// func TestCreateCustomer(t *testing.T) {
// 	t.Run("when everything goes as expected; should return response 200 and body", func(t *testing.T) {
// 		jsonBytes, err := json.Marshal(customerModelInput)
// 		fmt.Println(jsonBytes)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest(
// 			http.MethodPost,
// 			"/",
// 			bytes.NewBuffer(jsonBytes))
// 		w := httptest.NewRecorder()

// 		ctxGin, _ := gin.CreateTestContext(w)
// 		ctxGin.Request = req

// 		usecaseMock := mocks.NewCustomerUseCase(t)
// 		usecaseMock.On("CreateCustomer", ctxGin, customerEntityInput).Return(customerEntityOutput, nil).Once()

// 		handler := controller.NewHandler(usecaseMock)

// 		handler.CreateCustomer(ctxGin)

// 		res := w.Result()
// 		defer res.Body.Close()
// 		got, err := json.Marshal(customerModelOutput)

// 		assert.NoError(t, err)
// 		assert.EqualValues(t, strings.TrimSuffix(w.Body.String(), "\n"), string(got))
// 		assert.Equal(t, http.StatusCreated, res.StatusCode)
// 		usecaseMock.AssertExpectations(t)
// 	})

// 	t.Run("when body is invalid; should return response 400", func(t *testing.T) {
// 		req := httptest.NewRequest(http.MethodPost, "/customer", bytes.NewBuffer([]byte(`{>}`)))
// 		w := httptest.NewRecorder()
// 		_, engine := gin.CreateTestContext(w)

// 		usecaseMock := mocks.NewCustomerUseCase(t)

// 		handler := controller.NewHandler(usecaseMock)

// 		engine.POST("/", handler.CreateCustomer)
// 		engine.ServeHTTP(w, req)
// 		assert.Equal(t, http.StatusNotFound, w.Code)
// 		usecaseMock.AssertExpectations(t)
// 	})

// 	t.Run("when use case return error; should return error", func(t *testing.T) {
// 		jsonBytes, err := json.Marshal(customerModelInput)
// 		fmt.Println(jsonBytes)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest(
// 			http.MethodPost,
// 			"/",
// 			bytes.NewBuffer(jsonBytes))

// 		w := httptest.NewRecorder()

// 		ctxGin, _ := gin.CreateTestContext(w)
// 		ctxGin.Request = req

// 		usecaseMock := mocks.NewCustomerUseCase(t)
// 		wantError := errors.New("error")
// 		usecaseMock.On("CreateCustomer", ctxGin, customerEntityInput).Return(nil, wantError).Once()

// 		handler := controller.NewHandler(usecaseMock)

// 		handler.CreateCustomer(ctxGin)

// 		assert.ErrorContains(t, wantError, "error")
// 		usecaseMock.AssertExpectations(t)
// 	})
// }

// func TestRetrieveCustomer(t *testing.T) {
// 	t.Run("when everything goes as expected; should return response 200 and body", func(t *testing.T) {

// 		req := httptest.NewRequest(http.MethodGet, retrievePath, nil)
// 		w := httptest.NewRecorder()

// 		_, engine := gin.CreateTestContext(w)

// 		usecaseMock := mocks.NewCustomerUseCase(t)
// 		usecaseMock.
// 			On("RetrieveCustomer", mock.AnythingOfType("*gin.Context"), cpf).Return(customerEntityOutput, nil).Once()

// 		handler := controller.NewHandler(usecaseMock)

// 		engine.GET("/customer/:cpf", handler.RetrieveCustomer)
// 		engine.ServeHTTP(w, req)

// 		res := w.Result()
// 		defer res.Body.Close()
// 		wantGot, err := json.Marshal(customerModelOutput)
// 		assert.NoError(t, err)

// 		assert.EqualValues(t, strings.TrimSuffix(w.Body.String(), "\n"), string(wantGot))
// 		assert.Equal(t, http.StatusOK, w.Code)
// 	})
// 	t.Run("when use case return error; should return response error", func(t *testing.T) {
// 		req := httptest.NewRequest(http.MethodGet, retrievePath, nil)
// 		w := httptest.NewRecorder()
// 		ctxGin, _ := gin.CreateTestContext(w)
// 		ctxGin.Request = req

// 		wantError := errors.New("error")

// 		usecaseMock := mocks.NewCustomerUseCase(t)
// 		usecaseMock.
// 			On("RetrieveCustomer", ctxGin, mock.AnythingOfType("*entity.Customer")).Return(nil, wantError).Once()

// 		handler := controller.NewHandler(usecaseMock)
// 		handler.RetrieveCustomer(ctxGin)

// 		assert.ErrorContains(t, wantError, "error")
// 		usecaseMock.AssertExpectations(t)
// 	})
// }

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
			Paths:    []string{"../../../features/"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(s *godog.ScenarioContext) {
	h := &handlerContext{}
	s.Step(`^the following customer details`, h.theFollowingCustomerDetails)
	s.Step(`^a request is made to create the customer`, h.aRequestIsMadeToCreateTheCustomer)
	s.Step(`^the response should have status code (\d+)`, h.theResponseShouldHaveStatusCode)
	s.Step(`^the response body should match the expected customer details`, h.theResponseBodyShouldMatchTheExpectedCustomerDetails)
	s.Given(`^a customer with CPF "([^"]*)" exists`, h.aCustomerWithCPFExists)
	s.Step(`^a request is made to retrieve the customer`, h.aRequestIsMadeToRetrieveTheCustomer)

	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		h.err = nil
		return ctx, nil
	})

	s.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		return ctx, nil
	})
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {}
