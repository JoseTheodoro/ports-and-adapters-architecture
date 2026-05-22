package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	httpapi "app/internal/adapters/http"
	"app/internal/adapters/postgres"
	"app/internal/core/application/createorder"
)

func main() {

	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	conn, err := pgxpool.New(ctx, "postgres://order:order@localhost/order?sslmode=disable")
	if err != nil {
		slog.Error("error on connect to database", "err", err)
	}
	defer conn.Close()

	repository := postgres.NewOrderRepositoryPostgress(conn)
	createOrderHandler := createorder.NewCreateOrderHandler(repository)
	h := httpapi.NewHandleCreateOrder(createOrderHandler)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/order", h.CreateOrder)

	httpSever := http.Server{
		Addr:    ":3001",
		Handler: mux,
	}

	slog.Info("HTTP Server Started at http://localhost:3001")
	if err := httpSever.ListenAndServe(); err != nil {
		slog.Error("received error from http server", "err", err)
	}

}
