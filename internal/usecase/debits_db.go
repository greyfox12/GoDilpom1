package usecase

import (
	"context"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

// Списание балов

func (c *UseCase) DebitsDB(ctx context.Context, userID int, vReq domain.TPostWithdraw) (int, error) {

	return c.postWithdraw.Execute(ctx, userID, vReq)
}
