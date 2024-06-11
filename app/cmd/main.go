package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mikeyQwn/server-ping/config"
	"github.com/mikeyQwn/server-ping/internal/delivery"
	"github.com/mikeyQwn/server-ping/internal/usecase"
	"github.com/mikeyQwn/server-ping/pkg/logger"
)

func main() {
	log := logger.New(&logger.LoggerConfig{})

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err, "could not load the config")
	}

	uc := usecase.New()

	s := delivery.New(log, uc, cfg.PingAddress)
	s.MapHandlers()

	go func(s *delivery.Server) {
		addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
		log.Infof("Server is running on %s", addr)
		log.Fatal(s.Run(addr), "could not run the server")
	}(s)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
}
