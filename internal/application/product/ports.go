package product

import (
	"context"

	"github.com/edmiltonVinicius/go-api-catalog/internal/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product domain.Product) error
	FindByID(ctx context.Context, id string) (*domain.Product, error)
}

type PriceRepository interface {
	GetByProductID(ctx context.Context, productID string) (*domain.Price, error)
}

type StockRepository interface {
	FindByProductID(ctx context.Context, productID string) (*domain.Stock, error)
}
