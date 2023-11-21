package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

func (c *Handler) postRegister(res http.ResponseWriter, req *http.Request) {

	var vRegister domain.TRegister
	namefunc := "postRegister"

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.cfg.HTTP.TimeoutContexDB)*time.Second)
	defer cancel()

	c.logger.Debug(fmt.Sprintf("enter in %s", namefunc))

	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		c.logger.Info(fmt.Sprintf("%v: read body request: %v", namefunc, err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(body) <= 0 {
		c.logger.Info(fmt.Sprintf("%v: read empty body request", namefunc))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &vRegister)
	if err != nil {
		c.logger.Info(fmt.Sprintf("%v: decode json: %v", namefunc, err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if vRegister.Login == "" || vRegister.Password == "" {
		c.logger.Info(fmt.Sprintf("%v login or password empty: login/passwd: %v/%v", namefunc, vRegister.Login, vRegister.Password))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	c.logger.Debug(fmt.Sprintf("%v vRegister =%v", namefunc, vRegister))

	if vRegister.PasswordHash, err = c.auth.GetBcryptHash(vRegister.Password); err != nil {
		c.logger.Error(fmt.Sprintf("%v hash password: %v", namefunc, err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	ret, err := c.uc.RegisterDB(ctx, vRegister.Login, vRegister.PasswordHash)
	if err != nil {
		c.logger.Error(fmt.Sprintf("%v: db register: %v", namefunc, err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if ret == http.StatusOK {
		token, err := c.auth.CreateToken(vRegister.Login)
		if err != nil {
			c.logger.Error(fmt.Sprintf("%v: create token: %v", namefunc, err))
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		c.logger.Debug(fmt.Sprintf("%v: create token=%v", namefunc, token))

		res.Header().Set("Authorization", "Bearer "+token)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(ret) // тк нет возврата тела - сразу ответ без ZIP
	res.Write(nil)
}

/*
// Регистрация пользователя
func (c *BaseController) registerDB(ctx context.Context, login string, passwd string) (int, error) {

	rows, err := c.DB.ResendDB(ctx,
		`INSERT INTO user_ref (login, user_pass) VALUES ($1,  $2)
		  ON CONFLICT (login) DO NOTHING
		  returning user_id`,
		login, passwd)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("execute insert query: %w", err) // внутренняя ошибка сервера
	}

	if rows == 0 {
		return http.StatusConflict, nil
	}
	return http.StatusOK, nil
}
*/
