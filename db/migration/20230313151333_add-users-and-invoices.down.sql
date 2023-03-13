BEGIN;

DELETE FROM users WHERE id IN (18, 19, 20, 21, 22, 23, 24);

DELETE FROM invoices WHERE id IN (1, 2, 3, 4, 5, 6, 7);

COMMIT;