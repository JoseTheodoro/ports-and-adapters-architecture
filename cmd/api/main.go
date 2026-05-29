package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	amqp "github.com/rabbitmq/amqp091-go"

	httpapi "app/internal/adapters/http"
	approveorderrepo "app/internal/adapters/postgres/approveorder"
	createorderrepo "app/internal/adapters/postgres/createorder"
	"app/internal/adapters/rabbitmq"
	"app/internal/core/application/approveorder"
	"app/internal/core/application/createorder"
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

	// publisher adapters implementations
	publishOrderApproved := rabbitmq.New(ch)

	// repositories implementations
	createOrderRepo := createorderrepo.New(conn)
	approveOrderRepo := approveorderrepo.New(conn)

	// use cases implementations (port)
	createOrderService := createorder.New(createOrderRepo)
	approveOrderService := approveorder.New(approveOrderRepo, publishOrderApproved)

	// http handlers
	createOrderHandle := httpapi.NewHTTPHandler(createOrderService)
	approveOrderHandle := httpapi.NewApproveOrderHTTPHandler(approveOrderService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/order", createOrderHandle.CreateOrder)
	mux.HandleFunc("GET /api/order/approve/{id}", approveOrderHandle.ApproveOrder)

	httpSever := http.Server{
		Addr:    ":3001",
		Handler: mux,
	}

	slog.Info("HTTP Server Started at http://localhost:3001")
	if err := httpSever.ListenAndServe(); err != nil {
		slog.Error("received error from http server", "err", err)
	}

}
