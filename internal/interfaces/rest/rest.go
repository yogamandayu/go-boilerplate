package rest

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yogamandayu/go-boilerplate/internal/app"
	"github.com/yogamandayu/go-boilerplate/internal/interfaces/rest/route"
)

// Server is a http rest api struct.
type Server struct {
	app *app.App

	Port    string
	Handler http.Handler
}

// NewServer is a constructor.
func NewServer(app *app.App) *Server {
	return &Server{
		app:  app,
		Port: ":8080",
	}
}

// With is to set option.
func (s *Server) With(opts ...Option) *Server {
	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Run is to run http rest api service.
func (s *Server) Run() error {
	router := route.NewRouter(s.app)
	server := http.Server{
		Addr:         fmt.Sprintf(":%s", s.Port),
		Handler:      router.Handler(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("Shutdown error: %v", err)
		}
		log.Println("Shutdown complete")
	}()

	log.Println("HTTP Server server is starting ...")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen %s error: %v", s.Port, err)
	}
	log.Println("Shutting down HTTP Server server gracefully...")

	return nil
}
