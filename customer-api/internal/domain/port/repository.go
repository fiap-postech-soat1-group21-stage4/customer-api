package port

import (
	"context"

	"github.com/fiap-postech-soat1-group21/customer-api/customer-api/internal/domain/entity"
)

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, c *entity.Customer) (*entity.Customer, error)
	RetrieveCustomer(ctx context.Context, c *entity.Customer) (*entity.Customer, error)
}
