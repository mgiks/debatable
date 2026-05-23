-- +goose Up
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    email citext UNIQUE NOT NULL,
    passhash bytea NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now()
);

-- +goose Down
DROP EXTENSION IF EXISTS citext;

DROP TABLE IF EXISTS users;
