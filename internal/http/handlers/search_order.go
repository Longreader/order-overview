package handlers

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) SearchOrder(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		logrus.Errorf("error while parsing form, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	order_uid := r.Form.Get("order_uid")
	fmt.Printf("%v", order_uid)
}
