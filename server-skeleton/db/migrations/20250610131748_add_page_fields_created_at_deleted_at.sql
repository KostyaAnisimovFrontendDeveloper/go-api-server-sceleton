-- +goose Up
-- +goose StatementBegin
ALTER TABLE pages
    ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN updated_at TIMESTAMP  NULL DEFAULT NULL,
    ADD COLUMN deleted_at TIMESTAMP  NULL DEFAULT NULL
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE pages
    DROP COLUMN created_at,
    DROP COLUMN updated_at,
    DROP COLUMN deleted_at
-- +goose StatementEnd
