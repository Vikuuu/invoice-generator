-- name: CreateItem :one
INSERT INTO item (
    name, hsn, price, gst
) VALUES (
    ?, ?, ?, ?
)
RETURNING *;

-- name: GetItem :one
SELECT * FROM item
WHERE id = ? LIMIT 1;

-- name: ListItem :many
SELECT * FROM item;

-- name: UpdateItem :one
UPDATE item
SET name = ?,
hsn = ?,
price = ?,
gst = ?
WHERE id = ?
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM item
where id = ?;
