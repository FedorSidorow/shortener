-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS content;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA content;
-- +goose StatementEnd
