package myrepo

import (
	"context"
	"database/sql"
)

// Выполняю SELECT в Базе данных
func (r *Repository) GetAll(ctx context.Context, sqlQuery string, ids ...any) (*sql.Rows, error) {

	rows, err := r.GetConn().QueryContext(ctx, sqlQuery, ids...)
	if err == nil {
		return rows, nil
	}

	return nil, err
}
