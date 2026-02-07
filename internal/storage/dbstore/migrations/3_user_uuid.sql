-- +goose Up
-- +goose StatementBegin
ALTER TABLE "content"."shorturl"
    ADD COLUMN "user_id" uuid DEFAULT '00000000-0000-0000-0000-000000000000',
    ADD COLUMN "created_at" timestamp DEFAULT NOW()
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "content"."shorturl"
    DROP COLUMN "user_id",
    DROP COLUMN "created_at";
-- +goose StatementEnd