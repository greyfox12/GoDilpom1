// получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях
package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (c *Handler) getOrders(res http.ResponseWriter, req *http.Request) {

	namefunc := "getOrders"
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

	str, err := c.uc.GetOrderDB(ctx, userID)
	if err != nil {
		c.logger.Error(fmt.Sprintf("%v: db getOrderDB: %v", namefunc, err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	if len(str) == 0 {
		res.WriteHeader(http.StatusNoContent)
		return // 204
	}

	jsonData, err := json.Marshal(str)
	if err != nil {
		c.logger.Error(fmt.Sprintf("%v: Marshal: %v", namefunc, err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonData)
}
