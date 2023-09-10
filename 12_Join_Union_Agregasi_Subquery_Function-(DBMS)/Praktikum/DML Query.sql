-- Nomor 1A
INSERT INTO operators (name) VALUES
	('Andi'),
    ('Budi'),
    ('Putra'),
    ('Agung'),
    ('Agus');
    
-- Nomor 1B
INSERT INTO product_types (name) VALUES
	('Alat Tulis'),
    ('Elektronik'),
    ('Perhiasan');
    
-- Nomor 1C
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES
	(1, 3, 'P001', 'Pensil', 1),
    (1, 3, 'P002', 'Pulpen', 1);
    
-- Nomor 1D
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES
    (2, 1, 'K001', 'Komputer', 1),
    (2, 1, 'T001', 'Televisi', 1),
    (2, 1, 'L002', 'Lampu', 1);
    
-- Nomor 1E
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES
    (3, 4, 'C001', 'Cincin', 1),
	(3, 4, 'G001', 'Gelang', 1),
	(3, 4, 'K002', 'Kalung', 1);

-- Nomor 1F
INSERT INTO product_descriptions (product_id, description) VALUES
	(1, 'Pensil mekanik yang di dalamnya terdapat mekanisme dimana bila di tekan akan mengeluarkan grafit kecil melalui lubang kecil yang menonjol di bawahnya. Seperti pensil pada umumnya, juga terdapat karet penghapus di ujung lainnya'),
    (2, 'Alat tulis berupa mata pena berujung tajam yang dilengkapi pegangan berisi kantong tinta yang bisa diisi kembali'),
    (3, 'Alat elektronik yang mengubungkan komponen satu dengan yang lainnya sehingga menghasilkan informasi yang sebelumnya telah diolah terlebih dahulu'),
    (4, 'Salah satu media publik yang memiliki tiga fungsi sebagai alat komunikasi massa'),
    (5, 'Peranti yang menghasilkan cahaya'),
    (6, 'Perhiasan yang melingkar di jari'),
    (7, 'Perhiasan melingkar yang diselipkan atau dikaitkan pada pergelangan tangan seseorang'),
    (8, 'Benda atau barang berupa perhiasan yang dikenakan melingkari leher');
    
-- Nomor 1G
INSERT INTO payment_methods (name, status) VALUES
	('Cash', 1),
    ('Transfer', 1),
    ('Kartu Kredit', 1);
    
-- Nomor 1H
INSERT INTO users (name, status, dob, gender) VALUES
	('Davin', 1, '2003-09-23', 'M'),
    ('Kevin', 1, '2004-10-03', 'M'),
    ('Carlos', 1, '2002-03-14', 'M'),
    ('Jessica', 1, '2000-05-30', 'P'),
    ('Angel', 1, '1999-12-21', 'P');
    
-- Nomor 1I
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES
    (1, 1, 'Paid', 28, 292000),
    (1, 2, 'Paid', 4, 16000000),
    (1, 3, 'Paid', 12, 2611000),
    (2, 1, 'Not Paid', 15, 9220000),
    (2, 2, 'Paid', 6, 12020000),
    (2, 3, 'Partial', 14, 10600000),
    (3, 3, 'Paid', 16, 8135000),
    (3, 2, 'Not Paid', 20, 8660000),
    (3, 1, 'Paid', 19, 18060000),
    (4, 2, 'Paid', 17, 12075000),
    (4, 1, 'Not Paid', 30, 770000),
    (4, 3, 'Paid', 4, 8060000),
    (5, 3, 'Paid', 12, 6620000),
    (5, 2, 'Partial', 21, 4540000),
    (5, 1, 'Paid', 19, 4696000);

-- Nomor 1J
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price) VALUES
    (1, 1, 'Paid', 6, 72000),
    (1, 2, 'Paid', 20, 100000),
    (1, 5, 'Paid', 2, 120000),
    (2, 3, 'Paid', 1, 8000000),
    (2, 4, 'Paid', 1, 4000000),
    (2, 6, 'Paid', 2, 4000000),
    (3, 1, 'Paid', 8, 96000),
    (3, 2, 'Paid', 3, 15000),
    (3, 8, 'Paid', 1, 2500000),
    (4, 5, 'Not Paid', 12, 720000),
    (4, 7, 'Not Paid', 2, 6000000),
    (4, 8, 'Not Paid', 1, 2500000),
    (5, 2, 'Paid', 4, 20000),
    (5, 3, 'Paid', 1, 8000000),
    (5, 4, 'Paid', 1, 4000000),
    (6, 5, 'Paid', 10, 600000),
    (6, 6, 'Not Paid', 2, 4000000),
    (6, 7, 'Not Paid', 2, 6000000),
    (7, 1, 'Paid', 5, 60000),
    (7, 2, 'Paid', 15, 75000),
    (7, 3, 'Paid', 1, 8000000),
    (8, 1, 'Not Paid', 10, 120000),
    (8, 3, 'Not Paid', 1, 8000000),
    (8, 5, 'Not Paid', 9, 540000),
    (9, 2, 'Paid', 12, 60000),
    (9, 4, 'Paid', 2, 8000000),
    (9, 6, 'Paid', 5, 10000000),
    (10, 2, 'Paid', 15, 75000),
    (10, 3, 'Paid', 1, 8000000),
    (10, 4, 'Paid', 1, 4000000),
    (11, 1, 'Not Paid', 10, 120000),
    (11, 2, 'Not Paid', 10, 50000),
    (11, 5, 'Not Paid', 10, 600000),
    (12, 4, 'Paid', 1, 4000000),
    (12, 5, 'Paid', 1, 60000),
    (12, 6, 'Paid', 2, 4000000),
    (13, 1, 'Paid', 10, 120000),
    (13, 4, 'Paid', 1, 4000000),
    (13, 7, 'Paid', 1, 2500000),
    (14, 2, 'Paid', 12, 60000),
    (14, 4, 'Paid', 1, 4000000),
    (14, 5, 'Not Paid', 8, 480000),
    (15, 1, 'Paid', 8, 96000),
    (15, 4, 'Paid', 1, 4000000),
    (15, 5, 'Paid', 10, 600000);

-- Nomor 2A
SELECT * FROM users WHERE gender = 'M';
-- Nomor 2B
SELECT * FROM products WHERE id = 3;
-- Nomor 2C
SELECT * FROM users WHERE created_at between date_sub(now(),INTERVAL 1 WEEK) and now() AND name LIKE '%a%';
-- Nomor 2D
SELECT COUNT(gender) pelangganPerempuan FROM users WHERE gender = 'P';
-- Nomor 2E
SELECT * FROM users ORDER BY name;
-- Nomor 2F
SELECT * FROM products LIMIT 5;

-- Nomor 3A
UPDATE products SET name = 'product dummy' WHERE id = 1;
-- Nomor 3B
UPDATE transaction_details SET qty = 3 WHERE product_id = 1;

-- Nomor 4A
DELIMITER $$
CREATE TRIGGER delete_all_data_product BEFORE DELETE ON products FOR EACH ROW
BEGIN
DELETE FROM transaction_details WHERE product_id = OLD.id;
DELETE FROM product_descriptions WHERE product_id = OLD.id;
END$$
DELIMITER ;
DELETE FROM products WHERE id = 1;

-- Nomor 4B
DELETE FROM products WHERE product_type_id = 1;



-- JOIN, UNION, Subquery, Function
-- Nomor 1
SELECT * FROM transactions WHERE user_id = 1 UNION SELECT * FROM transactions WHERE user_id = 2;

-- Nomor 2
SELECT SUM(total_price) 'Total Price' FROM transactions WHERE user_id = 1;

-- Nomor 3
SELECT COUNT(product_id) FROM transaction_details WHERE product_id IN (SELECT id FROM products WHERE product_type_id = 2);

-- Nomor 4
SELECT p.*, pt.name productType FROM products p INNER JOIN product_types pt ON p.product_type_id = pt.id;

-- Nomor 5
SELECT t.*, u.name user_name, p.name
FROM transactions t LEFT JOIN users u ON t.user_id = u.id
LEFT JOIN transaction_details td ON td.transaction_id = t.id
LEFT JOIN products p ON td.product_id = p.id;

-- Nomor 6
DELIMITER $$
CREATE TRIGGER delete_transaction_details BEFORE DELETE ON transactions FOR EACH ROW
BEGIN
DELETE FROM transaction_details WHERE transaction_id = OLD.id;
END$$
DELIMITER ;

-- Nomor 7
DELIMITER $$
CREATE TRIGGER update_total_qty BEFORE DELETE ON transaction_details FOR EACH ROW
BEGIN
UPDATE transactions SET total_qty = (SELECT SUM(qty) FROM transaction_details WHERE transaction_id = OLD.transaction_id)-OLD.qty WHERE id = OLD.transaction_id;
END$$
DELIMITER ;

-- Nomor 8
SELECT * FROM products WHERE id NOT IN (SELECT product_id FROM transaction_details GROUP BY product_id);
