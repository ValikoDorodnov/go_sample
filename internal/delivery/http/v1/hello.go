package v1

import (
	"net/http"

	"github.com/ValikoDorodnov/go_sample/pkg/rest"
)

func (h *Handler) hello(w http.ResponseWriter, _ *http.Request) {
	resp, err := h.greetingService.HelloProcess()
	if err != nil {
		rest.ResponseErrors(w, err)
	}
	rest.ResponseOk(w, resp)
}

func (h *Handler) dbTest(w http.ResponseWriter, _ *http.Request) {
	resp, err := h.greetingService.DbProcess()
	if err != nil {
		rest.ResponseErrors(w, err)
	}
	rest.ResponseOk(w, resp)
}
