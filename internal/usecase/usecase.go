// Code generated by goro;

package usecase

import (
	"context"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

// This file was generated by the goro tool.

type getOrders interface {
	Execute(context.Context, int) ([]*domain.TOrders, error)
}

type getBalance interface {
	Execute(context.Context, int) (*domain.TBallance, error)
}

type getWithdrawals interface {
	Execute(ctx context.Context, userID int) ([]*domain.TWithdrawals, error)
}

type postOrder interface {
	Execute(ctx context.Context, UserID int, ordNum string) (int, error)
}

type postWithdraw interface {
	Execute(ctx context.Context, userID int, vReq domain.TPostWithdraw) (int, error)
}

type postRegister interface {
	Execute(context.Context, string, string) (int, error)
}

type postLogin interface {
	Execute(context.Context, string) (string, error)
}

type testLogin interface {
	ExecuteTestLogin(context.Context, string) (int, error)
}

type migrateShema interface {
	Execute() error
}

// Система начисления баллов
type accrual interface {
	ExecuteResetOrders(ctx context.Context, orderNum string, AccurualTimeReset int) (int, error)
	ExecuteSaveOrders(ctx context.Context, order string, status string, accrual float32) error
	ExecuteGetOrders(ctx context.Context, AccurualTimeReset int) (string, error)
	ExecuteGetRequestHTTP(ctx context.Context, orderNum string, accrualURL string) (*domain.TAccrualReq, error)
}

type UseCase struct {
	getOrders      getOrders
	getBalance     getBalance
	getWithdrawals getWithdrawals
	postOrder      postOrder
	postWithdraw   postWithdraw
	postRegister   postRegister
	postLogin      postLogin
	testLogin      testLogin
	migrateShema   migrateShema
	accrual        accrual
}

func NewUseCase(getOrders getOrders, getBalance getBalance, getWithdrawals getWithdrawals, postOrder postOrder,
	postWithdraw postWithdraw, postRegister postRegister, postLogin postLogin, testLogin testLogin,
	mirateShema migrateShema, accrual accrual) *UseCase {
	return &UseCase{
		getOrders:      getOrders,
		getBalance:     getBalance,
		getWithdrawals: getWithdrawals,
		postOrder:      postOrder,
		postWithdraw:   postWithdraw,
		postRegister:   postRegister,
		postLogin:      postLogin,
		testLogin:      testLogin,
		migrateShema:   mirateShema,
		accrual:        accrual,
	}
}
