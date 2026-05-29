package rabbitmq

import (
	"app/internal/core/domain"
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderApprovedPublisher struct {
	rabbitch *amqp.Channel
}

func New(ch *amqp.Channel) *OrderApprovedPublisher {
	return &OrderApprovedPublisher{
		rabbitch: ch,
	}
}

func (o *OrderApprovedPublisher) PublishOrderApproved(ctx context.Context, event *domain.OrderApprovedEvent) error {
	fmt.Printf("dispatching OrderApproved event: %v\n", event)

	messageBody := toJSON(event)

	err := o.rabbitch.PublishWithContext(ctx,
		"orders", // exchange
		"",       // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: amqp.MimeApplicationJSON,
			Body:        messageBody,
		},
	)

	if err != nil {
		return fmt.Errorf("publish order approved event error > %w", err)
	}

	return nil
}

func toJSON(event *domain.OrderApprovedEvent) []byte {
	m := new(bytes.Buffer{})
	json.NewEncoder(m).Encode(event)
	return m.Bytes()
}
