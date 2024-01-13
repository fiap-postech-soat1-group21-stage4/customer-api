package usecase

import (
	"context"

	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/entity"
	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/port"
)

type useCaseCustomer struct {
	repository port.CustomerRepository
}

func NewCustomerUseCase(ctm port.CustomerRepository) port.CustomerUseCase {
	return &useCaseCustomer{
		repository: ctm,
	}
}

func (u *useCaseCustomer) CreateCustomer(ctx context.Context, c *entity.Customer) (*entity.Customer, error) {
	res, err := u.repository.CreateCustomer(ctx, c)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *useCaseCustomer) RetrieveCustomer(ctx context.Context, c *entity.Customer) (*entity.Customer, error) {
	res, err := u.repository.RetrieveCustomer(ctx, c)
	if err != nil {
		return nil, err
	}

	return res, nil
}
