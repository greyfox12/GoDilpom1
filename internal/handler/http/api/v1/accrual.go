// получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях
package v1

import (
	"context"
	"fmt"
)

func (c Handler) ExecAccrual(ctx context.Context) error {

	namefunc := "orderNumberGet"

	c.logger.Debug(fmt.Sprintf("enter to %v", namefunc))
	//	ctx1, cancel := context.WithTimeout(context.Background(), time.Duration(c.cfg.HTTP.TimeoutContexDB)*time.Second)
	//	defer cancel()

	bk, err := c.uc.AccrualGetOrder(ctx, c.cfg.AccurualService.TimeReset, c.cfg.AccurualService.Url)
	if err != nil || bk == nil {
		c.logger.Error(fmt.Sprintf("%v: accruaGetOrderExec: %v", namefunc, err))
		return err
	}
	c.logger.Debug(fmt.Sprintf("%v: ordNum=%v, status=%v, accrual=%v, ordResetCN=%v", namefunc, bk.Order, bk.Status, bk.Accrual, bk.OrdResetCn))
	return nil
}
