-- +goose Up
CREATE TABLE room (
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL,
    room_name TEXT NOT NULL, 
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

Create TABLE room_member (
    room_id UUID NOT NULL REFERENCES room(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (room_id, user_id)
);

-- +goose Down 
DROP TABLE room;
DROP TABLE room_member; 