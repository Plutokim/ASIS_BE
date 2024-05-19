-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_product (
    id UUID NOT NULL PRIMARY KEY,
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id UUID NOT NULL REFERENCES product(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_product;
-- +goose StatementEnd
