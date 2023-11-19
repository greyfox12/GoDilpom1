package usecase

import (
	"context"
)

func (u *UseCase) RegisterDB(ctx context.Context, login string, passwd string) (int, error) {
	// TODO: put your service call logic here
	//	return &postRegisterUC{}, nil
	//	return "implement UseCase method RegisterDB", nil
	return u.postRegister.Execute(ctx, login, passwd)
}
