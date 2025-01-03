package worker

import (
	"context"
	"encoding/json"
	"ticketing/expiration/internal/model"
	"ticketing/expiration/internal/publisher"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

type AsynqWorker struct {
	ExpirationPublisher publisher.ExpirationPublisher
	Logger              *logrus.Logger
}

func NewAsynqWorker(expirationPublisher publisher.ExpirationPublisher, log *logrus.Logger) *AsynqWorker {
	return &AsynqWorker{
		ExpirationPublisher: expirationPublisher,
		Logger:              log,
	}
}

func (w *AsynqWorker) HandleOrderExpirationTask(ctx context.Context, task *asynq.Task) error {
	var payload model.OrderExpirationPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		w.Logger.WithError(err).Error("failed unmarshal order expiration payload")
		return err
	}

	w.Logger.Infof("Processing expiration for OrderID: %v", payload.OrderID)

	// publish the expiration complete event to NATS
	if err := w.ExpirationPublisher.Expired(payload.OrderID); err != nil {
		w.Logger.WithError(err).Error("failed to publish expiration complete event")
		return err
	}

	return nil
}
