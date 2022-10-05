package service

import (
	"fmt"
	"net/http"
)

type GreetingService struct {
	HelloHandler func(w http.ResponseWriter, r *http.Request)
}

func NewGreetingService() *GreetingService {
	return &GreetingService{
		HelloHandler: HelloHandler,
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
