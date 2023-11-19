package accrualServ

import (
	"context"
	"fmt"
)

// Добавить новое начисление баллов

func (s *Service) ExecuteSaveOrders(ctx context.Context, order string, status string, accrual float32) error {

	rows, err := s.myRepo.Update(ctx,
		`with sel as (select o.user_id, order_number  from orders o  where o.order_number = $1), 
			 up1  as (update user_ref 
				set ballans  = ballans  + $3 
				where user_id = (select user_id from sel) 
				returning user_id) 
				update orders 
				  set order_status = $2, 
					  accrual  = $3, 
					  update_at = now() 
				where order_number = $1 
				  and user_id = (select user_id from up1)`,
		order, status, accrual)

	if err != nil {
		return fmt.Errorf("execute update query: %w", err) // внутренняя ошибка сервера
	}

	if rows == 0 {
		return fmt.Errorf("no update rows") // внутренняя ошибка сервера
	}

	return nil
}
