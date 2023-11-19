package getwithdrawals

import (
	"context"
	"fmt"
	"time"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

func (s *Service) Execute(ctx context.Context, userID int) ([]*domain.TWithdrawals, error) {

	var tm time.Time
	rows, err := s.myRepo.GetAll(ctx,
		`select w.order_number, w.summa, w.uploaded_at
		   from withdraw w
	       where w.user_id = $1
		 order by w.uploaded_at`,
		userID)

	if err != nil {
		return nil, fmt.Errorf("execute select query: %w", err) // внутренняя ошибка сервера
	}
	defer rows.Close()

	stats := make([]*domain.TWithdrawals, 0)
	for rows.Next() {
		bk := new(domain.TWithdrawals)
		err := rows.Scan(&bk.Order, &bk.Sum, &tm)
		if err != nil {
			return nil, fmt.Errorf("scan select query: %w", err) // внутренняя ошибка сервера
		}

		bk.ProcessedAt = tm.Format(time.RFC3339)
		err = rows.Err()
		if err != nil {
			return nil, fmt.Errorf("fetch rows: %w", err) // внутренняя ошибка сервера
		}
		stats = append(stats, bk)
	}

	return stats, nil
}
