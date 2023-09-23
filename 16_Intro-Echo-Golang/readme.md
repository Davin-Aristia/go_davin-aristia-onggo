Third Party Library adalah sebuah kumpulan code yang memiliki fungsi-fungsi tertentu, yang dimana fungsi tersebut bisa dipanggil oleh program lain. Menggunakan library akan memudahkan dalam pembuatan sebuah projek, namun memiliki kekurangan/dampak negatif yaitu ketergantungan projek terhadap sebuah library. Apabila library memiliki update, maka projek juga harus diupdate agar sesuai dengan fungsi yang dijalankan. Beberapa library populer pada golang:
- echo
- Go Kit
- Logrus
- gorm.io
- cobra

Echo merupakan web framework golang yang memiliki performa tinggi, extensible, dan minimalis. Framework Echo cukup terkenal dan banyak dipakai oleh komunitas dan developer. Echo memiliki performa yang lebih cepat dibandingkan framework yang lain seperti Gin, dan Goji. Beberapa keuntungan menggunakan echo
- Optimized Router: penggunaan router echo yang cukup powerful karena adanya suatu optimasi berupa prioritize routing yang cerdas yang dilakukan oleh echo untuk menggunakan alokasi dynamic memory yang bisa bernilai 0
- Middleware: Di dalam echo terdapat banyak jenis middleware seperti JWT, dan Redirect HTTPS
- Data Rendering: Echo menggunakan custom render dalam memberikan sebuah respon dari sebuah request
- Scalable: Echo cukup mudah jika ingin dikembangkan pada REST API dalam skala kecil hingga besar
- Data Binding: ketika client melakukan sebuah request ke dalam REST API dengan mengirimkan query, form, json maka Echo akan cukup mudah dalam mendapatkan data tersebut dengan cara membuat struct dan binding data yang diterima ke dalam struct

Echo merupakan framework minimalis artinya framework echo cukup sederhana berukuran tidak besar, namun cukup powerful di fitur utamanya terutama di bagian REST API. Namun ketika ingin menambahkan database ORM, Echo belum bisa handle sehingga diperlukan untuk menggunakan libra tambahan. Echo juga tidak mepunyai sebuah struktur dalam pengembangannya, maka dari itu developer bisa memiliki kebebasan untuk memilih dengan konsep struktur seperti MVP, MVVM, Hexagonal, dan sebagainya. Walaupun minimalis, Echo memberikan kebebasan kepada developer dengan memberikan template engine yang dapat memudahkan developer untuk membangun REST API, template engine tersebut bisa diakses pada: https://echo.labstack.com/guide/templates

1. Contoh penggunaan Echo:
    func main(){
        e := echo.New()                 <- menginisiasi instansi echo
        e.GET("/", HelloController)     <- apabila endpoint yang digunakan adalah / dan method yang digunakan adalah GET, maka akan menjalankan fungsi HelloController
        e.Start(":8000")                <- menjalankan server pada port 8000
    }

    func Hello Controller (c echo.Context) error{
        return c.String(http.StatusOK, "Hello World")     <- memberikan respon berupa string Hello World dengan status 200 (OK)
    }

2. Render Data:
    type User struct{              <- define struct yang digunakan untuk mereturn response
        Name string
        Email string
    }
    func GetUser(c echo.Context) error{
        user := User{Name: "Davin", Email: "davin@gmail.com"}   <- Isi data pada struct
        return c.JSON(http.StatusOK, user)                      <- return struct
    }

3. Cara pengambilan URL Params
    e.GET("/users/:id, GetUserController)   <- routing yang di define untuk mengambil param id
    localhost:8000/users/3                  <- contoh request maka akan mendapatkan param id = 3
    id,_ := strconv.Atoi(c.Param("id"))      <- cara memanggil param id

4. Cara pengambilan Query Params
    e.GET("/users, UserSearchController)   <- routing yang di define untuk mengambil param id
    localhost:8000/users?match=a           <- contoh request maka akan mendapatkan query param match = 3
    match := c.QueryParam("match")         <- cara memanggil param id

5. Cara Pengambilan Form Value:
    name := c.FormValue("name")

6. Binding Data 
    type User struct{                               <- define struct untuk bind data ke dalam struct ini
        Name string `json:"name" form:"name"`       <- key data yang dikirim harus sesuai dengan variabel yang di define (dalam kasus ini name)
        Email string `json:"email" form:"email"`
    }
    
    user := User {}                                 <- deklarasi variabel struct
    c.Bind(&user)                                   <- Bind data yang diterima ke dalam struct