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

func (c *Handler) postLogin() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var vRegister domain.TRegister
		namefunc := "postLogin"

		c.logger.Debug(fmt.Sprintf("enter in %v", namefunc))

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.cfg.HTTP.TimeoutContexDB)*time.Second)
		defer cancel()

		c.logger.Debug(fmt.Sprintf("req.Header: %v", req.Header.Get("Content-Encoding")))

		body, err := io.ReadAll(req.Body)
		defer req.Body.Close()

		if err != nil {
			c.logger.Debug(fmt.Sprintf("%v: read body, Body: %v", namefunc, body))
			c.logger.Warn(fmt.Sprintf("%v: read body request: %v", namefunc, err))
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(body) <= 0 {
			c.logger.Warn(fmt.Sprintf("%v: read empty body request", namefunc))
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &vRegister)
		if err != nil {
			c.logger.Warn(fmt.Sprintf("%v: decode json: %v", namefunc, err))
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if vRegister.Login == "" || vRegister.Password == "" {
			c.logger.Warn(fmt.Sprintf("%v: empty login/passwd: %v/%v", namefunc, vRegister.Login, vRegister.Password))
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		c.logger.Debug(fmt.Sprintf("%v: login/passwd: %v/%v", namefunc, vRegister.Login, vRegister.Password))

		ret, err := c.uc.LogingDB(ctx, vRegister.Login, vRegister.Password)
		if err != nil {
			c.logger.Info(fmt.Sprintf("%v: db loging: %v", namefunc, err))
			res.WriteHeader(ret)
			return
		}

		token, err := c.auth.CreateToken(vRegister.Login)
		if err != nil {
			c.logger.Error(fmt.Sprintf("%v: create token: %v", namefunc, err))
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		c.logger.Debug(fmt.Sprintf("%v: create token=%v", namefunc, token))

		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Authorization", "Bearer "+token)
		res.WriteHeader(http.StatusOK) // тк нет возврата тела - сразу ответ без ZIP

		res.Write(nil)
	}
}
