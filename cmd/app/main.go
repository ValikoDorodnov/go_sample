package main

import (
	"context"
	"fmt"
	"github.com/ValikoDorodnov/go_sample/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/ValikoDorodnov/go_sample/internal/repository"
	"github.com/ValikoDorodnov/go_sample/pkg/db"

	"github.com/ValikoDorodnov/go_sample/internal/config"
	"github.com/ValikoDorodnov/go_sample/internal/delivery/http"
	"github.com/ValikoDorodnov/go_sample/internal/delivery/http/v1"
	"github.com/ValikoDorodnov/go_sample/internal/service"
)

func main() {
	conf := config.InitConfig()
	log := logger.NewLogger()

	dbConnection, err := db.Init(conf.Db)
	if err != nil {
		log.Error(fmt.Sprintf("err %v", err))
	}
	defer dbConnection.Close()

	repo := repository.NewGreetingRepository(dbConnection)
	greetService := service.NewGreetingService(repo)
	handler := v1.NewHandler(greetService, log)

	srv := http.NewRestServer(conf.Rest, handler.GetRouter())

	go func() {
		fmt.Printf("rest server started at port %s", conf.Rest.Port)
		if err := srv.Run(); err != nil {
			log.Error(fmt.Sprintf("error occured while running http server: %s", err.Error()))
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Error(fmt.Sprintf("error occured on server shutting down: %s", err.Error()))
	}
}
