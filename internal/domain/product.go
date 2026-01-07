package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          string
	Sku         string
	Name        string
	Description string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateProduct struct {
	Name        string
	Price       int64
	Description string
	Active      bool
}

type ProductDetails struct {
	ID          string
	Sku         string
	Name        string
	Description string
	Active      bool
	Price       PriceInfo
	Stock       StockInfo
}

func generateSKU() string {
	id := uuid.New().String()
	short := strings.ReplaceAll(id[:8], "-", "")
	return fmt.Sprintf("SKU-%s", strings.ToUpper(short))
}

func NewProduct(p CreateProduct) (*Product, error) {
	if p.Name == "" || p.Price <= 0 {
		return nil, errors.New("invalid data to create product")
	}

	return &Product{
		ID:          uuid.New().String(),
		Sku:         generateSKU(),
		Name:        p.Name,
		Description: p.Description,
		Active:      p.Active,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
