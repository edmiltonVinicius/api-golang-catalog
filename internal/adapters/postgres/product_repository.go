package postgres

import (
	"context"
	"errors"

	"github.com/edmiltonVinicius/go-api-catalog/internal/core/product"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) Create(
	ctx context.Context,
	p product.Product,
) error {

	query := `
		INSERT INTO products (id, sku, name, description, active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		p.ID,
		p.Sku,
		p.Name,
		p.Description,
		p.Active,
		p.CreatedAt,
	)

	if err != nil {
		if isUniqueViolation(err) {
			return product.ErrDuplicateSKU
		}
		return err
	}

	return nil
}

func (r *Repository) FindByID(
	ctx context.Context,
	id string,
) (*product.Product, error) {

	query := `
		SELECT id, sku, name, description, active, created_at
		FROM products
		WHERE id = $1
	`

	row := r.db.QueryRow(ctx, query, id)

	var p product.Product
	err := row.Scan(
		&p.ID,
		&p.Sku,
		&p.Name,
		&p.Description,
		&p.Active,
		&p.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, product.ErrProductNotFound
		}
		return nil, err
	}

	return &p, nil
}
