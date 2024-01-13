// package controller_test

// import (
// 	"net/http/httptest"
// 	"testing"

// 	handler "github.com/fiap-postech-soat1-group21/customer-api/customer-api/adapter/handler/controller"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestRegisterRoutes(t *testing.T) {
// 	h := handler.NewHandler(nil)
// 	w := httptest.NewRecorder()
// 	_, engine := gin.CreateTestContext(w)
// 	h.RegisterRoutes(engine.Group("/api/v1"))

// 	routesInfo := engine.Routes()
// 	routesMethodAndPath := make([][]string, 0, len(routesInfo))
// 	for _, routeInfo := range routesInfo {
// 		routesMethodAndPath = append(routesMethodAndPath, []string{routeInfo.Method, routeInfo.Path})
// 	}

// 	expectedRoutesMethodAndPath := [][]string{
// 		{"POST", "/api/v1/customer/"},
// 		{"GET", "/api/v1/customer/:cpf"},
// 	}

// 	assert.Equal(t, expectedRoutesMethodAndPath, routesMethodAndPath)
// }

package controller_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/cucumber/godog"
	handler "github.com/fiap-postech-soat1-group21/customer-api/customer-api/adapter/handler/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	h                     *handler.Handler
	w                     *httptest.ResponseRecorder
	engine                *gin.Engine
	expectedRoutesAndPath [][]string
	actualRoutesAndPath   [][]string
)

func aNewCustomerHandler() error {
	return nil
}

func registeringRoutesForAPI(apiVersion string) error {
	h.RegisterRoutes(engine.Group(apiVersion))
	routesInfo := engine.Routes()
	actualRoutesAndPath = make([][]string, 0, len(routesInfo))
	for _, routeInfo := range routesInfo {
		actualRoutesAndPath = append(actualRoutesAndPath, []string{routeInfo.Method, routeInfo.Path})
	}
	return nil
}

func theRegisteredRoutesShouldMatchTheExpectedRoutes(expectedRoutesTable [][]string) error {
	expectedRoutesAndPath = expectedRoutesTable
	assertExpectedAndActual(assert.Equal, expectedRoutesAndPath, actualRoutesAndPath, "registered routes")
	return nil
}

func TestFeaturesHandler(t *testing.T) {
	suite := godog.TestSuite{
		Name:                 "handler",
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

func InitializeScenarioHandler(s *godog.ScenarioContext) {
	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		h = handler.NewHandler(nil)
		w = httptest.NewRecorder()
		engine = gin.New()
		expectedRoutesAndPath = nil
		actualRoutesAndPath = nil
		return ctx, nil
	})

	s.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		return ctx, nil
	})

	s.Given(`^a new customer handler`, aNewCustomerHandler)
	s.When(`^registering routes for API version "([^"]*)"`, registeringRoutesForAPI)
	s.Then(`^the registered routes should match the expected routes:`, theRegisteredRoutesShouldMatchTheExpectedRoutes)
}

func InitializeTestSuiteHandler(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() { fmt.Println("Get the party started!") })
}
