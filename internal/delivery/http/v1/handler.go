package v1

import (
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

func (h Handler) Init() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", h.greetingService.HelloHandler)
	return r
}
