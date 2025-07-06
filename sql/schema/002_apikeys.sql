-- +goose Up
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (SHA2(UUID(), 256));

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;