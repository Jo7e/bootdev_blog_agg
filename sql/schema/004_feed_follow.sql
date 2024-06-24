-- +goose Up
CREATE TABLE feed_follows (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    feed_id uuid NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE feed_follows;
