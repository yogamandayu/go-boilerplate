package healthcheck

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type StorageWorkflow struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

type StorageStatus struct {
	StackStatus StackStatus
}

type StackStatus struct {
	Minio MinioStatus
}

type MinioStatus struct {
	Status string
}

func NewStorageWorkflow(db *pgxpool.Pool, redis *redis.Client) *StorageWorkflow {
	return &StorageWorkflow{
		db:    db,
		redis: redis,
	}
}

func (p *StorageWorkflow) Health(ctx context.Context) StorageStatus {
	status := StorageStatus{
		StackStatus: StackStatus{
			Minio: MinioStatus{},
		},
	}

	return status
}
