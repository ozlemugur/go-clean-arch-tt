package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/ozlemugur/go-clean-arch-tt/config"
	v1 "github.com/ozlemugur/go-clean-arch-tt/internal/controller/http/v1"
	"github.com/ozlemugur/go-clean-arch-tt/internal/controller/scheduler"
	"github.com/ozlemugur/go-clean-arch-tt/internal/usecase"
	"github.com/ozlemugur/go-clean-arch-tt/internal/usecase/repo"
	"github.com/ozlemugur/go-clean-arch-tt/internal/usecase/webapi"
	"github.com/ozlemugur/go-clean-arch-tt/pkg/httpserver"
	"github.com/ozlemugur/go-clean-arch-tt/pkg/logger"
	"github.com/ozlemugur/go-clean-arch-tt/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	messageSenderWebAPI := webapi.NewMessageSenderWebAPI(cfg.Mocky.URL)
	// Use case
	messageUseCase := usecase.NewMessageUseCase(
		repo.NewMessage(pg),
		messageSenderWebAPI,
		l,
	)

	l.Debug("messageUseCase created")
	autoMessageScheduler := scheduler.NewAutoMessageScheduler(messageUseCase, l)
	autoMessageScheduler.Start()

	autoMessageUseCase := usecase.NewAutoMessageSchedulerUseCase(messageUseCase, autoMessageScheduler, l)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, autoMessageUseCase, messageUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
