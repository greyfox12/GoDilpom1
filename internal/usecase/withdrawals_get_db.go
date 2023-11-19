package usecase

import (
	"context"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

func (u *UseCase) WithdrawalsGetDB(ctx context.Context, userID int) ([]*domain.TWithdrawals, error) {
	// TODO: put your service call logic here
	return u.getWithdrawals.Execute(ctx, userID)
}
