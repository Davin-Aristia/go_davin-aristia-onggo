Arsitektur yang baik adalah pemisahan perhatian menggunakan lapisan untuk membangun aplikasi yang modular, terukur, dapat dipelihara, dan dapat diuji. Mengerjakan proyek pemrograman mirip dengan merencanakan dan membangun kawasan perumahan. Jika sudah direncanakan/diketahui bahwa satu distrik akan diperluas, maka diperlukan untuk menyediakan sebuah ruang/tempat untuk mengembangkan distrik tersebut. Jika tidak, untuk mengembangkan distrik baru terpaksa harus menghancurkan bangunan dan jalanan untuk membuat ruang/tempat untuk bangunan baru

Setiap tim memiliki strukturnya masing-masing:
- How to maintain this
- Mobility issue: Sulit untuk mentransfer pengetahuan dalam business domain dan ditambah dengan struktur kode yang unik sehingga dibutuhkan waktu untuk mempelajarinya secara top down
- Another issue: Sulit untuk mengimplementasi unit test khususnya test yang memiliki koneksi dengan database

Salah satu solusi untuk beberapa permasalahan diatas adalah Hexagonal Architecture. Dalam beberapa tahun terakhir, terdapat beberapa sistem arsitektur diantaranya:
- Hexagonal Architecture
- Onion Architecture
- Screaming Architecture
- DCI from Agile Development
- BCE from Object Oriented Software
- Clean Architecture
Meskipun arsitektur tersebut berbeda memiliki detail yang berbeda-beda, namun arsitektur tersebut sangat mirip. Semua memiliki objektif yang sama, yaitu separation of concern. Semua architecture tersebut mencapai separation dengan membagi software menjadi lapisan. Setiap lapisan memiliki satu lapisan untuk aturan bisnis dan yang lain untuk interface

Beberapa constraint sebelum design Clean Architecture:
- Independent of Frameworks: Mengizinkan untuk penggunaan framework sebagai alat daripada memiliki keterbatasan constraint
- Testable: Aturan bisnis bisa diuji tanpa barang lainnya
- Independent of UI: UI bisa berubah dengan gampang, tanpa mengubah seluruh sistem
- Independent of Database: Aturan bisnis tidak terikat dengan database, sehingga bisa menukar database dengan gampang
- Independent of any external: Aturan bisnis sama sekali tidak mengetahui apa pun tentang dunia luar

Beberapa benefit menggunakan design Clean Architecture:
- Merupakan struktur yang standar sehingga gampang untuk switching projek
- Development yang lebih cepat dalam jangka panjang
- Mocking dependencies menjadi hal yang tidak begitu penting dalam unit test
- Perpindahan yang mudah dari prototype ke solusi yang layak (misalnya mengubah in-memory storage menjadi SQL database)

Classic 3-layer architecture (Bisa memiliki lebih dari 3 layer):
- Entities layer (Optional): Objek bisnis yang mencerminkan konsep yang dikelola aplikasi
- Use Case - Domain Layer: Berisi business logic
- Controller - Presentation layer: Menyajikan data ke dalam layar dan menangani user interactions
- Drivers - Data layer: Mengelolaa data aplikasi, misalnya mengambil data dari network, dan mengelola data cache

Terdapat beberapa pendekatan 3-layer architecture:
- Pendekatan 1: CA Layer dalam 1 modul
    Misalnya: presentation layer, Domain Layer, Data Layer berada pada 1 module yaitu app module
- Pendekatan 2: 1 CA Layer per modul
    Misalnya: Presentation layer berada pada presentation module, Domain layer berada pada domain module, Data layer berada pada data module
- Pendekatan 3: CA Layer di dalam feature module
    Misalnya: Terdapat feature_a module dan feature_b module yang masing-masing berisi presentation layer, Domain Layer, Data Layer. Yang dimana modul feature tersebut berada di dalam app module
- Pendekatan 4: CA Layer per feature modules
    MIsalnya: Presentation layer berada pada feature_a presentation module, Domain layer berada pada feature_a domain module, Data layer berada pada feature_a data module yang dimana semua module tersebut berada di dalam app module

Domain Driven Design(DDD) merupakan sebuah pendekatan untuk mengembangkan software kompleks yang menghubungkan konsep bisnis dengan implementasi teknis secara mendalam. DDD ini bukanlah sebuah metodologi atau teknologi tetapi lebih ke pendekatan secara praktikal dan pendefinisian (terminologi) yang fokus terhadap keputusan desain software dalam hubungannya dengan domain bisnis yang kompleks. DDD ini umumnya digunakan dalam pengembangan software dengan berorientasi object dan pendekatan agile. 

Dalam Domain Driven Design terdapat istilah "You Need to go slow, to go fast", yang berarti memastikan bahwa kita telah menyelesaikan masalah yang valid secara optimal. Seletah itu baru implementasi sedemikian rupa sehingga bisnis kita dapat memahaminya tanpa memerlukan terjemahan tambahan dari bahasa teknis

Kesimpulan:
- Clean Architecture adalah Software Architecture
- Domain Driven Design adalah software design technique, sebelum membuat sebuah software, kita bisa memetakan dan mencari solusi terbaik dahulu dengan Domain Driven Design
Keduanya melengkapi satu sama lain

Package context atau bisa disebut context adalah seuatu package yang membawa deadline, cancellation signal, atau value lain pada request/permintaan api:
1. Implementasi Context
    Code: var ctx = context.Background()
    Penjelasan: Fungsi background() akan mengembalikan root context dimana kita bisa memakainya untuk operasi lain
2. Context dengan Value
    Code: ctx = context.WithValue(ctx, "key", "value")
    Penjelasan: Implementasi context dengan value akan sering kita lihat pada middleware. Kita dapat mengirim data dari middleware (contoh user_id dari token) ke handler menggunakan context

    Code: var newRequest = c.Request().WithContext(
            context.WithValue(c.Request().Context(), "key", user_id),
          )
    Penjelasan: - Menambahkan value ke context yang terdapat dari http.Request
                - Parent/root context diambil dari http.Request

    Code: var ctx = c.Request().Context()
          user_id := ctx.Value("key").(int)
    Penjelasan: - Mengambil context dari http.Request
                - Mengambil value dari context menggunakan key yang sama, lalu parsing ke tipe datanya

3. Context dengan Cancellation:
    Normal Request: Request -> Service -> Database
    Apabila user cancel request tanpa menggunakan context cancellation maka data masih tetap diteruskan ke database, namun dengan menggunakan context cancellation maka data tidak diteruskan ke database


Terdapat 3 opsi bila melakukan migrasi arsitektur kode dari mvc ke clean code:
- Pertahankan desain sekarang dengan memisahkan dependensi
- Pertahankan desainnya tetapi pindahkan kodenya kedalam suatu layer
- Ubah desainnya dan pisahkan dependensi

Biasanya struktur kode yang dipakai dalam migrasi arsitektur kode dari mvc ke clean code:
- Controller: berisi kode yang berhubungan langsung dengan user (interface layer)
- Repository: berisi kode yang berhubungan langsung dengan database (interface layer)
- Usecase: berisi bisnis logic yang dipakai