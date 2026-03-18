-- +goose Up
CREATE TABLE room (
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL,
    room_name TEXT NOT NULL
);

-- +goose Down 
DROP TABLE room;