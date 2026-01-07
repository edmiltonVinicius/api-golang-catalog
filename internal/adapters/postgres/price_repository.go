package postgres

import (
	"context"
	"errors"

	"github.com/edmiltonVinicius/go-api-catalog/internal/domain"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetByProductID(
	ctx context.Context,
	productID string,
) (*domain.Price, error) {

	query := `
		SELECT product_id, amount, currency, updated_at
		FROM prices
		WHERE product_id = $1
	`

	row := r.db.QueryRow(ctx, query, productID)

	var p domain.Price
	err := row.Scan(
		&p.ProductID,
		&p.Amount,
		&p.Currency,
		&p.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}
