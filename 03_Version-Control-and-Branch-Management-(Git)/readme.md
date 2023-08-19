Versioning: Pengaturan versi/pelacakan perubahan dari sebuah program. Contohnya adalah apabila kita membuat revisi skripsi, maka kita perlu membuat beberapa file untuk tracking perubahan apa saja yang telah kita buat. Versioning dapat membuat pelacakan terhadap semua perubahan yang kita buat, dan siapa yang membuat perubahan tersebut

Git: Salah satu version control system populer yang digunakan para developer untuk mengembangkan software secara bersama-sama, berguna untuk tracking perubahan dalam suatu projek, dan dari tracking perubahan tersebut, bisa digunakan separalel dengan berkolaborasi dengan developer lain. Git bersifat terdistribusi dan bukan tersentralisasi artinya setiap developer memiliki file codingan dari git tersebut dan akan terkoneksi dengan developer lain melalui git itu sendiri

Terdapat 3 Git Server yang populer saat ini yaitu Github, Gitlab, dan Bitbucket yang bisa disebut sebagai sosial medianya para developer. Git Server tersebut menampung repository atau projek. Repository Github terbagi menjadi 2, yaitu:
1. Public: Bisa diakses oleh orang lain, dan biasanya bersifat open source
2. Private: Tidak bisa diakses oleh orang lain, biasanya digunakan untuk internal perusahaan

Cara Link projek lokal ke dalam Git:
1. Mengetikkan command git init untuk menginisialisasikan git, yang akan membuat sebuah hidden folder bernama .git
2. git remote add [nama remote] [link ssh or https] , untuk menambahkan remote git yang tersinkronisasi ke git

Cara push projek pada git:
1. Git add [nama file] / Git add .(untuk add semua perubahan): Untuk memindahkan perubahan pada file yang dipilih masuk ke dalam Staging Area
2. git commit -m "[komentar yang dilakukan pada perubahan]": Setelah Staging Area selesai ditandain, langkah selanjutnya adalah membuat satu paket staging area tersebut menjadi satu commit. Commit berisi komentar perubahan per staging area.
3. git push [nama remote] [nama branch]: untuk mengupload perubahan yang terdapat pada staging area ke dalam git repository

Command-command git yang sering digunakan:
1. git status: menunjukkan pada saat ini berada di branch mana, dan menunjukkan perubahan yang siap di commit.
2. git diff: menunjukkan perubahan apa saja yang terjadi pada projek
3. git stash: menyimpan perubahan yang dibuat pada tempat penyimpanan sementara di lokal yang dimana tidak dicommit maupun dipush pada repo dan perubahan yang dibuat pada projek akan hilang. Cara untuk mengambil perubahan yang distash adalah dengan menggunakan command "git stash apply"
4. git log: memunculkan commit id, author, dan tanggal dilakukan commit
5. git log --oneline: memunculkan commit id dan komentar commit
6. git reset [commit id] --soft: balik ke suatu titik commit dan perubahan yang terjadi tetap akan disimpan di staging area
7. git reset [commit id] --HARD: balik ke suatu titik commit dan tidak perubahan yang terjadi