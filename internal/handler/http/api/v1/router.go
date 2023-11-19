package v1

import (
	"github.com/go-chi/chi"
)

func (h *Handler) GetVersion() string {
	return "v1"
}

func (h *Handler) GetContentType() string {
	return ""
}
func (h *Handler) GetHandler() *Handler {
	return h
}

func (h *Handler) AddRoutes(r *chi.Mux) {

	r.Group(func(r chi.Router) {
		r.Use(h.Autoriz())
		//получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях
		r.Get("/api/user/orders", h.getOrders)
		//получение текущего баланса счёта баллов лояльности пользователя
		r.Get("/api/user/balance", h.getBalance)
		//запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа
		r.Get("/api/user/withdrawals", h.getWithdrawals)

		//загрузка пользователем номера заказа для расчёт
		r.Post("/api/user/orders", h.postOrder)
		//запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа
		r.Post("/api/user/balance/withdraw", h.postWithdraw)

	})

	r.Group(func(r chi.Router) {
		r.Post("/api/user/register", h.postRegister)
		//аутентификация пользователя
		r.Post("/api/user/login", h.postLogin())

		// Ошибочный путь
		r.Post("/*", h.errorPath)
		r.Get("/*", h.errorPath)
	})

}
