package usecase

import (
	"context"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

func (u *UseCase) GetOrderDB(ctx context.Context, userID int) ([]*domain.TOrders, error) {
	// TODO: put your service call logic here
	return u.getOrders.Execute(ctx, userID)
}
