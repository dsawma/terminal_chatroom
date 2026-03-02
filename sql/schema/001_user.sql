-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    username TEXT NOT NULL, 
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL, 
    hashed_password TEXT NOT NULL DEFAULT 'unset'
);

-- +goose Down 
DROP TABLE users; 