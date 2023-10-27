package routers

import (
	"github.com/Longreader/order-owerview/internal/http/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(h *handlers.Handler) chi.Router {

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Get("/order/{order_uid}", h.GetOrder)

	return r
}
