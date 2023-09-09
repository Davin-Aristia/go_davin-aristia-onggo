Clean code adalah istilah untuk kode yang mudah dibaca, dipahami dan diubah oleh programmer. Clean code memiliki karakteristik sebagai berikut:
- Mudah dipahami: penamaan variabel ataupun function yang dibuat diusahakan sesuai dengan kegunaannya, misalnya: var nama, untuk menyimpan nama
- Mudah dieja dan dicari: contohnya untuk set timeout selama 1 hari, daripada langsung menuliskan 86400000(milidetik pada 1 hari), alangkah baiknya menyimpan pada sebuah variabel yang dinamakan milidetik_pada_sehari
- Singkat namun mendeskripsikan konteks: variabel yang dibuat akan lebih baik apabila singkat dan tepat sesuai dengan konteksnya
- Konsisten: terdapat beberapa gaya dalam penulisan sebuah variabel, misalnya adalah camel case yang setiap kata diawali dengan huruf besar kecuali kata pertama seperti "iniCamelCase" dan terdapat pula gaya snake case yang setiap kata dipisah dengan _ seperti "ini_snake_case". Arti karakteristik konsisten adalah apabila penamaan sebuah variabel menggunakan camel case, maka penamaan semua variabel juga harus menggunakan camel case
- Hindari penambahan konteks yang tidak perlu: contohnya pada sebuah struct car yang memiliki variabel carColor yang dimana kata car sudah tidak diperlukan karena sudah terdeskripsikan pada nama struct car
- Komentar: Komentar pada code biasanya digunakan untuk mendeskripsikan apa yang dilakukan kode tersebut dan komentar tersebut tidak akan dijalan oleh program. Akan lebih baik apabila menambahkan komentar pada setiap block dan menjelaskan block dari kode tersevbut dan tidak perlu per baris yang menjelaskan setiap baris dari kode
- Good Function: sebuah function yang bagus tidak menerima banyak parameter, apabila parameter yang perlu dikirimkan banyak, akan lebih baik apabila dibuat sebuah struct ataupun array
- Gunakan konvensi: Bisa menggunakan style guide Airbnb Javascript: https://github.com/airbnb/javascript ataupun style guide Google Python: https://google.github.io/styleguide/pyguide.html
- Formatting: Saran formatting pada code
    1. Lebar baris code 80-120 karakter
    2. Satu class 300-500 baris
    3. Baris code yang berhubungan saling berdekatan
    4. Dekatkan fungsi dengan pemanggilnya
    5. Deklarasi variabel berdekatan dengan penggunanya
    6. Perhatikan indentasi
    7. Menggunakan prettier atau formatter

Beberapa prinsip pada "Clean Code":
- Keep it So Simple(KISS): Hindari membuat fungsi yang dibuat untuk melakukan A sekaligus memodifikasi B, mengecek fungsi C.
    1. Fungsi atau class harus kecil
    2. Fungsi dibuat untuk melakukan satu tugas saja
    3. Jangan digunakan terlalu banyak argumen pada fungsi
    4. Harus diperhatikan untuk mencapai kondisi yang seimbang, kecil dan jumlahnya minimal
- Don't Repeat Yourself: Duplikasi code sering terjadi karena copy paste. Untuk menghindari duplikasi code buatlah fungsi yang dapat digunakan secara berulang-ulang
- Refactoring: proses restrukturisasi kode yang dibuat, dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal. Prinsip KISS dan DRY bisa dicapai dengan cara refactor. Teknik refactoring:
    1. Membuat sebuah abstraksi
    2. Memecah kode dengan fungsi/class
    3. Perbaiki penamaan dan lokasi kode
    4. Deteksi kode yang memiliki duplikasi