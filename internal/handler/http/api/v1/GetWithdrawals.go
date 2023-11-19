// Получение списка списания баллов
package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (c *Handler) getWithdrawals(res http.ResponseWriter, req *http.Request) {

	namefunc := "getWithdrawals"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.cfg.HTTP.TimeoutContexDB)*time.Second)
	defer cancel()

	str, err := c.uc.WithdrawalsGetDB(ctx, userID)
	if err != nil {
		c.logger.Warn(fmt.Sprintf("%v: withdrawalsGetDB: %v", namefunc, err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(str) == 0 {
		c.logger.Info(fmt.Sprintf("%v: withdrawalsGetDB: empty", namefunc))
		res.WriteHeader(http.StatusNoContent)
		return //204
	}

	jsonData, err := json.Marshal(str)
	if err != nil {
		c.logger.Warn(fmt.Sprintf("%v: convert JSON: %v", namefunc, err))
		res.WriteHeader(http.StatusInternalServerError)
		return //500
	}

	c.logger.Debug(fmt.Sprintf("%v login: %v return: %v", namefunc, login, str))
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}
