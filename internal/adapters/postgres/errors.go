package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if err == nil {
		return false
	}

	if ok := errors.As(err, &pgErr); !ok {
		return false
	}

	return pgErr.Code == "23505"
}
