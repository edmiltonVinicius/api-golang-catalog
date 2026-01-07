package domain

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

type ProductDetails struct {
	ID          string
	Sku         string
	Name        string
	Description string
	Active      bool
	Price       PriceInfo
	Stock       StockInfo
}
