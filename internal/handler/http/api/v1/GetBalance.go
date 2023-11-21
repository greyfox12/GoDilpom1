package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Получить балланс пользователя
func (c *Handler) getBalance(res http.ResponseWriter, req *http.Request) {

	namefunc := "getBalance"
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

	//	fmt.Printf("c.cfg.HTTP.TimeoutContexDB %v\n", c.cfg.HTTP.TimeoutContexDB)
	bk, err := c.uc.BalanceGetDB(ctx, userID)
	if err != nil {
		c.logger.Warn(fmt.Sprintf("%v: balanceGetDB: %v", namefunc, err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(bk)
	if err != nil {
		c.logger.Warn(fmt.Sprintf("%v: convert JSON: %v", namefunc, err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}
