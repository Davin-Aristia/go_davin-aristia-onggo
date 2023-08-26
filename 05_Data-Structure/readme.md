Array adalah struktur data yang berisi grup dari sebuah elemen, array biasanya hanya dapat berisi satu tipe variabel dan ukuran dari array tidak bisa lebih dari yang telah dideklarasikan sebelumnya. Tipe data array dapat berupa Numeric, String, dan Boolean

Cara deklarasi array hampir sama dengan deklarasi variabel, bedanya adalah terdapat kurung siku tepat disebelum tipe data yang berisi angka yang menentukan ukuran dari array tersebut, misalnya adalah: var countries [5]string

Terdapat 2 cara untuk mendeklarasikan array sekaligus mengisi nilai dari array tersebut:
1. x := [5]int{1, 3, 5, 7, 9}
2. var x [5]int = [5]int{1, 3, 5, 7, 9}

Terdapat 3 cara untuk looping array:
1. for index := 0; index < len(x); index++ : dimana kita bisa mendapatkan nilai dari array dengan cara x[index]
2. for index, element := range x: dimana index akan menunjukkan index dari array tersebut yang akan mulai dari angka 0, dan element akan menunjukkan nilai dari array tersebut
3. for range x: dimana kita dapat mendapatkan nilai dari array dengan cara x[index] dan index tersebut di increment (index++)

Slice hampir mirip dengan array yaitu struktur data yang berisi grup dari sebuah elemen dan hanya bisa berisi satu tipe variabel, bedanya adalah ukuran dari slice ini dapat diubah. Selain itu, slice memiliki sebuah pointer yang merupakan sebuah referensi array yang diambil. Cara deklarasi slice juga mirip dengan dengan array hanya saja kurung siku yang terdapat pada array tidak diisi oleh angka melainkan kosong

Contoh penggunaan slice:
1. var colors = []string{"red", "green", "yellow"}
2. colors = append(colors, "purple")
3. copied_colors := make([]string, 10)
4. copy(copied_colors, colors)
Penjelasan:
1. Pada baris pertama adalah deklarasi sekaligus menambahkan nilai pada slice
2. Pada baris kedua ada menambah nilai "purple" pada slice
3. Pada baris ketiga adalah cara lain mendeklarasikan slice
4. Pada baris keempat adalah copy nilai pada slice colors ke copied colors

Map juga merupakan struktur data pada golang, namun pada map terdapat pasangan key dan value dimana nilai pada setiap key bersifat unik. Cara mendeklarasikan map adalah:
var harga = map[string]int{"siomay": 1000, "batagor": 2000}

Function adalah sekumpulan kode yang dikelompokkan sehingga dapat dipanggil dengan nama tertentu. Tujuan dari sebuah function adalah membuat clean code sehingga lebih mudah untuk di maintenance dan lebih mudah membuat unit testing

Contoh dari function:
func calculateCircle(diameter float64) (float64, float64){
    var keliling = math.Pi * math.Pow(diameter/2, 2)
    var luas = math.Pi * diameter
    return keliling, luas
}

Lalu, function tersebut bisa dipanggil dengan cara:
keliling, luas := calculateCircle(diameter)