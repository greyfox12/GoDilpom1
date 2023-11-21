package myrepo

import (
	"context"
)

// Выполнение скриптов
func (r *Repository) Update(ctx context.Context, Script string, ids ...any) (int, error) {
	// TODO: put your repository logic here
	//	panic("implement me")

	result, err := r.GetConn().ExecContext(ctx, Script, ids...)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err == nil {
		return int(rows), nil
	}

	return 0, err
}
