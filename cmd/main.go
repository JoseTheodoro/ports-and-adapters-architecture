package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	httpapi "app/internal/adapters/http"
	approveorderrepo "app/internal/adapters/postgres/approveorder"
	createorderrepo "app/internal/adapters/postgres/createorder"
	"app/internal/core/application/approveorder"
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

	// repositories implementations
	createOrderRepo := createorderrepo.New(conn)
	approveOrderRepo := approveorderrepo.New(conn)

	// use cases implementations (port)
	createOrderService := createorder.New(createOrderRepo)
	approveOrderService := approveorder.New(approveOrderRepo)

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
