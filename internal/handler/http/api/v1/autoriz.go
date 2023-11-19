package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Autoriz — middleware-авторизация для входящих HTTP-запросов.

func (c *Handler) Autoriz() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			c.logger.Debug("enter in Autoriz")
			// Пропускаю авторизацию для регистрации и логина
			//		if r.URL.String() == "/api/user/register" || r.URL.String() == "/api/user/login" {
			//			next.ServeHTTP(w, r)
			//			return
			//		}

			// Получаю токен авторизации
			login, err := c.auth.CheckAuth(r.Header.Get("Authorization"))
			if err != nil {
				c.logger.Warn(fmt.Sprintf("autorization: error autorization: %v", err))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			c.logger.Debug(fmt.Sprintf("autorization: login %v ", login))

			// Добавляю логин для дальнейшего использования в хендлере
			r.Header.Add("LoginUser", login)

			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.cfg.HTTP.TimeoutContexDB)*time.Second)
			defer cancel()

			userID, err := c.uc.TestLoginDB(ctx, login)
			if err != nil {
				c.logger.Warn(fmt.Sprintf("autorization: login %v: error get userID : %v", login, err))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// Добавляю ID пользователя для дальнейшего использования в хендлере
			r.Header.Add("UserID", fmt.Sprint(userID))

			next.ServeHTTP(w, r)
		})
	}
}
