package main

import (
	"github.com/ValikoDorodnov/go_sample/internal/app"
	"github.com/ValikoDorodnov/go_sample/internal/delivery/http/v1"
	"github.com/ValikoDorodnov/go_sample/internal/service"
	"net/http"
)

func main() {
	s := service.NewGreetingService()
	h := v1.NewHandler(s)
	a := app.NewApp(h)
	r := a.Handler.Init()

	http.Handle("/test", r)
}
