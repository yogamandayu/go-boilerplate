package healthcheck

import (
	"context"
	"time"
)

type PingWorkflow struct {
}

type PingStatus struct {
	Message   string
	Timestamp string
}

func NewPingWorkflow() *PingWorkflow {
	return &PingWorkflow{}
}

func (p *PingWorkflow) Ping(ctx context.Context) PingStatus {
	status := PingStatus{
		Message:   "Pong!",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return status
}
