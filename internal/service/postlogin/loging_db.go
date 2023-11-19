package postlogin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (s *Service) Execute(ctx context.Context, login string) (string, error) {

	// TODO: put your business logic here
	//	panic("implement me")
	var ret string

	rows, err := s.myRepo.GetAll(ctx,
		`select user_pass  from  user_ref 	where  login = $1`,
		login)

	if err != nil {
		return "", fmt.Errorf("execute select query: %w", err) // внутренняя ошибка сервера
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&ret)
	if errors.Is(err, sql.ErrNoRows) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("scan select query: %w", err) // внутренняя ошибка сервера
	}

	err = rows.Err()
	if err != nil {
		return "", fmt.Errorf("fetch rows: %w", err) // внутренняя ошибка сервера
	}
	// fmt.Printf("ret=%v", ret)
	return ret, nil

}
