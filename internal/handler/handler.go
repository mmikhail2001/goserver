package handler

import (
	"fmt"
	"net/http"

	"github.com/mmikhail2001/goserver/internal/service"
)

type Handler struct {
	service service.Service
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.signUp)
	return mux
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello signUp!")
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
