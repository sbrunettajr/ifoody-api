package repository

import (
	"context"
	"database/sql"
)

type scanner interface {
	Scan(dest ...any) error
}

func findByID[T any](context context.Context, db *sql.DB, query string, ID uint32, function func(scan scanner) (T, error)) (T, error) {
	row := db.QueryRowContext(
		context,
		query,
		ID,
	)

	var t T

	if row.Err() != nil {
		return t, row.Err()
	}

	result, err := function(row)
	if err != nil {
		return t, err
	}

	return result, nil
}

func findByUUID[T any](context context.Context, db *sql.DB, query string, UUID string, function func(scan scanner) (T, error)) (T, error) {
	row := db.QueryRowContext(
		context,
		query,
		UUID,
	)

	var t T

	if row.Err() != nil {
		return t, row.Err()
	}

	result, err := function(row)
	if err != nil {
		return t, err
	}

	return result, nil
}

func parseEntities[T any](rows *sql.Rows, function func(scan scanner) (T, error)) ([]T, error) {
	response := make([]T, 0)

	for rows.Next() {
		r, err := function(rows)
		if err != nil {
			return nil, err
		}
		response = append(response, r)
	}
	return response, nil
}
