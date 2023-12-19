-- +goose Up
ALTER TABLE users ADD COLUMN password VARCHAR(64);

-- +goose Down
ALTER TABLE users DROP COLUMN password;