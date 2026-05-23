-- +goose Up
ALTER TABLE posts
    ADD COLUMN IF NOT EXISTS user_id bigint REFERENCES users (id);

-- +goose Down
ALTER TABLE posts
    DROP COLUMN IF EXISTS user_id;
