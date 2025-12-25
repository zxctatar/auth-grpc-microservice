package grpcserv

import (
	"auth/internal/transport/grpc/handler"
	authv1 "auth/internal/transport/grpc/pb"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	log  *slog.Logger
	serv *grpc.Server
}

func NewServer(log *slog.Logger, authHand *handler.AuthHandler) *Server {
	grpc := grpc.NewServer()
	authv1.RegisterAuthServiceServer(grpc, authHand)

	return &Server{
		log:  log,
		serv: grpc,
	}
}

func (s *Server) MustLoad(port int) {
	const op = "grpcserv.MustLoad"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic("failed to start listening")
	}

	defer l.Close()

	s.log.Info("starting the server", slog.String("op", op), slog.Int("port", port))

	if err := s.serv.Serve(l); err != nil {
		panic("error with server")
	}
}

func (s *Server) Stop() {
	const op = "grpcserv.Stop"

	s.log.Info("stopping the server...", slog.String("op", op))
	s.serv.GracefulStop()
	s.log.Info("the server is stopped", slog.String("op", op))
}
