Object Relational Mapping (ORM, O/RM, dan O/R mapping tool) dalam computer science merupakan sebuah teknik pemrograman untuk konversi data antara sistem tipe yang tidak kompatibel menggunakan bahasa Object-Oriented Programming (OOP)

GORM merupakan library ORM untuk Golang. Beberapa keunggulan menggunakan ORM:
- Mengurangi query yang berulang-ulang
- Mengotomasi pengambilan data menjadi object yang siap pakai
- Merupakan cara yang simpel apabila ingin melakukan screening data sebelum disimpan ke dalam database
- Beberapa memiliki fitur cache query

Beberapa kerugian dalam penggunaan ORM:
- Menambah lapisan dalam code dan menambah biaya dalam proses overhead
- Load unnecessary relationship data
- Query mentah yang kompleks akan menjadi panjang dengan ORM (> 10 join tabel)
- Fungsi sql tertentu yang terkait dengan satu vendor mungkin tidak didukung atau tidak ada fungsi tertentu (MYSQL: FORCE INDEX)

Database Migration merupakan sebuah cara untuk update versi database untuk memenuhi perubahan versi aplikasi. Perubahan bisa dengan upgrade ke versi terbaru ataupun rollback ke versi sebelumnya. Alasan mengapa menggunakan Database Migration:
- Kesederhanaan dalam melakukan update database
- Kesederhanaan dalam melakukan rollback database
- Melacak perubahan dalam struktur database
- Riwayat struktur basis data yang ditulis dalam code
- Selalu kompatibel dengan perubahan versi aplikasi

Relasi Database dalam GORM:
- Belongs To: Cocok untuk relasi yang terhubung satu sama lain
    Contoh: user belongs to company
- Has One: Cocok untuk merepresentasi sebuah entitas memiliki hanya satu entitas
    Contoh: user memiliki satu kartu kredit
- Has Many: Cocok untuk merepresentasi sebuah entitas memiliki banyak entitas
    Contoh: user memiliki banyak kartu kredit
- Many to Many: Cocok untuk merepresentasi banyak entitas memiliki banyak entitas
    Contoh: banyak user bisa berbicara banyak bahasa, dan banyak bahasa bisa diucapkan oleh banyak user

Database Transaction in GORM:
- Transaksi merupakan sequence dari operasi database. Jika satu operasi gagal, maka semua transaksi tersebut akan gagal juga
- Transaksi berguna untuk memastikan integritas dan konsistensi dari data
- Transaksi didukung dalam GORM

Cara Koneksi aplikasi dengan database:
1. Membuat fungsi InitDB() untuk koneksi database:
    func InitDB(){

        // deklarasi struct config & variabel connectionString

        var err error
        DB, err = gorm.Open("mysql", connectionString)
        if err != nil{
            panic(err)
        }
    }

2. Membuat User Schema & InitialMigration()
    var (
        DB *gorm.DB
    )

    type User struct {              <- Membuat User Schema untuk database
        gorm.Model
        Name     string `json:"name" form:"name"`
        Email    string `json:"email" form:"email"`
        Password string `json:"password" form:"password"`
    }

    func InitialMigration(){        <- Membuat Database pada mysql
        DB.AutoMigrate(&User{})
    }

3. Memanggil InitDB() dan InitialMigration
    package main

    // import library

    func init(){
        InitDB()
        InitialMigration
    }

4. Membuat GetUsersController()
    func GetUsersController(c echo.Context) error{
        var users []User

        if err := DB.Find(&users).Error; err != nil{
            return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }
        return c.JSON(http.StatusOK, map[string]interface{}{
            "message": "success get all users",
            "users": users,
        })
    }

5. Membuat CreaterUserController()
    func CreateUserController(c echo.Context) error {
        user := User{}
        c.Bind(&user)

        if err := DB.Save(&user).Error; err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, err.Error())
        }
        return c.JSON(http.StatusOK, map[string]interface{}{
            "message": "success create new user",
            "user": user,
        })
    }

6. Routing
    func main() {
        // create a new echo instance
        e := echo.New()

        // Route / to handler function
        e.GET("/users", GetUsersController)
        e.POST("/users", CreateUserController)

        // start the server, and log if it fails
        e.Start(":8000")
    }


MVC merupakan akronim dari Model, View, dan Controller. MVC merupakan code structure yang cukup populer untuk mengorganisasi code. Ide utama dari MVC ini adalah setiap bagian dari code memiliki tujuan, dan setiap tujuan tersebut berbeda-beda. Alasan penggunaan Code Structure ada;ah:
- Untuk mencapai aplikasi modular
- Implementasi separation of concerns
- Lebih sedikit konflik dalam versioning