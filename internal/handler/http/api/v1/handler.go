package v1

import (
	"context"

	"github.com/greyfox12/GoDilpom1/internal/config"
	"github.com/greyfox12/GoDilpom1/pkg/cripto"
	"github.com/greyfox12/GoDilpom1/pkg/domain"
	"github.com/greyfox12/GoDilpom1/pkg/logger"
)

type UseCase interface {
	GetOrderDB(ctx context.Context, userID int) ([]*domain.TOrders, error)
	BalanceGetDB(ctx context.Context, userID int) (*domain.TBallance, error)
	WithdrawalsGetDB(ctx context.Context, userID int) ([]*domain.TWithdrawals, error)
	LoadOrderDB(ctx context.Context, UserID int, ordNum string) (int, error)
	DebitsDB(ctx context.Context, userID int, vReq domain.TPostWithdraw) (int, error)
	RegisterDB(ctx context.Context, login string, passwd string) (int, error)
	LogingDB(ctx context.Context, login string, passwd string) (int, error)
	TestLoginDB(ctx context.Context, login string) (int, error)
	AccrualGetOrder(ctx context.Context, accualTimeReset int, URL string) (*domain.TAccrualReq, error)
}

type Handler struct {
	uc     UseCase
	logger logger.Logger
	cfg    config.Config
	auth   cripto.Auth
}

func NewHandler(uc UseCase, logger logger.Logger, cfg config.Config) *Handler {
	return &Handler{uc: uc, logger: logger, cfg: cfg}
}
