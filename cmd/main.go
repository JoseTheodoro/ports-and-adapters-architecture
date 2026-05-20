package main

import (
	http2 "app/internal/adapters/inbound/http"
	"app/internal/adapters/outbound/postgres"
	"app/internal/core/application"
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	ctx := context.Background()
	conn, _ := pgxpool.New(ctx, "postgres://order:order@localhost/order?sslmode=disable")

	repository := postgres.NewRepositoryOrderPostgress(conn)
	usecase := application.NewCreateOrderUseCase(repository)
	h := http2.NewHandleCreateOrder(usecase)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/order", h.CreateOrder)

	httpSever := http.Server{
		Addr:    ":3001",
		Handler: mux,
	}

	fmt.Println("HTTP Server Started at http://localhost:3001")
	httpSever.ListenAndServe()

}
