package product

import "errors"

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductInactive = errors.New("product is inactive")
	ErrInvalidPrice    = errors.New("invalid price")
	ErrDuplicateSKU    = errors.New("duplicate SKU")
)
