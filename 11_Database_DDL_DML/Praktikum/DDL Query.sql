-- DDL
-- Nomor 1
CREATE DATABASE alta_online_shop;
USE alta_online_shop;

-- Nomor 2A
CREATE TABLE user(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255),
    dob DATE,
    status_user SMALLINT,
    gender CHAR(1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Nomor 2B
CREATE TABLE product_type(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE operators(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE product(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    operator_id INT,
    product_type_id INT,
    code VARCHAR(50),
    name VARCHAR(255),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_product_type FOREIGN KEY (product_type_id) REFERENCES product_type(id),
    CONSTRAINT fk_product_operator FOREIGN KEY (operator_id) REFERENCES operators(id)
);
CREATE TABLE product_description(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_id INT,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_product_description FOREIGN KEY (product_id) REFERENCES product(id)
);
CREATE TABLE payment_method(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Nomor 2C
CREATE TABLE transaction(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    payment_method_id INT,
    status VARCHAR(10),
    total_qty INT,
    total_price NUMERIC(25,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_transaction_user FOREIGN KEY (user_id) REFERENCES user(id),
    CONSTRAINT fk_transaction_payment_method FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
);
CREATE TABLE transaction_detail(
	transaction_id INT,
    product_id INT,
    status VARCHAR(10),
    qty INT,
    price NUMERIC(25,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (transaction_id, product_id)
);

-- Nomor 3
CREATE TABLE kurir(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Nomor 4
ALTER TABLE kurir ADD COLUMN ongkos_dasar INT AFTER name;

-- Nomor 5
ALTER TABLE kurir RENAME TO shipping;

-- Nomor 6
DROP TABLE shipping;

-- Nomor 7A
CREATE TABLE payment_method_description(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
ALTER TABLE payment_method ADD COLUMN payment_method_description_id INT AFTER id;
ALTER TABLE payment_method ADD CONSTRAINT fk_payment_method_description FOREIGN KEY (payment_method_description_id) REFERENCES payment_method_description(id);

-- Nomor 7B
CREATE TABLE user_address(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    address VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
ALTER TABLE user RENAME COLUMN address TO address_id;
ALTER TABLE user MODIFY COLUMN address_id INT;
ALTER TABLE user ADD CONSTRAINT fk_user_address FOREIGN KEY (address_id) REFERENCES user_address(id);

-- Nomor 7C
CREATE TABLE user_payment_method_detail(
    payment_method_id INT,
	user_id INT,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (user_id, payment_method_id)
);