package v1

import (
	"net/http"

	"github.com/ValikoDorodnov/go_sample/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	greetingService *service.GreetingService
}

func NewHandler(greetingService *service.GreetingService) *Handler {
	return &Handler{
		greetingService: greetingService,
	}
}

func (h *Handler) GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.hello).Methods(http.MethodGet)
	r.HandleFunc("/db-test", h.dbTest).Methods(http.MethodGet)

	return r
}
