package usecase

import (
	"context"
	"encoding/json"
	"time"

	"ticketing/expiration/internal/domain"
	"ticketing/expiration/internal/model"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

type ExpirationUsecase interface {
	ScheduleExpiration(ctx context.Context, request *model.OrderExpirationPayload) error
}

type ExpirationUsecaseImpl struct {
	TS     *asynq.Client
	Logger *logrus.Logger
}

func NewExpirationUsecase(ts *asynq.Client, log *logrus.Logger) ExpirationUsecase {
	return &ExpirationUsecaseImpl{
		TS:     ts,
		Logger: log,
	}
}

func (uc *ExpirationUsecaseImpl) ScheduleExpiration(ctx context.Context, request *model.OrderExpirationPayload) error {
	delay := time.Duration(request.Delay) * time.Millisecond

	if delay < 0 {
		delay = 0
	}

	payload := new(model.OrderExpirationPayload)
	payload.OrderID = request.OrderID

	// Create a new Asynq task withg the expiration payload
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Enqueue the task to Asynq with a delay
	task := asynq.NewTask(domain.OrderExpirationTask, data)
	_, err = uc.TS.Enqueue(task, asynq.ProcessIn(delay))
	if err != nil {
		uc.Logger.WithError(err).Error("failed schedule order expiration")
		return err
	}

	uc.Logger.Infof("Scheduled expiration for OrderID: %v", request.OrderID)

	return nil
}
