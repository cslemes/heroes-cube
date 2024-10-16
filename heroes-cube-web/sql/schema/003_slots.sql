-- +goose Up
CREATE TABLE slots (
    character_id INTEGER NOT NULL,
    slot TEXT NOT NULL,
    name TEXT NOT NULL,
    damage INTEGER NOT NULL,
    price INTEGER NOT NULL,
    class TEXT NOT NULL,
    FOREIGN KEY (character_id) REFERENCES characters(id),
    PRIMARY KEY (character_id, slot)
);
-- +goose Down
DROP TABLE slots;