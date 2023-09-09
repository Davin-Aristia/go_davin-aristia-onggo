CREATE DATABASE praktikum_database;
USE praktikum_database;

CREATE TABLE product_types(
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
CREATE TABLE products(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	product_type_id INT,
    operator_id INT,
    code VARCHAR(50),
    name VARCHAR(100),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_product_type FOREIGN KEY (product_type_id) REFERENCES product_types(id),
    CONSTRAINT fk_product_operator FOREIGN KEY (operator_id) REFERENCES operators(id)
);
CREATE TABLE product_descriptions(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_id INT,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_product_description FOREIGN KEY (product_id) REFERENCES products(id)
);
CREATE TABLE payment_methods(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE users(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    status SMALLINT,
    dob DATE,
    gender CHAR(1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE transactions(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    payment_method_id INT,
    status VARCHAR(10),
    total_qty INT,
    total_price NUMERIC(25,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_transaction_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_transaction_payment_method FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);
CREATE TABLE transaction_details(
	transaction_id INT,
    product_id INT,
    status VARCHAR(10),
    qty INT,
    price NUMERIC(25,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (transaction_id, product_id)
);