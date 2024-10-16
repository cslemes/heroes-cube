-- name: GetCharacter :one
SELECT *
FROM characters
WHERE name = ?
LIMIT 1;
-- name: GetInventory :many
SELECT *
FROM inventory
WHERE character_id = ?;
-- name: GetSlots :many
SELECT *
FROM slots
WHERE character_id = ?;
-- name: CreateCharacter :execresult
INSERT INTO characters (name, race, class, damage, level)
VALUES (?, ?, ?, ?, ?);
-- name: AddInventoryItem :execresult
INSERT INTO inventory (
        character_id,
        name,
        damage,
        price,
        class,
        quantity
    )
VALUES (?, ?, ?, ?, ?, ?);
-- name: UpdateInventoryItemQuantity :exec
UPDATE inventory
SET quantity = quantity + ?
WHERE id = ?;
-- name: EquipItem :exec
INSERT INTO slots (character_id, slot, name, damage, price, class)
VALUES (?, ?, ?, ?, ?, ?) ON CONFLICT (character_id, slot) DO
UPDATE
SET name = excluded.name,
    damage = excluded.damage,
    price = excluded.price,
    class = excluded.class;
-- name: UnequipItem :exec
DELETE FROM slots
WHERE character_id = ?
    AND slot = ?;