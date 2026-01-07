package domain

import "time"

type Price struct {
	ProductID string
	Amount    float64
	Currency  string
	UpdatedAt time.Time
}

type PriceInfo struct {
	Amount   float64
	Currency string
}
