// Загрузка Номеров Заказов
package v1

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func (c *Handler) postOrder(res http.ResponseWriter, req *http.Request) {

	namefunc := "postOrder"

	c.logger.Debug(fmt.Sprintf("enter in %v", namefunc))

	login := req.Header.Get("LoginUser")
	if login == "" {
		c.logger.Warn(fmt.Sprintf("%v: error autorization", namefunc))
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	userID, err := strconv.Atoi(req.Header.Get("UserID"))
	if err != nil {
		c.logger.Warn(fmt.Sprintf("%v: error autorization: convert userID: %v", namefunc, err))
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	if req.Header.Get("Content-Type") != "text/plain" {
		c.logger.Info(fmt.Sprintf("%v: incorrect content-type head: %v", namefunc, req.Header.Get("Content-Type")))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		c.logger.Warn(fmt.Sprintf("%v: read body request: %v", namefunc, err))
		res.WriteHeader(http.StatusUnprocessableEntity) //422
		return
	}
	if len(body) <= 0 {
		c.logger.Warn(fmt.Sprintf("%v: read empty body request", namefunc))
		res.WriteHeader(http.StatusUnprocessableEntity) //422
		return
	}

	// получил номер заказа
	// Проверка корректности
	numeric := regexp.MustCompile(`\d`).MatchString(string(body))
	if !numeric {
		c.logger.Info(fmt.Sprintf("%v: number incorrect symbol: %v", namefunc, string(body)))
		res.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Проверка алгоритмом Луна
	if !c.auth.ValidLunaStr(string(body)) {
		c.logger.Info(fmt.Sprintf("%v: number incorrect luna: %v", namefunc, string(body)))
		res.WriteHeader(http.StatusUnprocessableEntity) //422
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.cfg.HTTP.TimeoutContexDB)*time.Second)
	defer cancel()

	ret, err := c.uc.LoadOrderDB(ctx, userID, string(body))
	if err != nil {
		c.logger.Error(fmt.Sprintf("%v: db loader: %v", namefunc, err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	c.logger.Debug(fmt.Sprintf("%v load ok: number: %v ret:%v", namefunc, string(body), ret))
	res.WriteHeader(ret) // тк нет возврата тела - сразу ответ без ZIP
	res.Write(nil)
}
