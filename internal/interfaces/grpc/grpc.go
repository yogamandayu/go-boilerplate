package grpc

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/yogamandayu/go-boilerplate/internal/app"
	"github.com/yogamandayu/go-boilerplate/internal/interfaces/grpc/handler/healthcheck"
	healthcheckPB "github.com/yogamandayu/go-boilerplate/internal/interfaces/grpc/protobuf/healthcheck"
	"github.com/yogamandayu/go-boilerplate/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	app *app.App

	Port string
}

// Option is option to rest struct.
type Option func(r *Server)

func NewServer(app *app.App) *Server {
	return &Server{
		app:  app,
		Port: "4080",
	}
}

// With is to set option.
func (s *Server) With(opts ...Option) *Server {
	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", s.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	healthcheckPB.RegisterPingServiceServer(grpcServer, healthcheck.Handler{})

	if util.GetEnv("APP_ENV", "") != "production" {
		reflection.Register(grpcServer)
	}

	log.Println("HTTP gRPC server is starting ...")
	if err = grpcServer.Serve(lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen %s error: %v", s.Port, err)
	}
	log.Println("Shutting down HTTP gRPC server gracefully...")

	return nil
}
