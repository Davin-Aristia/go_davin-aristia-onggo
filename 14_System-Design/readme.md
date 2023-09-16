Diagram adalah representasi simbolis dari informasi dengan menggunakan teknik visualisasi. Beberapa aplikasi yang bisa digunakan untuk mendesign diagram yaitu Smart Draw, Lucid Chart, Whimsical, draw.io, dan visio.

Diagram yang paling sering digunakan adalah:
- Flowchart: Biasa disebut juga bagan alur adalah diagram yang menampilkan langkah-langkah dan keputusan untuk melakukan sebuah proses dari suatu program. Setiap langkah digambarkan dalam bentuk diagram dan dihubungkan dengan garis atau arah panah.
- Use Case: Menggambarkan hubungan interaksi antara sistem dan aktor. Use Case dapat mendeskripsikan tipe interaksi antara si pengguna sistem dengan sistemnya. 
- Entity Relationship Diagram: Biasa disebut juga diagram hubungan entitas adalah sebuah diagram yang digunakan untuk perancangan suatu database dan menunjukan relasi atau hubungan antar objek atau entitas beserta atribut-atributnya secara detail. Dengan menggunakan ERD, sistem database yang sedang dibentuk dapat digambarkan dengan lebih terstruktur dan terlihat rapi.
- High Level Architecture: standar untuk simulasi terdistribusi, digunakan saat membangun simulasi untuk tujuan yang lebih besar dengan menggabungkan beberapa simulasi.

Beberapa hal yang harus dipertimbangkan saat design sistem yang besar:
- Bagian-bagian arsitektur berbeda apa saja yang dapat digunakan
- Bagaimana bagian-bagian tersebut dapat bekerja satu sama lain
- Bagaimana kita dapat memanfaatkan bagian-bagian ini dengan sebaik-baiknya: apa saja pengorbanan yang tepat

Karakteristik utama dalam sistem distribusi:
- Scalability: merupakan kemampuan sebuah sistem, proses, atau network untuk tumbuh dan memanajemen demand yang bertambah banyak. Setiap sistem terdistribusi yang terus berkembang guna mendukung bertambahnya jumlah pekerjaan. Sebuah sistem harus scalable karena alasan seperti daa yang semakin banyak atau pekerjaan yang semakin banyak, seperti banyaknya jumlah transaksi. Dalam scalability terdapat vertical scale yang dimana menambah power ke sebuah prosesor, dan horizontal scale adalah menambah mesin atau node ke dalam sistem
- Reliability: Probabilitas sebuah sistem akan gagal dalam sebuah periode tertentu. Sistem terdistribusi akan dianggap reliable apabila bisa menyampaikan layanannya meskipun satu atau beberapa komponen hardware atau software rusak
- Availability: Waktu sebuah mesin tetap bisa beroperasi untuk melakukan fungsinya dalam periode tertentu. Availability merupakan pengukuran sederhana dari persentase waktu suatu sistem, layanan, atau mesin tetap beroperasi dalam kondisi normal
- Eficiency: Untuk memahami cara mengukur efisiensi sistem terdistribusi, apabila diasumsikan operasi yang berjalan secara terdistribusi dan menghasilkan serangkaian item sebagai hasilnya. Dua ukuran standar efisiensinya adalah waktu respons (atau latensi) yang menunjukkan penundaan untuk mendapatkan item pertama dan throughput (atau bandwidth) yang menunjukkan jumlah item yang dikirimkan dalam satuan waktu tertentu (misalnya, satu detik).
- Serviceability atau Manageability: adalah kesederhanaan dan kecepatan perbaikan atau pemeliharaan suatu sistem; jika waktu untuk memperbaiki sistem yang gagal semakin lama, maka availability akan berkurang

Dalam sebuah sistem software, sebuah job queue(batch queue) merupakan sebuah data struktur yang dikelola oleh perangkat lunak penjadwal pekerjaan yang berisi pekerjaan yang akan dijalankan

Load Balancer (LB) adalah komponen penting lainnya dari sistem terdistribusi. LB membantu menyebarkan lalu lintas ke sekelompok server untuk meningkatkan daya tanggap dan ketersediaan aplikasi, situs web, atau database. LB juga melacak status semua sumber daya saat mendistribusikan permintaan. Jika server tidak tersedia untuk menerima permintaan baru atau tidak merespons atau memiliki tingkat kesalahan yang tinggi, LB akan berhenti mengirimkan lalu lintas ke server tersebut. Untuk memanfaatkan skalabilitas dan redundansi, kita dapat mencoba menyeimbangkan beban pada setiap lapisan sistem. Kita dapat menambahkan LB di tiga tempat:
- Diantara user dan web server
- Diantara web server dan internal platform layer, seperti server aplikasi dan server cache
- Diantara internal platform layer dan database

Aplikasi Monolithic memiliki single code base dengan banyak modul. Modul dibagi berdasarkan fitur bisnis atau fitur teknis. Monolithic memiliki sistem build tunggal yang membangun seluruh aplikasi dan/atau dependency. Ia juga memiliki biner tunggal yang dapat dieksekusi atau diterapkan.
Microservices adalah layanan yang dapat diterapkan secara independen yang dimodelkan di sekitar domain bisnis. Microservices berkomunikasi satu sama lain melalui jaringan, dan sebagai pilihan arsitektur, mereka menawarkan banyak pilihan untuk memecahkan masalah yang mungkin dihadapi. Oleh karena itu, arsitektur microservices didasarkan pada beberapa microservices yang berkolaborasi.

SQL merupakan relational database yang terstruktur dan memiliki skema yang telah ditentukan seperti buku telepon yang menyimpan nomor telepon dan alamat. Beberapa benefit relational DB:
- Dirancang untuk segala keperluan
- SQL memiliki standar yang jelas
- Memiliki banyak tool(administrasi, reporting, framework)
SQL juga memiliki sifat ACID:
- Atomicity: transaksi terjadi semua atau tidak sama sekali
- Consistency: data tertulis merupakan data valid yang ditentukan berdasarkan aturan tertentu
- Isolation: pada saat terjadi request yang bersamaan(concurrent), memastikan bahwa transaksi dieksekusi seperti dijalankan secara sekuensial
- Durability: jaminan bahwa transaksi yang telah tersimpan, tetap tersimpan

NoSQL(Not only SQL) merupakan non-relational database yang tidak terstruktur dan memiliki skema yang dinamis seperti file folder yang menyimpan semuanya mulai dari alamat dan nomor telepon seseorang hingga 'like' Facebook dan preferensi belanja online mereka. NoSQL menyediakan mekanisme yang lebih fleksibel dibandingkan dengan model RDBMS (Sifat ACID) dan menghindari:
- Effort pada sifat transaksi ACID
- Kompleksitas SQL
- Design schema di depan
- Transactions (ditangani oleh aplikasi)

Kelebihan NoSQL: schema less, fast development
Kapan menggunakan NoSQL: membutuhkan skema fleksibel, ACID tidak diperlukan, data logging (json), data sementara (cache)
Kapan tidak direkomendasikan: data finansial, data transaksi, business critical, ACID - compliant

Perbedaan Skema SQL dan NoSQL:
Skema SQL:
- Tidak bisa menambah data yang tidak sesuai skema
- Perlu menambahkan data NULL pada item yang tidak memiliki data
- Memiliki tipa data yang strict
- Tidak dapat menambah beberapa item data pada sebuah field

Skema NoSQL:
- Tidak memiliki skema ketika menambahkan data
- Aplikasi menangani proses validasi tipe data
- Mendukung proses agregasi dokumen pada item

Dalam hal teknologi basis data, tidak ada satu solusi yang bisa diterapkan untuk semua masalah. Itu sebabnya banyak bisnis mengandalkan database relasional dan non-relasional untuk berbagai kebutuhan. Meskipun database NoSQL semakin populer karena kecepatan dan skalabilitasnya, masih ada situasi di mana database SQL yang sangat terstruktur mungkin memiliki kinerja yang lebih baik; memilih teknologi yang tepat bergantung pada kasus penggunaan


Cache yang digunakan pada data yang barusan diminta kemungkinan besar akan diminta lagi. Cache digunakan di hampir setiap lapisan komputasi: perangkat keras, sistem operasi, browser web, aplikasi web, dan banyak lagi. Cache seperti memori jangka pendek: memiliki jumlah ruang terbatas, namun biasanya lebih cepat daripada sumber data asli dan berisi item yang paling terakhir diakses. Cache bisa ada di semua level dalam arsitektur, namun sering kali ditemukan di level yang paling dekat dengan front end di mana cache diterapkan untuk mengembalikan data dengan cepat tanpa membebani level downstream.

Redundansi adalah duplikasi komponen atau fungsi penting suatu sistem dengan tujuan untuk meningkatkan keandalan sistem, biasanya dalam bentuk cadangan atau fail-safe, atau untuk meningkatkan kinerja sistem yang sebenarnya. Misalnya, jika hanya ada satu salinan file yang disimpan di satu server, maka kehilangan server tersebut berarti kehilangan file tersebut. Karena kehilangan data jarang merupakan hal yang baik, kita dapat membuat salinan file duplikat atau berlebihan untuk mengatasi masalah ini

Pengindeksan adalah cara untuk mengoptimalkan kinerja database dengan meminimalkan jumlah akses disk yang diperlukan saat query diproses. Pengindeksan adalah teknik struktur data yang digunakan untuk menemukan dan mengakses data dalam database dengan cepat. Kebanyakan database menggunakan B-Tree sebagai struktur data untuk pengindeksan yang memiliki complexity O(log n) untuk pencarian, penghapusan, dan penambahan

