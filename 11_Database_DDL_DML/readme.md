Database adalah sekumpulan data yang terorganisir. Contohnya pada akun twitter terdapat data yang tersimpan pada database dalam bentuk:
- Display Name: Davin Aristia
- Username: @davin_aristia
- Bio: Saya adalah seorang Programmer...
- Location: Indonesia
- Join Date: 21/12/2018

Database memiliki beberapa relationship:
- One to One: 1 user hanya memiliki 1 foto profil
- One to Many: 1 user bisa memiliki banyak tweets
- Many to many: 1 user bisa memiliki banyak follower user, 1 user bisa di follow banyak user

Relational Database Management Systems(RDBMS) merupakan sebuah software yang menggunakan Relational Database Model sebagai dasarnya seperti MySQL yang bisa didownload dan install dari https://dev.mysql.com/downloads/windows/installer/


Pada SQL terdapat beberapa jenis perintah:
- Data Definition Language(DDL): perintah yang digunakan unruk membuat, mengubah, dan menghapus struktur database, namun bukan data di dalamnya
    1. CREATE DATABASE database_name: perintah ini digunakan untuk membuat database dengan nama database_name
    2. USE database_name: perintah ini digunakan untuk aktif pada database_name
    3. CREATE TABLE table_name: perintah ini digunakan untuk membuat tabel table_name pada database yang sedang kita aktif menggunakan perintah use database_name. Contoh query untuk create table adalah:
        CREATE TABLE table_name(
            column1 data_type PRIMARY KEY,
            column2 data_type FOREIGN KEY,
            ....
            column data_type
            PRIMARY KEY(one or more columns)
        );
    4. DROP TABLE table_name: perintah ini digunakan untuk menghapus tabel table_name
    5. RENAME TABLE table_name: perintah ini digunakan untuk mengubah nama tabel table_name

- Data Manipulation Language(DML): perintah yang digunakan untuk memanipulasi data dalam tabel dari suatu database, yang terdiri dari:
    1. INSERT: Input data ke table user. Contoh query untuk insert adalah:
        INSERT INTO USERS (username,fullname,status,gender,email,password,location)
        VALUES ('Davin','Davin Aristia',1,'m','davin@gmail.com','123456','Jawa Timur')
    2. SELECT: Menampilkan semua data pada tabel. Contoh query untuk select adalah:
        - SELECT * FROM USERS: menampilkan semua data yang terdapat pada tabel users
        - SELECT username,fullname FROM USERS WHERE id = 1: menampilkan username dan fullname pada tabel user yang id nya 1
        - SELECT username,fullname FROM USERS WHERE fullname IS NOT NULL: menampilkan username dan fullname dari table user yang usernamenya tidak kosong
    3. UPDATE: mengubah data yang terdapat pada tabel. Contoh query untuk update adalah:
        UPDATE USERS SET fullname ='Davin Aristia Onggo' WHERE id = 1
    4. DELETE: Hapus data pada tabel user. Contoh query untuk delete adalah:
        DELETE FROM USERS WHERE id = 1

- Data Control Language(DCL): perintah yang digunakan untuk mengontrol akses ke data yang disimpan dalam database, terdiri dari:
    1. LIKE / BETWEEN: 
        - Tampilkan data user_id dan message table tweets yang message mengandung huruf H didepan:
            SELECT user_id,type,message,parent_id FROM tweets WHERE message LIKE 'H%'
        - Tampilkan data user_id dan message table tweets user_id antara 1 dan 3
            SELECT user_id,type,message,parent_id FROM tweets WHERE user_id BETWEEN 1 AND 3
    2. AND / OR: 
        - Tampilkan data user_id dan message tabke tweets yang message mengandung huruf H didepan atau user_id antara 1 dan 3:
            SELECT user_id,type,message,parent_id FROM tweets WHERE message LIKE 'H%' OR user_id BETWEEN 1 AND 3
        - Tampilkan data user_id dan message tabke tweets yang message mengandung huruf H didepan dan user_id antara 1 dan 3:
            SELECT user_id,type,message,parent_id FROM tweets WHERE message LIKE 'H%' AND user_id BETWEEN 1 AND 3
    3. ORDER BY:
        - Tampilkan data user_id dan message table tweets yang message mengandung huruf H didepan atau user_id antara 1 dan 3 diurutkan berdasarkan id tweets berurutan dari terbesar ke terkecil:
            SELECT id,user_id,type,message,parent_id
            FROM tweets
            WHERE message LIKE 'H%' OR user_id BETWEEN 1 AND 2 ORDER BY id DESC
    4. LIMIT
        - Tampilkan data user_id dan message table tweets yang message mengandung huruf H didepan atau user_id antara 1 dan 3 diurutkan berdasarkan id tweets berurutan dari terbesar ke terkecil di batasi 2 data:
            SELECT id,user_id,type,message,parent_id
            FROM tweets
            WHERE message LIKE 'H%' AND user_id BETWEEN 1 AND 3 ORDER BY id DESC LIMIT 2