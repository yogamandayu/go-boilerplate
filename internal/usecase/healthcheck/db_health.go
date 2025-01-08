package healthcheck

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBHealthWorkflow struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

type DBHealthStatus struct {
	StackStatus DBHealthStackStatus
}

type DBHealthStackStatus struct {
	Postgres PostgresStatus
}

type PostgresStatus struct {
	Status        string
	TotalConns    uint32
	IdleConns     uint32
	AcquiredConns uint32
}

func NewDBHealthWorkflow(db *pgxpool.Pool, log *slog.Logger) *DBHealthWorkflow {
	return &DBHealthWorkflow{
		db:  db,
		log: log,
	}
}

func (p *DBHealthWorkflow) Health(ctx context.Context) DBHealthStatus {
	status := DBHealthStatus{
		StackStatus: DBHealthStackStatus{
			Postgres: PostgresStatus{},
		},
	}

	err := p.db.Ping(ctx)
	if err == nil {
		status.StackStatus.Postgres.Status = "OK"
		status.StackStatus.Postgres.TotalConns = uint32(p.db.Stat().TotalConns())
		status.StackStatus.Postgres.IdleConns = uint32(p.db.Stat().IdleConns())
		status.StackStatus.Postgres.AcquiredConns = uint32(p.db.Stat().AcquiredConns())
	}

	return status
}
