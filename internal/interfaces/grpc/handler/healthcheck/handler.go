// Package healthcheck is a package to all health check related gRPC API handler request.
package healthcheck

import (
	"context"
	"time"

	healthcheckPB "github.com/yogamandayu/go-boilerplate/internal/interfaces/grpc/protobuf/healthcheck"
	"github.com/yogamandayu/go-boilerplate/internal/usecase/healthcheck"
)

type Handler struct {
	healthcheckPB.UnimplementedPingServiceServer
}

var _ healthcheckPB.PingServiceServer = &Handler{}

func (h Handler) Ping(ctx context.Context, request *healthcheckPB.PingRequest) (*healthcheckPB.PingResponse, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, 5*time.Second)
	defer cancelCtx()

	pingWorkflow := healthcheck.NewPingWorkflow()
	status := pingWorkflow.Ping(ctx)
	return &healthcheckPB.PingResponse{
		Timestamp: status.Timestamp,
		Message:   status.Message,
	}, nil
}
