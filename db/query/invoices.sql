-- name: GetInvoice :one
SELECT * FROM invoices
WHERE id = $1 LIMIT 1;

-- name: ListInvoices :many
SELECT * FROM invoices
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateInvoice :one
INSERT INTO invoices (
  user_id, label, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateInvoiceStatus :one
UPDATE invoices 
SET status = 'paid'
WHERE id = $1
RETURNING *;

-- name: DeleteInvoice :exec
DELETE FROM invoices 
WHERE id = $1;