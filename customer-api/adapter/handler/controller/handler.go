package controller

import (
	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/port"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase port.CustomerUseCase
}

func NewHandler(u port.CustomerUseCase) *Handler {
	return &Handler{
		useCase: u,
	}
}

func (h *Handler) RegisterRoutes(routes *gin.RouterGroup) {
	customerRoute := routes.Group("/customer")
	customerRoute.POST("/", h.CreateCustomer)
	customerRoute.GET("/:cpf", h.RetrieveCustomer)
}
