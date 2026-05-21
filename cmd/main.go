package main

import (
	http2 "app/internal/adapters/http"
	"app/internal/adapters/postgres"
	"app/internal/core/application"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	ctx := context.Background()
	conn, err := pgxpool.New(ctx, "postgres://order:order@localhost/order?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	repository := postgres.NewOrderRepositoryPostgress(conn)
	createOrderInteractor := application.NewCreateOrderHandler(repository)
	h := http2.NewHandleCreateOrder(createOrderInteractor)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/order", h.CreateOrder)

	httpSever := http.Server{
		Addr:    ":3001",
		Handler: mux,
	}

	fmt.Println("HTTP Server Started at http://localhost:3001")
	if err := httpSever.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
