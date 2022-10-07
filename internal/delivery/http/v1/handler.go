package v1

import (
	"github.com/ValikoDorodnov/go_sample/pkg/logger"
	"net/http"

	"github.com/ValikoDorodnov/go_sample/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	greetingService *service.GreetingService
	log             *logger.Logger
}

func NewHandler(greetingService *service.GreetingService, log *logger.Logger) *Handler {
	return &Handler{
		greetingService: greetingService,
		log:             log,
	}
}

func (h *Handler) GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.hello).Methods(http.MethodGet)
	r.HandleFunc("/db-test", h.helloDb).Methods(http.MethodGet)

	return r
}
