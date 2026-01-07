package postgres

import (
	"context"
	"errors"

	"github.com/edmiltonVinicius/go-api-catalog/internal/core/product"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) FindByProductID(
	ctx context.Context,
	productID string,
) (*product.Stock, error) {

	query := `
		SELECT product_id, quantity, updated_at
		FROM stocks
		WHERE product_id = $1
	`

	row := r.db.QueryRow(ctx, query, productID)

	var s product.Stock
	err := row.Scan(
		&s.ProductID,
		&s.Quantity,
		&s.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &s, nil
}
