package main

import (
	"auth/internal/config"
	grpcserv "auth/internal/transport/grpc"
	"auth/internal/transport/grpc/handler"
	"auth/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Logger.Level)

	authHandl := handler.NewAuthHandler(log)

	server := grpcserv.NewServer(log, authHandl)

	go server.MustLoad(cfg.GRPC.Port)

	sysChan := make(chan os.Signal, 1)
	signal.Notify(sysChan, syscall.SIGINT, syscall.SIGTERM)

	<-sysChan

	server.Stop()
}
