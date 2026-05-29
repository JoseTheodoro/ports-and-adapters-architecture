package rabbitmq

import (
	"context"
	"fmt"
	"log/slog"

	amqp "github.com/rabbitmq/amqp091-go"

	"app/internal/core/application/processorder"
)

type ConsumeApproveOrder struct {
	p  processorder.Processor
	ch *amqp.Channel
}

func NewConsumeApproverOrder(p processorder.Processor, ch *amqp.Channel) *ConsumeApproveOrder {
	return &ConsumeApproveOrder{
		p:  p,
		ch: ch,
	}
}

func (c *ConsumeApproveOrder) Consume(ctx context.Context) error {
	msgs, err := c.ch.Consume("orders_approved", "", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("consume approved orders error > %w", err)
	}

	quitch := make(chan struct{})

	go func() {
		for d := range msgs {

			message := NewOrderApprovedMessage(d.Body)
			input := message.ToInput()
			if err := c.p.Process(ctx, input); err != nil {
				slog.Error("order process error", "error", err)
				continue
			}
			slog.Info("order has been processed", "order", message)
		}
	}()
	slog.Info("waiting messages from orders_approved queue...")
	<-quitch
	return nil
}
