package healthcheck

import (
	"context"
	"time"

	healthcheckPB "github.com/yogamandayu/go-boilerplate/internal/interfaces/grpc/protobuf/healthcheck"
	"github.com/yogamandayu/go-boilerplate/internal/workflow/healthcheck"
)

type Handler struct {
	healthcheckPB.UnimplementedPingServiceServer
}

var _ healthcheckPB.PingServiceServer = &Handler{}

func (h Handler) Ping(ctx context.Context, request *healthcheckPB.PingRequest) (*healthcheckPB.PingResponse, error) {
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	pingWorkflow := healthcheck.NewPingWorkflow()
	status := pingWorkflow.Ping(ctx)
	return &healthcheckPB.PingResponse{
		Timestamp: status.Timestamp,
		Message:   status.Message,
	}, nil
}
