package usecase

import (
	"context"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

func (u *UseCase) BalanceGetDB(ctx context.Context, userID int) (*domain.TBallance, error) {
	// TODO: put your service call logic here
	//	return "implement UseCase method BalanceGetDB", nil
	return u.getBalance.Execute(ctx, userID)
}
