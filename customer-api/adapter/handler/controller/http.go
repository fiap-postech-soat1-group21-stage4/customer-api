package controller

import (
	"net/http"

	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/adapter/model"
	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCustomer(ctx *gin.Context) {
	var input *model.CustomerRequestDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return
	}

	domain := &entity.Customer{
		Name:  input.Name,
		Email: input.Email,
		CPF:   input.CPF,
	}

	res, err := h.useCase.CreateCustomer(ctx, domain)
	if err != nil {
		return
	}

	output := &model.CustomerResponseDTO{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CPF:       res.CPF,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, output)
}

func (h *Handler) RetrieveCustomer(ctx *gin.Context) {
	cpf := ctx.Param("cpf")

	domain := &entity.Customer{
		CPF: cpf,
	}

	res, err := h.useCase.RetrieveCustomer(ctx, domain)
	if err != nil {
		return
	}

	output := &model.CustomerResponseDTO{
		ID:        res.ID,
		Name:      res.Name,
		Email:     res.Email,
		CPF:       res.CPF,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, output)
}
