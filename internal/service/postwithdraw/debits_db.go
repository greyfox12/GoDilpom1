package postwithdraw

import (
	"context"
	"fmt"
	"net/http"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

func (s *Service) Execute(ctx context.Context, userID int, vReq domain.TPostWithdraw) (int, error) {

	var pBallans float32

	// запросить балланс
	rows, err := s.myRepo.GetAll(ctx, `select ballans  from  user_ref where  user_id = $1`, userID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("execute select query ballans: %w", err) // внутренняя ошибка сервера
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&pBallans)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("scan select query ballans: %w", err) // внутренняя ошибка сервера
	}

	err = rows.Err()
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("fetch rows ballans: %w", err) // внутренняя ошибка сервера
	}

	if pBallans < vReq.Sum {
		//		c.Loger.OutLogDebug(fmt.Errorf("ballans < summ"))
		return http.StatusPaymentRequired, nil
	}

	// Корректирую балланс и пишу журнал списания
	_, err = s.myRepo.Update(ctx,
		`with ins as (insert into withdraw (user_id, order_number, summa) VALUES($1, $2, $3) returning user_id, summa)
			update user_ref u
					   set withdrawn = withdrawn + $3,
						   ballans  = ballans - $3
					   where u.user_id = (select user_id from ins) `,
		userID, vReq.Order, vReq.Sum)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("execute insert query: %w", err) // внутренняя ошибка сервера
	}

	return http.StatusOK, nil

}
