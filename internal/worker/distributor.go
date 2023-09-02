package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/stg35/avito_test/internal/handler/dto"
)

type TaskDistributor interface {
	DistributeTaskSegmentExpiration(
		ctx context.Context,
		payload *dto.ChangeSegmentDto,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client,
	}
}
