-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payment (
    id UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    card_token VARCHAR(255) NOT NULL,
    card_number VARCHAR(20) NOT NULL,
    expiry_date VARCHAR(7) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment;
-- +goose StatementEnd
