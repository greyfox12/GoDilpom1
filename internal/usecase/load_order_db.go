package usecase

import (
	"context"
)

func (u *UseCase) LoadOrderDB(ctx context.Context, UserID int, ordNum string) (int, error) {
	// TODO: put your service call logic here

	return u.postOrder.Execute(ctx, UserID, ordNum)
}
