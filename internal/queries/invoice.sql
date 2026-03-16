-- name: CreateInvoice :one
INSERT INTO invoice (
    invoice_number,
    date,
    fk_from_company,
    fk_to_company,
    fk_payment_detail,
    igst,
    total,
    fk_shipping_address
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetInvoice :one
SELECT * FROM invoice
WHERE id = ? LIMIT 1;

-- name: ListInvoice :many
SELECT * FROM invoice;

-- name: UpdateInvoice :one
UPDATE invoice
SET invoice_number = ?,
date = ?,
updated_at = ?,
fk_from_company = ?,
fk_to_company = ?,
fk_payment_detail = ?,
igst = ?,
total = ?,
fk_shipping_address = ?
WHERE id = ?
RETURNING *;

-- name: DeleteInvoice :exec
DELETE from invoice
WHERE id = ?;

-- name: GetLatestInvoiceNumber :one
SELECT invoice_number FROM invoice
ORDER BY invoice_number DESC
LIMIT 1;
