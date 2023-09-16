Command Line merupakan antarmuka berbasis teks yang cepat dan kuat yang digunakan developer untuk berkomunikasi secara lebih efektif dan efisien dengan komputer untuk menyelesaikan serangkaian tugas yang lebih luas. Beberapa keuntungan menggunakan Command Line:
- Kontrol Granular dari sebuah OS atau aplikasi
- Pengelolaan lebih cepat dari sejumlah besar sistem operasi
- Kemampuan untuk menyimpan skrip untuk mengotomatiskan tugas-tugas rutin
- Command Line Interface menghubungkan pengetahuan untuk membantu troubleshooting, seperti masalah koneksi jaringan

Shell merupakan Command Line Interface untuk layanan OS seperti UNIX Shell dan Command Prompt (MSDOS). Selain itu juga terdapat beberapa aplikasi CLI seperti Python REPL, MySQL CLI client, Mongo shell, dan sebagainya

Pada sebuah terminal di dalamnya terdapat shell tempat ketik command yang sebelum tempat untuk mengetik terdapat tulisan seperti me@linuxbox:~$:
- me: username
- linuxbox: hostname
- ~: current path (home)
- $: shell untuk normal user(diizinkan untuk hanya bisa memanipulasi directory /home/username)
- #: shell untuk root user(diizinkan untuk memanipulasi semua directory /dir)
Selain itu terdapat normal user + perintah sudo yang memungkinkan normal user untuk bertindak sebagai root user pada 1 perintah tersebut

Perintah directory populer pada UNIX Shell:
- pwd: perintah untuk menampilkan working directory yang sedang dipakai
- ls: perintah untuk menampilkan list file atau list directory yang ada di dalam working directory, untuk menampilkan file dan folder yang hidden, bisa menggunakan ls -a
- mkdir: perintah untuk membuat folder, contohnya adalah mkdir nama_folder
- cd: perintah untuk pindah ke working directory lain, contohnya adalah cd /Users
- rm: perintah untuk delete sebuah file, sedangkan untuk menghapus folder menggunakan rmdir
- cp: perintah untuk mengcopy sebuah file ke folder yang lain
- mw: perintah untuk memindahkan sebuah file ke folder yang lain
- ln: perintah untuk membuat sebuah link

* perintah ls juga memiliki ls -l yang akan menampilkan list secara detail, yang akan memunculkan hasil seperti:
drwxrwxrwx@ 2 Fransiskaariana wheel 64 Feb 8 06:58 access_directory
- drwxrwxrwx: merupakan sebuah permission dimana d menandakan directory dan - menandakan file, rwx yang pertama artinya permission untuk read write execute untuk owner, rwx yang kedua artinya permission untuk read write execute untuk grup owner, rwx yang pertama artinya permission untuk read write execute untuk other user
- 2: merupakan terdapat 2 file
- Fransiskaariana: nama owner
- wheel: group user
- 64: ukuran
- Feb 8 06:58: tanggal dibuat
- access_directory: nama file/directory

Perintah files populer pada UNIX Shell:
- touch: perintah untuk membuat file
- vim, nano: perintah untuk mengedit sebuah file dengan membuka sebuah text editor yang berisi konten dari file tersebut
- head, cat, tail, less: menampilkan konten dari sebuah file
- chmod: perintah untuk mengganti permission dari sebuah file
- chown: perintah untuk mengganti owner dari sebuah file
- diff: perintah untuk membandingkan 1 file dengan file lainnya perbedaannya ada dimana

Perintah network populer pada UNIX Shell:
- ping: perintah untuk melakukan pengecekan koneksi
- ssh: perintah untuk melakukan koneksi ke remote server
- netstat: perintah untuk mengetahui port yang sedang digunakan dan apa saja yang sedang aktif
- nmap: perintah untuk menganalisis port
- ip addr (ifconfig): perintah untuk menampilkan ip address
- wget, curl: perintah untuk mengunduh file dari url
- telnet: hampir sama seperti ssh yaitu untuk melakukan koneksi ke remote server namun secara keamanan telnet lebih rendah
- netcat: perintah yang digunakan untuk monitoring jaringan

Perintah utility populer pada UNIX Shell:
- man: deskripsi dan penjelasan dari sebuah perintah
- env: environment yang berisi variabel path
- echo: perintah untuk menampilkan/print sebuah text
- date: perintah untuk menampilkan waktu saat ini
- which: perintah untuk lokasi dari sebuah program file
- watch: perintah untuk memonitoring sebuah program 
- sudo: perintah sudo yang memungkinkan normal user untuk bertindak sebagai root user pada 1 perintah tersebut 
- history: perintah untuk menampilkan semua history command pada CLI
- grep: perintah untuk mencari file yang mengandung kata tertentu
- locate: perintah untuk pencarian sebuah file

Shell adalah sebuah program fungsi dimana menjembatani antara user dan kernel. Shell script adalah bahasa pemrograman yang dicomplie berdasarkan perintah shell