// Списание баллов с баланса
package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/greyfox12/GoDilpom1/pkg/domain"
)

func (c *Handler) postWithdraw(res http.ResponseWriter, req *http.Request) {

	namefunc := "postWithdraw"
	var err error
	var vRequest domain.TPostWithdraw

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

	if userID == 0 { // Логин не найден в базе
		c.logger.Warn(fmt.Sprintf("%v: error autorization", namefunc))
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	if req.Header.Get("Content-Type") != "application/json" {
		c.logger.Info(fmt.Sprintf("%v: incorrect content-type head: %v", namefunc, req.Header.Get("Content-Type")))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		c.logger.Warn(fmt.Sprintf("%v: read body request: %v", namefunc, err))
		res.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if len(body) <= 0 {
		c.logger.Warn(fmt.Sprintf("%v: read empty body request", namefunc))
		res.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = json.Unmarshal(body, &vRequest)
	if err != nil {
		c.logger.Warn(fmt.Sprintf("%v: decode json: %v", namefunc, err))
		res.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	if vRequest.Order == "" || vRequest.Sum == 0 {
		c.logger.Warn(fmt.Sprintf("%v: empty order/sum: %v/%v", namefunc, vRequest.Order, vRequest.Sum))
		res.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// получил номер заказа
	// Проверка корректности
	numeric := regexp.MustCompile(`\d`).MatchString(vRequest.Order)
	if !numeric {
		c.logger.Warn(fmt.Sprintf("%v: number incorrect: %v", namefunc, vRequest.Order))
		res.WriteHeader(http.StatusUnprocessableEntity) //422
		return
	}

	// Проверка алгоритмом Луна
	/*		if !hash.ValidLunaStr(vRequest.Order) {
				logmy.OutLog(fmt.Sprintf("debitingpage: number incorrect: %v", vRequest.Order))
				res.WriteHeader(422)
				return
			}
	*/

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.cfg.HTTP.TimeoutContexDB)*time.Second)
	defer cancel()

	retCod, err := c.uc.DebitsDB(ctx, userID, vRequest)
	if err != nil {
		c.logger.Error(fmt.Sprintf("%v: db debits: %v", namefunc, err))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(retCod) // тк нет возврата тела - сразу ответ без ZIP
	res.Write(nil)
}
