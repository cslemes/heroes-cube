-- +goose Up
CREATE TABLE inventory (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    character_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    damage INTEGER NOT NULL,
    price INTEGER NOT NULL,
    class TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    FOREIGN KEY (character_id) REFERENCES characters(id)
);
-- +goose Down
DROP TABLE inventory;