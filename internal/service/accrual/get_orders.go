package accrualServ

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// Получить строку для расчета балов из БД

func (s *Service) ExecuteGetOrders(ctx context.Context, AccurualTimeReset int) (string, error) {

	var ordNumber string

	rows, err := s.myRepo.GetAll(ctx,
		`with q as (select o.order_number, id from orders o 
		where order_status = 'NEW' or 
			 (order_status = 'PROCESSING' and trunc(EXTRACT( 
				EPOCH from now() -o.update_at)) > $1 ) 
		order by update_at 
		limit 1 
		for update nowait) 
		update orders o 
		 set order_status = 'PROCESSING', 
			 update_at  = now() 
			 from Q  where o.id = q.id 
		returning o.order_number `,
		AccurualTimeReset)

	if err != nil {
		return "", fmt.Errorf("execute select query: %w", err) // внутренняя ошибка сервера
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&ordNumber)

		if err == nil {
			return ordNumber, nil
		}
	}

	if errors.Is(err, sql.ErrNoRows) { // Нет строк
		return "", nil
	}

	if err != nil {
		return "", fmt.Errorf("scan select query: %w", err) // внутренняя ошибка сервера
	}

	err = rows.Err()
	if err != nil {
		return "", fmt.Errorf("fetch rows: %w", err) // внутренняя ошибка сервера
	}
	return ordNumber, nil
}
