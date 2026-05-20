-- name: CreateOrder :one
INSERT INTO orders (order_id, price, status) VALUES ($1, $2, $3) RETURNING *;