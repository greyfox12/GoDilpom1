package postregister

import (
	"context"
	"fmt"
	"net/http"
)

// func (s *Service) RegisterDB(ctx context.Context, login string, passwd string) (int error) {
func (s *Service) Execute(ctx context.Context, login string, passwd string) (int, error) {

	// TODO: put your business logic here
	//	panic("implement me")
	rows, err := s.myRepo.Update(ctx,
		`INSERT INTO user_ref (login, user_pass) VALUES ($1,  $2) 
			ON CONFLICT (login) DO NOTHING 
			returning user_id`,
		login, passwd)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("execute insert query: %w", err) // внутренняя ошибка сервера
	}

	if rows == 0 {
		return http.StatusConflict, nil
	}
	return http.StatusOK, nil
}
