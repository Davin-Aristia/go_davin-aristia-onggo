Dalam sebuah pencarian Big Search, search engine tersebut akan mencari mengenai website, gambar, dan video. Pencarian tersebut terbagi menjadi beberapa istilah:
- Sequential Program: dalam sequential program, sebelum tugas berikutnya berjalan, tugas sebelumnya harus selesai. Contohnya adalah dalam pencarian Big Search tadi, misalnya dalam mencari video membutuhkan waktu 2 detik, dalam mencari website membutuhkan waktu 3 detik, dan mencari gambar membutuhkan waktu 4 detik, maka sebelum sequential program bisa mencari website, sequential program harus mencari video terlebih dahulu selama 2 detik
- Parallel Program: dalam Parallel program, beberapa tugas bisa dijalankan secara bersamaan namun membutuhkan mesin multi-core. Contohnya adalah dalam pencarian Big Search tadi, waktu yang dibutuhkan hanyalah 4 detik apabila sebuah mesin memiliki minimal 3 core, yaitu core pertama akan mencari video, core kedua akan mencari website, dan core keempat akan mencari gambar
- Concurrent Program: dalam Concurrent program, beberapa tugas bisa dieksekusi secara independen dan hasilnya bisa muncul secara bersamaan. Misalnya adalah dalam pencarian big search tadi, concurrent tidak perlu menunggu pencarian video selesai, bisa dilanjutkan dengan mencari website

Go Concurrency (goroutines) menggabungkan konsep parallel program dengan concurrent program. Biaya pembuatan goroutine terbilang kecil jika dibandingkan dengan thread. Thread adalah proses yang ringan, atau dengan kata lain thread adalah unit yang mengeksekusi kode di bawah program. Penggunaan goroutine bisa dengan cara menambahkan go di depan pemanggilan sebuah function. Gomaxprocs digunakan untuk mengontrol jumlah thread sistem operasi yang tersedia untuk eksekusi goroutine pada program go tertentu

Channel merupakan objek komunikasi yang dimana digunakan oleh goroutine untuk berkomunikasi antara satu dengan yang lainnya. Cara deklarasi sebuah channel adalah dengan menggunakan <- contohnya adalah data := <- c. Dalam channel juga terdapat istilah:
- Unbuffered Channel: yang dimana Unbuffered artinya sebuah channel hanya bisa menyimpan sebuah data dan harus dikeluarkan terlebih dahulu sebelum diisi dengan data lainnya.
- Buffered Channel: menentukan seberapa banyak data yang bisa disimpan oleh sebuah channel dalam sebuah waktu.

Select digunakan untuk mengontrol komunikasi data antar 1 atau banyak channel. Contoh penggunaannya adalah:
select{
    case avg := <-ch1: fmt.Println(//something)
    case max := <-ch2: fmt.Println(//another thing)
}

Race Condition adalah saat 2 thread yang mengakses memori yang sama pada saat yang sama, yang dimana satu sedang mengubah. Race condition terjadi karena akses yang tidak tersinkronisasi dengan memori yang berbagi(shared memory). Race condition bisa dicegah dengan cara:
- Blocking with Waitgroups: Cara yang paling mudah dalam menyelesaikan sebuah data race adalah dengan block akses untuk membcara data tersebut sampai operasi mengubah data selesai.
Contoh:
var wg sync.WaitGroup : inisialisasi variabel waitgroup
wg.Add(n): menunjukkan bahwa ada n task yang perlu ditunggu
wg.Wait(): block akses sampai wg.Done() dipanggil
wg.Done(): mengindikasi bahwa data tersebut sudah bisa diakses
- Channel Blocking: mengizinkan goroutine untuk sinkronisasi satu sama lain untuk sesaat
Contoh:
done := make(chan struct{}): membuat channel untuk push struct kosong saat sudah selesai
<-done: block akses sampai adanya struct dipush ke done dipanggil
done <- struct(){}: push struct kosong untuk menandakan bahwa data tersebut sudah bisa diakses
- Mutex: Dimana kita tidak memedulikan urutan dari membaca maupun mengubah data, yang diperlukan hanya mereka tidak terjadi bersama-sama
Contoh:
type SafeNumber struct{
    val int
    m sync.Mutex
}: deklarasi struct dengan instansi mutex
i := &SafeNumber(): membuat instansi SafeNumber
i.m.Lock(): block akses sampe metode unlock dipanggil
i.m.Unlock(): membuka akses untuk pengaksesan data