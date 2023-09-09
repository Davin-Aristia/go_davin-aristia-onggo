JOIN adalah sebuah klausa untuk mengkombinasikan record dari dua atau lebih tabel, join terdiri dari beberapa jenis:
- Inner Join: mengembalikan baris-baris dari dua tabel atau lebih yang memenuhi syarat, contohnya adalah:
    SELECT t.message FROM users u
    INNER JOIN tweets t
    On u.id = t.user_id
- LEFT JOIN: mengembalikan seluruh baris dari tabel disebelah kiri yang dikenai kondisi ON dan hanya baris dari tabel disebelah kanan yang memenuhi kondisi join
    SELECT u.username, t.message
    FROM users u
    LEFT JOIN tweets t
    ON u.id = t.user_id
- RIGHT JOIN: mengembalikan seluruh baris dari tabel disebelah kanan yang dikenai kondisi ON dengan data dari tabel sebelah kiri yang memenuhi kondisi join. Teknik ini merupakan kebalikan dari left join
    SELECT u.username, t.message
    FROM users u
    RIGHT JOIN tweets t
    ON u.id = t.user_id

UNION digunakan untuk menggabungkan data dari hasil dua atau lebih kueri perintah SELECT menjadi satu kumpulan hasil yang berbeda. Operator ini menghapus semua duplikat yang ada dalam hasil yang digabungkan. Contoh query Union adalah:
    SELECT username,fullname
    FROM users WHERE id=1
    UNION
    SELECT username,fullname
    FROM users WHERE id=2

Fungsi Agregasi adalah fungsi di mana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal, agregasi terdiri dari:
- MIN: fungsi dimana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal
    Contoh: SELECT MIN(id) AS id FROM users: menampilkan id terkecil dari tabel users
- MAX: digunakan untuk mendapatkan nilai maximum atau nilai terbesar dari sebuah data record di tabel
    Contoh: SELECT MAX(id) FROM users: menampilkan nilai id terbesar dari tabel users
- SUM: digunakan untuk mendapatkan jumlah total nilai dari sebuah data atau record di tabel
    Contoh: SELECT SUM(favourite_count) FROM tweets WHERE user_id = 1: menampilkan jumlah total favourite_count dari tabel tweets dengan user_id 1
- AVG: digunakan untuk mencari nilai rata-rata(average) dari sebuah data atau record di tabel
    Contoh: SELECT AVG(favourite_count) FROM tweets WHERE user_id = 1: menampilkan nilai rata-rata favourite_count dari tabel tweets dengan user_id 1
-COUNT: digunakan untuk mencari jumlah dari sebuah data atau record di tabel
    Contoh: SELECT COUNT(1) FROM tweets WHERE user_id = 1: menampilkan jumlah data dari tabel tweets dengan user_id 1
- HAVING: digunakan untuk menyeleksi data berdasarkan kriteria tertentu, dimana kriteria berupa fungsi agregat
    Contoh: SELECT user_id FROM tweets GROUP BY user_id HAVING SUM(favourite_count)> 2: menampilkan data dari tabel tweets dengan jumlah total favourita_count per user lebih dari 2

Subquery atau Inner query atau Nested query adalah query di dalam query SQL lain yang digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil. Subquery dapat digunakan dengan SELECT, INSERT, UPDATE, dan DELETE statements bersama dengan operator seperti =,<,>,>=,<=, IN, BETWEEN, dan sebagainya. Beberapa peraturan dalam penggunaan subquery:
- Harus tertutup dalam tanda kurung
- Sebuah subquery hanya dapat memiliki satu kolom pada klausa SELECT, kecuali beberapa kolom yang di query utama untuk subquery untuk membandingkan kolom yang dipilih
- Subqueries yang kembali lebih dari satu baris hanya dapat digunakan dengan beberapa value operator, seperti operator IN
- Daftar SELECT tidak bisa menyertakan referensi referensi ke nilia-nilai yang mengevaluasi ke BLOB, ARRAY, CLOB, atau NCLOB
- Sebuah subquery tidak dapat segera tertutup dalam fungsi set

Contoh Subquery:
- Tampilkan data tabel users yang user_id nya ada pada tabel tweets:
    SELECT * FROM USERS WHERE id IN
    (SELECT user_id FROM tweets GROUP BY user_id)
- Tampilkan data tabel users yang jumlah total favourite_count per user lebih dari 5 pada tabel tweets:
    SELECT * FROM USERS WHERE id IN
    (SELECT user_id FROM tweets GROUP BY user id HAVING SUM(favourite_count)> 5)

Function merupakan sebuah kumpulan statement yang akan mengembalikan sebuah nilai balik pada pemanggilnya. Contoh sebuah function untuk mengembalikan jumlah data tweets per user:
    DELIMITER $$
    CREATE FUNCTION sf_count_tweet_peruser (user_id_p int) RETURNS INT DETERMINISTIC
    BEGIN
    DECLARE total INT;
    SELECT COUNT(*) INTO total FROM tweets
    WHERE user_id = user_id_p AND type = 'tweets';
    RETURN total;
    EMD$$
    DELIMITER ;
Penjelasan:
- Delimiter: memberitahu kepada mysql soal delimiter yang digunakan, secara default menggunakan ; jadi bila ada tanda ; mysql akan mengartikan akhir dari statement, pada contoh di atas delimiter yang digunakan $$ jadi akhir statementnya adalah $$
- CREATE FUNCTION: adalah header untuk membuat function
- RETURN: adalah untuk menentukan tipe data yang di return-kan oleh function
- DETERMINISTIC/NOT DETERMINISTIC: adalah untuk menentukan yang bisa menggunakan function ini adalah untuk user pembuatnya saja (deterministic) atau user siapa saja (not deterministic)
- BEGIN END: body dari function jadi semua SQL tulis diantara begin dan end

Contoh lainnya adalah membuat trigger function untuk delete data yang berhubungan dengan table users yang sedang di delete datanya:
    DELIMITER $$
    CREATE TRIGGER delete_all_data_user
    BEFORE DELETE ON users FOR EACH ROW
    BEGIN
    // declare variables
    DECLARE v_user_id INT;
    SET v_user_id = OLD.id;
    // trigger code
    DELETE FROM tweets WHERE user_id = v_user_id;
    DELETE FROM user_followers WHERE
    user_id=v_user_id;
    EMD$$
Penjelasan:
- CREATE TRIGGER: adalah header untuk membuat trigger function
- DECLARE: adalah syntax untuk mendeclare kan variabel
- OLD: variabel yang menyimpan nilai dari dalam data yang sedang berinteraksi/dipanggil
- NEW: adalah variabel yang menyimpan nilai dari data yang baru masuk/sedang di input