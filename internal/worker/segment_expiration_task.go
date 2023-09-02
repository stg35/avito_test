package worker

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/stg35/avito_test/internal/handler/dto"
)

const TaskSegmentExpiration = "task:segment_expiration"

func (r *RedisTaskDistributor) DistributeTaskSegmentExpiration(
	ctx context.Context,
	payload *dto.ChangeSegmentDto,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	task := asynq.NewTask(TaskSegmentExpiration, jsonPayload, opts...)
	_, err = r.client.EnqueueContext(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisTaskProcessor) ProcessTaskSegmentExpiration(ctx context.Context, task *asynq.Task) error {
	var payload dto.ChangeSegmentDto
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	err := r.service.User.DeleteSegments(payload)
	if err != nil {
		return err
	}

	return nil

}

func (r *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSegmentExpiration, r.ProcessTaskSegmentExpiration)

	return r.server.Start(mux)
}
