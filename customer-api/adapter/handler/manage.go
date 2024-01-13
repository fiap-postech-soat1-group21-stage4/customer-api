package manage

import (
	c "github.com/fiap-postech-soat1-group21/customer-api/customer-api/adapter/handler/controller"
	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/port"
	"github.com/gin-gonic/gin"
)

type apps interface {
	RegisterRoutes(routes *gin.RouterGroup)
}

type Manage struct {
	customer apps
}

type UseCases struct {
	Customer port.CustomerUseCase
}

func New(uc *UseCases) *Manage {

	customerHandler := c.NewHandler(uc.Customer)

	return &Manage{
		customer: customerHandler,
	}
}

func (m *Manage) RegisterRoutes(group *gin.RouterGroup) {
	m.customer.RegisterRoutes(group)
}
