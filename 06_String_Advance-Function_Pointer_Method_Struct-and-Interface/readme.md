String merupakan salah satu tipe data pada golang yang dapat menyimpan teks. Di Golang terdapat beberapa fungsi string, seperti:
1. len(sentence): menghitung panjang sebuah string yang akan return angka dari panjang string tersebut
2. sentence1 == sentence2: compare nilai 2 buah nilai yang akan return nilai antara true (jika sama) dan false(jika beda)
3. strings.Contrains(sentence, word): mengecek apakah terdapat word di dalam sentence yang akan return nilai true dan false
4. sentence[5:len(value)]: mengambil string yang dimulai dari indeks ke 5
5. strinfs.Replace(word, "a","o", 3): mengganti huruf a yang terdapat pada word menjadi o sebanyak 3 huruf
6. word1 + word2: menggabungkan string

Function merupakan sekumpulan kode yang dikelompokkan sehingga dapat dipanggil dengan nama tertentu. Terdapat beberapa bentuk advanced function:
1. Variadic Function: digunakan ketika parameter yang digunakan untuk memanggil function tersebut berbeda-beda
Contoh: func sum (data ...int) int{//kode disini}
2. Anonymous function / Literal function: function yang tidak memiliki nama
Contoh: - anonymous := func(){//kode disini}
        - concurrency := go func(){//kode disini}
3. Closure: Merupakan sebuah penjabaran dari anonymous function, yang dimana setiap memanggil function ini bisa melanjutkan value dari variabel sebelumnya
Contoh:
func NewCounter() func() int{
    count := 0
    return func() int {
        count += 1
        return count
    }
}

counter := NewCounter()
fmt.Println(counter()) // 1
fmt.Println(counter()) // 2
4. Defer Function: Merupakan sebuah penjabaran dari anonymous function, yang dimana kelompok function yang di defer akan running diakhir
Contoh:
defer func(){
    fmt.Println("First")
}
defer fmt.Println("Second)
fmt.Println("Last")
// Last Second First


Pointer adalah sebuah variabel yang dapat menyimpan memory address dari variabel yang lain, cara deklarasi pointer adalah:
var p *int = &a

Struct berisi collection of named field/function, cara deklarasi struct:
type struct_variable_type struct{
    field <data_type>;
    field <data_type>;
    ...
    field <data_type>;
}

Method adalah function yang ditempel pada sebuah type (bisa berupa struck), cara deklarasi method:
func (receiver StructType) MethodName(parameterList) (returnTypes){//kode disini}