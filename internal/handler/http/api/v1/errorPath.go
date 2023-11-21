package v1

import (
	"net/http"
)

// Страница ошибочным URL
func (c *Handler) errorPath(res http.ResponseWriter, req *http.Request) {

	c.logger.Info("enter in ErrorPage")

	res.WriteHeader(http.StatusNotFound)
}
