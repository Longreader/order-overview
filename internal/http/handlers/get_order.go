package handlers

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "order_uid")

	if len(id) == 0 {
		http.Error(w, "wrong order id", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrder(id)

	if err != nil {
		http.Error(w, "can not get order", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("internal/http/template/index.html")
	if err != nil {
		logrus.Errorf("error occured, %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, order)
}
