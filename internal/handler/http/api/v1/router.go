package v1

import (
	"github.com/go-chi/chi"
)

func (c *Handler) GetVersion() string {
	return "v1"
}

func (c *Handler) GetContentType() string {
	return ""
}
func (c *Handler) GetHandler() *Handler {
	return c
}

func (c *Handler) AddRoutes(r *chi.Mux) {

	r.Group(func(r chi.Router) {
		r.Use(c.Autoriz())
		//получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях
		r.Get("/api/user/orders", c.getOrders)
		//получение текущего баланса счёта баллов лояльности пользователя
		r.Get("/api/user/balance", c.getBalance)
		//запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа
		r.Get("/api/user/withdrawals", c.getWithdrawals)

		//загрузка пользователем номера заказа для расчёт
		r.Post("/api/user/orders", c.postOrder)
		//запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа
		r.Post("/api/user/balance/withdraw", c.postWithdraw)

	})

	r.Group(func(r chi.Router) {
		r.Post("/api/user/register", c.postRegister)
		//аутентификация пользователя
		r.Post("/api/user/login", c.postLogin())

		// Ошибочный путь
		r.Post("/*", c.errorPath)
		r.Get("/*", c.errorPath)
	})

}
