package testlogin

import (
	"context"
	"database/sql"
	"fmt"
)

// Проверка действительности логина. Если есть - вернуть user_id или 0 если нет
func (s *Service) ExecuteTestLogin(ctx context.Context, login string) (int, error) {
	var ret int

	rows, err := s.myRepo.GetAll(ctx, `select user_id  from  user_ref where  login = $1 `, login)

	if err != nil {
		return 0, fmt.Errorf("testLogin function: execute select query: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&ret)
		if err == sql.ErrNoRows {
			return 0, nil
		}

		if err != nil {
			return 0, fmt.Errorf("testLogin function: scan select query: %w", err)
		}
	}
	err = rows.Err()
	if err != nil {
		return 0, fmt.Errorf("testLogin function: fetch rows: %w", err)
	}

	return ret, nil
}
