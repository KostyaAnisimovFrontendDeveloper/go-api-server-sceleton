-- +goose Up
-- +goose StatementBegin
CREATE TABLE pages
(
    id   uuid default uuid_generate_v4() not null
        primary key,
    name varchar(100)  null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pages
-- +goose StatementEnd
