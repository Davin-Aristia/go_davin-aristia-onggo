Time Complexity: perhitungan sebuah algoritma mengenai seberapa lama waktu yang dibutuhkan algoritma tersebut menyelesaikan suatu permasalahan dengan menentukan seberapa banyak operasi dominan dalam operasi primitif seperti penambahan, perkalian, assignment, dan sebagainya. Dalam sebuah algoritma mungkin akan ada banyak code dan dalam menghitung time complexity, kita harus menentukan operasi mana yang paling dominan

func dominant (n int) int {
    var result int = 0
    for i:= 0; i < n; i++{
        result += 1
    }
    result = result + 10
    return result
}

Dalam contoh function dominan diatas, terdapat 2 operasi primitif, yaitu result += 1 dan result = result + 10. Namun operasi dominan adalah result += 1 karena operasi tersebut berada di dalam loop yang dimana akan berjalan lebih banyak kali daripada result = result + 10. Dalam contoh algoritma diatas, Time Complexity adalah O(n) karena operasi tersebut akan berjalan sebanyak n kali

Jenis-jenis Time Complexity
1. Constant Time - O(1): Operasi primitif tetap hanya akan berjalan satu kali tidak peduli kita input seberapa banyak di dalam N
2. Linear Time - O(n): Operasi primitif akan berjalan sesuai seberapa banyak nilai n yang diinputkan, misalnya n = 10, maka operasi primitif tersebut akan berjalan 10 kali pula. Biasanya time complexity O(n) ini ditemui ketika digunakan sebuah loop dalam suatu algoritma
3. Linear Time - O(n+m): Hampir mirip dengan O(n),O(n+m) artinya Operasi primitif akan berjalan sesuai seberapa banyak nilai n ditambah nilai m yang diinputkan, misalnya nilai n = 10 dan m = 5, maka operasi primitif akan berjalan 15 kali. Biasanya time complexity O(n+m) ini ditemui ketika digunakan 2 loop didalam suatu algoritma, dimana loop yang pertama menggunakan nilai n dan loop yang kedua menggunakan nilai m
4. Logarithmic Time - O(log n): Merupakan time complexity yang paling efektif dalam sebuah problem solving terkait dengan kecepatan dan hanya sedikit lebih lama dibandingkan O(1). Apabila nilai n = 8, maka operasi primitif akan berjalan sebanyak 3 kali. Biasanya dalam O(log n) ini akan ditemui pada binary search yang menggunakan algoritma divide and conquer yang akan membagi 2 jumlah inputan secara terus-menerus sampai mendapatkan nilai yang kita cari
5. Quadratic Time - O(n^2): Operasi primitif akan berjalan sebanyak nilai n pangkat 2, misalnya n = 10, maka operasi primitif tersebut akan berjalan 100 kali. Biasanya time complexity O(n^2) ini ditemui ketika digunakan sebuah loop di dalam loop (nested loop) dalam suatu algoritma

Selain jenis Time Complexity diatas, terdapat juga factorial time O(n!) dan exponential time O(2^n). Algoritma dengan time complexity tersebut sangat tidak efisien dan hanya dapat menyelesaikan masalah dengan jumlah n yang sangat sedikit karena akan memakan waktu yang terlalu lama apabila nilai n besar


Space Complexity: perhitungan sebuah algoritma mengenai seberapa banyak memory yang dibutuhkan algoritma tersebut menyelesaikan suatu permasalahan. Space Complexity berhubungan dengan sebuah variabel jadi misalnya pada function dominant diatas, Space Complexity adalah O(2) karena hanya terdapat 2 variabel, yaitu n dan result. Berbeda dengan sebuah function yang terdapat looping dan didalam looping tersebut append nilai ke dalam array maka space complexity function tersebut adalah O(n)