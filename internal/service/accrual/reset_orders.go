package accrualServ

import (
	"context"
	"fmt"
)

// Сборсить не обработанные задания в ночальное состояние
// конкретное и все с истекшим тайаутом на обработку

func (s *Service) ExecuteResetOrders(ctx context.Context, orderNum string, AccurualTimeReset int) (int, error) {

	rows, err := s.myRepo.Update(ctx,
		`UPDATE orders SET order_status = 'NEW', update_at = now() 
			 WHERE order_number = $1 OR (order_status = 'PROCESSING' and trunc(EXTRACT( 
			 EPOCH from now() - update_at)) > $2 )`,
		orderNum, AccurualTimeReset)

	if err != nil {
		return 0, fmt.Errorf("execute select query: %w", err) // внутренняя ошибка сервера
	}

	return rows, nil
}
