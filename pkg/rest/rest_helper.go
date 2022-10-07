package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ValikoDorodnov/go_sample/pkg/logger"
	"net/http"
)

func ResponseOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func ResponseErrors(w http.ResponseWriter, err error, l *logger.Logger) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	l.Error(fmt.Sprintf("err %v", err))
	json.NewEncoder(w).Encode(ErrorResponse{Errors: []string{err.Error()}})
}
