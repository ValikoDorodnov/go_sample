package v1

import (
	"github.com/ValikoDorodnov/go_sample/pkg/rest"
	"net/http"
)

func (h *Handler) helloDb(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp, err := h.greetingService.HelloDbProcess(ctx)
	if err != nil {
		rest.ResponseErrors(w, err)
	}
	rest.ResponseOk(w, resp)
}
