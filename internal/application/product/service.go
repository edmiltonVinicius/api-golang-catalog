package product

import (
	"context"
	"errors"

	"github.com/edmiltonVinicius/go-api-catalog/internal/domain"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	products ProductRepository
	prices   PriceRepository
	stocks   StockRepository
}

func NewService(products ProductRepository, prices PriceRepository, stocks StockRepository) *Service {
	return &Service{
		products: products,
		prices:   prices,
		stocks:   stocks,
	}
}

func assembleDetails(product *domain.Product, price *domain.Price, stock *domain.Stock) *domain.ProductDetails {
	details := &domain.ProductDetails{
		ID:          product.ID,
		Sku:         product.Sku,
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
	}

	if price != nil {
		details.Price = domain.PriceInfo{
			Amount:   price.Amount,
			Currency: price.Currency,
		}
	}

	if stock != nil {
		details.Stock = domain.StockInfo{
			Quantity:  stock.Quantity,
			Available: stock.Quantity > 0,
		}
	}

	return details
}

func (s *Service) GetProduct(ctx context.Context, id string) (*domain.ProductDetails, error) {
	g, ctx := errgroup.WithContext(ctx)

	var (
		product *domain.Product
		price   *domain.Price
		stock   *domain.Stock
	)

	g.Go(func() error {
		var err error
		product, err = s.products.FindByID(ctx, id)
		return err
	})

	g.Go(func() error {
		var err error
		price, err = s.prices.GetByProductID(ctx, id)
		return err
	})

	g.Go(func() error {
		var err error
		stock, err = s.stocks.FindByProductID(ctx, id)
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	if product == nil {
		return nil, ErrProductNotFound
	}

	if !product.Active {
		return nil, ErrProductInactive
	}

	return assembleDetails(product, price, stock), nil
}

func (s *Service) CreateProduct(ctx context.Context, d domain.Product) error {
	if d.Name == "" {
		return errors.New("invalid data to create product")
	}

	err := s.products.Create(ctx, d)

	return err
}
