package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	amqp "github.com/rabbitmq/amqp091-go"

	repoprocessorder "app/internal/adapters/postgres/processorder"
	"app/internal/adapters/rabbitmq"
	"app/internal/core/application/processorder"
)

func main() {

	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	// connecting to postgres
	conn, err := pgxpool.New(ctx, "postgres://order:order@localhost/order?sslmode=disable")
	if err != nil {
		slog.Error("error on connect to database", "err", err)
	}
	defer conn.Close()

	// connecting to rabbitmq
	rabbitConn, err := amqp.Dial("amqp://user:root@localhost:5672/")
	if err != nil {
		slog.Error("rabbitmq connect error", "err", err)
	}
	// connection to the channel
	ch, err := rabbitConn.Channel()
	if err != nil {
		slog.Error("channel rabbitmq connect error", "err", err)
	}

	// start repository process orders
	repoProcessOrder := repoprocessorder.New(conn)
	// start processor orders
	processOrderService := processorder.New(repoProcessOrder)

	// start consume or handle messages from rabbitmq
	consumer := rabbitmq.NewConsumeApproverOrder(processOrderService, ch)

	if err := consumer.Consume(ctx); err != nil {
		slog.Error("consume start error", "error", err)
	}

}
