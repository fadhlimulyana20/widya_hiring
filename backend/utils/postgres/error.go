package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

// Translate Postgres Error to http status
func TranslatePostgresError(err error) int {
	httpStatus := 500
	if pgError := err.(*pgconn.PgError); errors.Is(err, pgError) {
		switch pgError.Code {
		case "23505":
			httpStatus = 409
		}
	}

	return httpStatus
}
