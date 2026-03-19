-- name: CreateCompany :one
INSERT INTO company (
    name, gst, address, signature_image_path
) VALUES (
    ?, ?, ?, ?
)
RETURNING *;

-- name: GetCompany :one
SELECT * FROM company
WHERE id = ? LIMIT 1;

-- name: ListCompany :many
SELECT * FROM company;

-- name: UpdateCompany :one
UPDATE company
SET name = ?,
gst = ?,
address = ?,
signature_image_path = ?
WHERE id = ?
RETURNING *;

-- name: DeleteCompany :exec
DELETE FROM company
WHERE id = ?;
