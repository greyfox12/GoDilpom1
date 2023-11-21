package usecase

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (u *UseCase) LogingDB(ctx context.Context, login string, passwd string) (int, error) {
	// TODO: put your service call logic here
	//	return "implement UseCase method LogingDB", nil

	dbHash, err := u.postLogin.Execute(ctx, login)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("db loging: %v", err)
	}

	if dbHash == "" {
		return http.StatusUnauthorized, fmt.Errorf("login %v not found in db", login) //401
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbHash), []byte(passwd)); err != nil {
		return http.StatusUnauthorized, fmt.Errorf("compare password and hash incorrect")
	}

	return 0, nil //u.postLogin.Execute(ctx, login)

}
