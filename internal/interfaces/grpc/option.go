package grpc

import (
	"github.com/yogamandayu/go-boilerplate/internal/app"
	"github.com/yogamandayu/go-boilerplate/internal/config"
)

// SetByConfig to set REST API with config.
func SetByConfig(config *config.Config) Option {
	return func(r *Server) {
		r.Port = config.REST.Port
	}
}

// WithApp to set app.
func WithApp(app *app.App) Option {
	return func(r *Server) {
		r.app = app
	}
}
