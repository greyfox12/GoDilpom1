package usecase

import (
	"context"
)

func (u *UseCase) TestLoginDB(ctx context.Context, login string) (int, error) {
	// TODO: put your service call logic here
	//	return &postRegisterUC{}, nil
	//	return "implement UseCase method RegisterDB", nil
	return u.testLogin.ExecuteTestLogin(ctx, login)
}
