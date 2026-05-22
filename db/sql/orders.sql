-- name: CreateOrder :one
INSERT INTO orders (order_id, price, status) VALUES ($1, $2, $3) RETURNING *;
-- name: FindByOrderID :one
SELECT id, order_id, price, status, created_at, updated_at FROM orders where order_id = $1;
-- name: ApproveOrder :exec
UPDATE orders SET status = $1, updated_at = now() WHERE order_id = $2;