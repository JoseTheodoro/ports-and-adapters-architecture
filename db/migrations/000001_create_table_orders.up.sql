CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    order_id UUID NOT NULL,
    price INT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NULL
)