package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

func (u *UseCase) AccrualGetOrder(ctxx context.Context, accualTimeReset int, accrualURL string) (*domain.TAccrualReq, error) {

	namefunc := "accrualGetOrder"

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(12)*time.Second)
	defer cancel()

	ordNum, err := u.accrual.ExecuteGetOrders(ctx, accualTimeReset)
	if err != nil {
		return nil, fmt.Errorf("%v: ExecuteGetOrders: %w", namefunc, err)
	}

	if ordNum == "" {
		//		logger.Logger.Debug(fmt.Sprintf("%v: executeGetOrders: orderNumber is null", namefunc))
		return nil, nil
	}

	bk, err := u.accrual.ExecuteGetRequestHTTP(ctx, ordNum, accrualURL)
	if err != nil {
		_, err1 := u.accrual.ExecuteResetOrders(ctx, ordNum, accualTimeReset)
		//bk.OrdResetCn = cn
		//		logger.Logger.Info(fmt.Errorf("%v: resetOrdersDB: reset %v orders", namefunc, cn))
		if err1 != nil {
			//			logger.Logger.Debug(fmt.Errorf("%v: executeResetOrders: orderNum=%v cn=%v %w", namefunc, ordNum, cn, err))
			return nil, fmt.Errorf("%v: ExecuteGetRequestHTTP=%w resetOrdersDB: orderNum=%v %w", namefunc, err, ordNum, err1)
		}
		return nil, err
	}

	err = u.accrual.ExecuteSaveOrders(ctx, bk.Order, bk.Status, bk.Accrual)
	if err != nil {
		_, err1 := u.accrual.ExecuteResetOrders(ctx, ordNum, accualTimeReset)
		//		bk.OrdResetCn += cn
		//		logger.Logger.Info(fmt.Sprintf("%v: resetOrdersDB: reset %v orders", namefunc, cn))
		if err1 != nil {
			//			logger.Logger.Error(fmt.Sprintf("%v: resetOrdersDB: orderNum=%v %v", namefunc, ordNum, err))
			return nil, fmt.Errorf("%v: ExecuteGetRequestHTTP=%w resetOrdersDB: orderNum=%v %w", namefunc, err, ordNum, err1)
		}
		return nil, err
	}
	return bk, nil
}
