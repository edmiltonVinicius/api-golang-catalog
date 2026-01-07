package product

import "time"

type Product struct {
	ID          string
	Sku         string
	Name        string
	Description string
	Active      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Price struct {
	ProductID string
	Amount    float64
	Currency  string
	UpdatedAt time.Time
}

type Stock struct {
	ProductID string
	Quantity  int
	UpdatedAt time.Time
}

type PriceInfo struct {
	Amount   float64
	Currency string
}

type StockInfo struct {
	Quantity  int
	Available bool
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

type CreateProduct struct {
	Name   string
	Price  int64
	Active bool
}
