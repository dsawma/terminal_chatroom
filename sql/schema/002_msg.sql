-- +goose Up
CREATE TABLE message (
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL, 
    body TEXT NOT NULL,
    room_id UUID NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down 
DROP TABLE message; 