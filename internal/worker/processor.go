package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/stg35/avito_test/internal/service"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSegmentExpiration(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server  *asynq.Server
	service *service.Service
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, service *service.Service) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{},
	)

	return &RedisTaskProcessor{
		server,
		service,
	}
}
