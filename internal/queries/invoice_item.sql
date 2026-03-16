-- name: CreateInvoiceItem :one
INSERT INTO invoice_item (
    fk_invoice, fk_item, qty
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: GetInvoiceItem :one
SELECT * FROM invoice_item
WHERE id = ? LIMIT 1;

-- name: ListInvoiceItem :many
SELECT * FROM invoice_item;

-- name: UpdateInvoiceItem :one
UPDATE invoice_item
SET fk_invoice = ?,
fk_item = ?,
qty = ?
WHERE id = ?
RETURNING *;

-- name: DeleteInvoiceItem :exec
DELETE FROM invoice_item
WHERE id = ?;
