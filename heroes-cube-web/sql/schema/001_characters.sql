-- +goose Up
CREATE TABLE characters (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    race TEXT NOT NULL,
    class TEXT NOT NULL,
    damage INTEGER NOT NULL,
    level INTEGER NOT NULL
);
-- +goose Down
DROP TABLE characters;