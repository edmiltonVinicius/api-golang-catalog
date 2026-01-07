package domain

import "time"

type Stock struct {
	ProductID string
	Quantity  int
	UpdatedAt time.Time
}

type StockInfo struct {
	Quantity  int
	Available bool
}
