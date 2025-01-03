package healthcheck

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

type CacheHealthWorkflow struct {
	log   *slog.Logger
	redis *redis.Client
}

type CacheHealthStatus struct {
	StackStatus CacheHealthStackStatus
}

type CacheHealthStackStatus struct {
	Redis RedisStatus
}

type RedisStatus struct {
	Status     string
	TotalConns uint32
	IdleConns  uint32
	StaleConns uint32
}

func NewCacheHealthWorkflow(redis *redis.Client, log *slog.Logger) *CacheHealthWorkflow {
	return &CacheHealthWorkflow{
		redis: redis,
		log:   log,
	}
}

func (p *CacheHealthWorkflow) Health(ctx context.Context) CacheHealthStatus {
	status := CacheHealthStatus{
		StackStatus: CacheHealthStackStatus{
			Redis: RedisStatus{},
		},
	}

	redisStatus := p.redis.Ping(ctx)
	if redisStatus.Err() != nil {
		p.log.Error(redisStatus.Err().Error())
		status.StackStatus.Redis.Status = "ERROR"
		return status
	}

	status.StackStatus.Redis = RedisStatus{
		Status:     "OK",
		TotalConns: p.redis.PoolStats().TotalConns,
		IdleConns:  p.redis.PoolStats().IdleConns,
		StaleConns: p.redis.PoolStats().StaleConns,
	}

	return status
}
