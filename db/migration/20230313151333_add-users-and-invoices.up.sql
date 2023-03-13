BEGIN;

INSERT INTO users (first_name, last_name, balance) VALUES
    ('Bob', 'Leponge', 241817),
    ('Kevin', 'Hart', 49297),
    ('Lynne', 'Franklin', 82540),
    ('Loren', 'Ipsun', 402758),
    ('Billy', 'Joe', 226777),
    ('Joe', 'The Kid', 144970),
    ('Wido', 'Black', 205387);

INSERT INTO invoices (user_id, label, amount) VALUES
    (11, 'Construction', 1876),
    (8, 'Work', 764),
    (11, 'Stuff', 956),
    (16, 'Wedding', 3000),
    (3, 'Construction', 1876),
    (7, 'House', 956),
    (11, 'Construction', 2000);
UPDATE invoices SET status = 'paid' WHERE id = 3;

COMMIT;