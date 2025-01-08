package healthcheck

import (
	"context"
	"net/http"
	"time"

	"github.com/yogamandayu/go-boilerplate/internal/interfaces/rest/response"
	"github.com/yogamandayu/go-boilerplate/internal/usecase/healthcheck"

	_ "github.com/yogamandayu/go-boilerplate/docs"
)

// Ping is healthcheck handler.
// @Summary      Ping
// @Description  Responds with "Pong" and stack status.
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200 {object} ResponseContract
// @Router       /healthcheck [get]
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancelCtx()

	pingWorkflow := healthcheck.NewPingWorkflow()
	status := pingWorkflow.Ping(ctx)
	data := PingResponseContract{
		Message:   status.Message,
		Timestamp: status.Timestamp,
	}

	response.NewHTTPSuccessResponse(data, "Success").WithStatusCode(http.StatusOK).AsJSON(w)
}
