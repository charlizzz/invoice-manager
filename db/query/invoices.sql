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
  user_id, status, label, amount
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateInvoiceAmount :one
UPDATE invoices 
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: UpdateInvoiceLabel :one
UPDATE invoices 
SET label = $2
WHERE id = $1
RETURNING *;

-- name: DeleteInvoice :exec
DELETE FROM invoices 
WHERE id = $1;