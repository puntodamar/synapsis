-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id           VARCHAR(255) PRIMARY KEY,
    customer_id  VARCHAR(255) NOT NULL,
    status       VARCHAR(32)  NOT NULL,
    created_at   TIMESTAMPTZ  DEFAULT NOW()
);

CREATE TABLE order_items (
     id        BIGSERIAL PRIMARY KEY,
     order_id  VARCHAR(255) NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
     sku       VARCHAR(255) NOT NULL,
     qty       INTEGER      NOT NULL
);

CREATE UNIQUE INDEX order_items_uniq_order_sku
    ON order_items (order_id, sku);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
