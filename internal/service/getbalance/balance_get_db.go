package getbalance

import (
	"context"
	"fmt"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

func (s *Service) Execute(ctx context.Context, userID int) (*domain.TBallance, error) { //BalanceGetDB(ctx context.Context) {

	var bk domain.TBallance

	rows, err := s.myRepo.GetAll(ctx,
		`select ur.ballans, ur.withdrawn 
	   		from user_ref ur 
	   		where ur.user_id = $1 `, userID)

	if err != nil {
		return nil, fmt.Errorf("execute select query: %w", err) // внутренняя ошибка сервера
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&bk.Current, &bk.Withdrawn)

		if err != nil {
			return nil, fmt.Errorf("scan select query: %w", err) // внутренняя ошибка сервера
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("fetch rows: %w", err) // внутренняя ошибка сервера
	}

	return &bk, nil
}
