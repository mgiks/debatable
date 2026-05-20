-- +goose Up
CREATE TABLE IF NOT EXISTS posts (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title text NOT NULL,
    body text NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS posts;
