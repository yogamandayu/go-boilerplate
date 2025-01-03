package route

import (
	"net/http"

	"github.com/yogamandayu/go-boilerplate/internal/app"
	"github.com/yogamandayu/go-boilerplate/internal/interfaces/rest/handler/healthcheck"
)

// HealthRoute is a health route to monitor service health.
func HealthRoute(mux *http.ServeMux, app *app.App) {
	pingHandler := healthcheck.NewHandler(app)
	mux.HandleFunc("/ping", pingHandler.Ping)
}
