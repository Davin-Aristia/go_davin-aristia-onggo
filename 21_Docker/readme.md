Container Bukan sebuah virtual machine, melainkan sebuah proses dengan isolasi file system. Segala sesuatu dalam Linux merupakan sebuah file:
- /dev/sda = harddisk
- /dev/proc = processes
- /dev/usb
- /dev/cpu
- /dev/std(in|out)
- /bin/bash = file binary

Perbedaan Container dan Virtual Machines:
Container:
- Abstraction berada pada app layer
- Container memakan lebih sedikit ruang daripada Virtual Machine (gambar pada container biasanya hanya berukuran puluhan MB)
- Menangani lebih banyak aplikasi dan memerlukan lebih sedikit Virtual Machine dan Sistem Operasi

Virtual Machine:
- Abstraction pada fisik hardware
- Setiap Virtual Machine termasuk salinan lengkap sistem operasi
- Slow to Boot

Beberapa Syntax pada Docker:
- FROM: Mendapatkan Image dari docker registry
- RUN: Mengeksekusi bash command saat building container
- ENV: Set variabel didalam container
- ADD: Copy file dengan proses lain
- COPY: Copy file
- WORKDIR: set working file directory
- ENTRYPOINT: Mengeksekusi command saat building container sudah selesai
- CMD: Mengeksekusi command tapi bisa di overwrite

Docker Volume bisa dianggap sebagai storage atau tempat penyimpanan data di container. Tentunya saat kita membuat container kita tidak ingin ketika container kita mati data yang ada pada container ikut terhapus juga. Untuk itu kita dapat memanfaatkan Volume pada docker. Command untuk membuat docker adalah: docker create volume <name_volume>

Defaultnya container pada docker akan saling terisolasi satu sama lain. Kita tidak dapat melakukan request api dari container satu ke container lain. Untuk itu kita harus membuat dan mendaftarkan container pada network yang sama