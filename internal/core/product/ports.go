package product

import "context"

type ProductRepository interface {
	Create(ctx context.Context, product Product) error
	FindByID(ctx context.Context, id string) (*Product, error)
}

type PriceRepository interface {
	GetByProductID(ctx context.Context, productID string) (*Price, error)
}

type StockRepository interface {
	FindByProductID(ctx context.Context, productID string) (*Stock, error)
}
