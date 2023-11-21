package postorder

import (
	"context"
	"fmt"
	"net/http"
)

func (s *Service) Execute(ctx context.Context, pUserID int, ordNum string) (int, error) {

	var userID int
	var userIDOrd int

	// Загрузка номера
	userID, err := s.myRepo.Update(ctx,
		`INSERT INTO orders(user_id, order_number, order_status) 
				   VALUES ($2, $1,'NEW') 
				   ON CONFLICT (order_number) DO NOTHING
				   returning user_id `,
		ordNum, pUserID)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("execute insert query: %w", err) // внутренняя ошибка сервера
	}

	if userID > 0 {
		return http.StatusAccepted, nil
	} // Записб добавлена

	// Запись конфликтует. Ищу причину
	rows, err := s.myRepo.GetAll(ctx,
		`select o.user_id, coalesce(u.user_id, -1)
				 from orders o 
				 left join user_ref u on u.user_id =$2 and u.user_id =o.user_id 
			   where o.order_number = $1`,
		ordNum, pUserID)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("execute select query: %w", err) // внутренняя ошибка сервера
	}
	defer rows.Close()

	if rows.Next() {

		err = rows.Scan(&userIDOrd, &userID)
		if err != nil {
			return http.StatusInternalServerError, fmt.Errorf("scan select query: %w", err) // внутренняя ошибка сервера
		}

		if userID == -1 {
			//			c.Loger.OutLogDebug(fmt.Errorf("order load othes user"))
			return http.StatusConflict, nil // загружено другим пользователем
		}
		//		c.Loger.OutLogDebug(fmt.Errorf("order load this user"))
		return http.StatusOK, nil // загружено этим пользователем
	}

	if err = rows.Err(); err != nil {
		//	c.Loger.OutLogError(fmt.Errorf("fetch rows: %w", err))
		return http.StatusInternalServerError, err // внутренняя ошибка сервера
	}

	return http.StatusInternalServerError, nil // Непонятная ошибка
}
