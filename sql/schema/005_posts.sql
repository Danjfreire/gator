-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    title TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    published_at TIMESTAMPTZ NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id)
);

-- +goose Down
DROP TABLE posts;