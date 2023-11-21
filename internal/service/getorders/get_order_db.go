package getorders

import (
	"context"
	"fmt"
	"time"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

func (s *Service) Execute(ctx context.Context, userID int) ([]*domain.TOrders, error) {

	var tm time.Time
	rows, err := s.myRepo.GetAll(ctx,
		`select o.order_number, o.order_status, o.uploaded_at, o.accrual 
				from orders o 
				where o.user_id = $1 
			 order by o.uploaded_at `,
		userID)

	if err != nil {
		return nil, fmt.Errorf("execute select query: %w", err) // внутренняя ошибка сервера
	}
	defer rows.Close()

	stats := make([]*domain.TOrders, 0)
	for rows.Next() {
		bk := new(domain.TOrders)
		err := rows.Scan(&bk.Number, &bk.Status, &tm, &bk.Accrual)
		if err != nil {
			return nil, fmt.Errorf("scan select query: %w", err) // внутренняя ошибка сервера
		}

		bk.UploadedAt = tm.Format(time.RFC3339)
		err = rows.Err()
		if err != nil {
			return nil, fmt.Errorf("fetch rows: %w", err) // внутренняя ошибка сервера
		}
		stats = append(stats, bk)
	}

	return stats, nil
}
