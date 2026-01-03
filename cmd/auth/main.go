package main

import (
	"auth/internal/config"
	jwtservice "auth/internal/infrastructure/jwt"
	"auth/internal/infrastructure/postgres"
	grpcserv "auth/internal/transport/grpc"
	"auth/internal/transport/grpc/handler"
	"auth/internal/usecase/implementations/login"
	"auth/internal/usecase/implementations/registration"
	"auth/pkg/logger"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Logger.Level)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
		cfg.Postgres.SslMode,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(fmt.Sprintf("cannot open sql: %s", err.Error()))
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("cannot ping db: %s", err.Error()))
	}

	postg := postgres.NewPostgres(log, db)
	tokService := jwtservice.NewJWTService(cfg.GRPC.JWTSecretKey, &cfg.GRPC.JWTTimeOut)

	regUC := registration.NewRegistrationUC(log, postg)
	loginUC := login.NewLoginUc(log, postg, tokService)

	authHandl := handler.NewAuthHandler(log, &cfg.GRPC.TimeOut, regUC, loginUC)

	server := grpcserv.NewServer(log, authHandl)

	go server.MustLoad(cfg.GRPC.Port)

	sysChan := make(chan os.Signal, 1)
	signal.Notify(sysChan, syscall.SIGINT, syscall.SIGTERM)

	<-sysChan

	server.Stop()
}
