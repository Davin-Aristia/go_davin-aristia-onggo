Recursion adalah keadaan di mana suatu fungsi menyelesaikan suatu masalah dengan memanggil dirinya sendiri berulang kali. Jika masalahnya cukup kecil, fungsi rekursi dapat menghasilkan jawaban langsung. Jika masalahnya terlalu besar, fungsi tersebut akan memanggil dirinya sendiri dengan cakupan masalah yang lebih kecil

Ada dua hal yang perlu dipikirkan ketika menggunakan strategi recursion:
- Base Case: apa kasus paling sederhana dari masalah ini
- Recurrence relations: apa hubungan rekursif dari masalah ini dengan masalah serupa yang lebih kecil

Contoh recursion:
func factorial(n int) int{
    if n == 1{
        return 1
    } else {
        return n * factorial(n-1)
    }
}

Number Theory adalah cabang matematika yang mempelajari bilangan bulat. Banyak sekali topik dalam bidang number theory yaitu Bilangan Prima, FPB, KPK, Faktorial, dan sebagainya. Dalam mencari bilangan prima, terdapat versi naif yang dimana kita looping nilai dari 2 sampai n-1 dan mengecek apakah habis dibagi. Terdapat pula solusi yang lebih efisien yaitu mengecek dari nilai 2 sampai akar dari n.

Searching merupakan proses mencari nilai dari sebuah list of value. Terdapat dua algoritma searching, yaitu linear search dan binary search

Sorting adalah proses menyusun data dalam urutan tertentu. Biasanya kita mengurutkan berdasarkan nilai elemennya. Kita bisa mengurutkan angka, kata, pasangan. Misalnya, kita dapat mengurutkan siswa berdasarkan tinggi badan mereka, dan kita dapat mengurutkan kota berdasarkan abjad atau berdasarkan jumlah penduduknya. Urutan yang paling banyak digunakan adalah urutan numerik dan urutan abjad. Terdapat beberapa algoritma sorting, yaitu Selection sort, Counting sort, Merge Sort, Quick sort, dan masih banyak lagi

Stack menerapkan prinsip Last-In First-Out. Stack memiliki beberapa properti:
- push(element): memasukkan element ke dalam stack
- pop(): menghapus element pertama yang terdapat di dalam stack
- top(): mengecek element pertama yang terdapat di dalam stack

Queue menerapkan prinsip First-In First-Out. Queue memiliki beberapa properti:
- enqueue(element): memasukkan element ke dalam queue
- dequeue(): menghapus element pertama yang terdapat di dalam queue
- front(): mengecek element pertama yang terdapat di dalam queue

Set adalah kumpulan elemen berbeda yang terdefinisi dengan baik, dianggap sebagai elemen tersendiri, sedangkan map merupakan sebuah dictionary. Set dan Map masing-masing memiliki dua implementasi:
- Unordered: wadah asosiatif yang berisi sekumpulan objek unik bertipe Kunci
- Ordered: wadah asosiatif yang diurutkan yang berisi pasangan nilai kunci dengan kunci unik