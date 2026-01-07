package product

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
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

func assembleDetails(product *Product, price *Price, stock *Stock) *ProductDetails {
	details := &ProductDetails{
		ID:          product.ID,
		Sku:         product.Sku,
		Name:        product.Name,
		Description: product.Description,
		Active:      product.Active,
	}

	if price != nil {
		details.Price = PriceInfo{
			Amount:   price.Amount,
			Currency: price.Currency,
		}
	}

	if stock != nil {
		details.Stock = StockInfo{
			Quantity:  stock.Quantity,
			Available: stock.Quantity > 0,
		}
	}

	return details
}

func (s *Service) GetProduct(ctx context.Context, id string) (*ProductDetails, error) {
	g, ctx := errgroup.WithContext(ctx)

	var (
		product *Product
		price   *Price
		stock   *Stock
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

func (s *Service) CreateProduct(ctx context.Context, data CreateProduct) error {
	if data.Name == "" || data.Price <= 0 {
		return errors.New("invalid data to create product")
	}

	err := s.products.Create(ctx, Product{
		ID:          uuid.New().String(),
		Sku:         "",
		Name:        data.Name,
		Description: "",
		Active:      data.Active,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	return err
}
