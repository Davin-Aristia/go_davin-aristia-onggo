Middleware merupakan entitas yang terhubung ke pemrosesan request/response server. Middleware mengizinkan untuk mengimplementasikan berbagai fungsi seperti memasukkan permintaan http ke server dan response yang keluar. Beberapa contoh Third Party Middleware:
- Negroni
- Echo
- Interpose
- Alice

Setup Middleware Echo
1. Echo#Pre()
    Dieksekusi sebelum router memproses request:
    - HTTPSRedirect
    - HTTPSWWWRedirect
    - WWWRedirect
    - AddTrailingSlash
    - RemoveTralingSlash
    - Rewrite
2. Echo#Use()
    Dieksekusi setelah router memproses request dan memiliki akses penuh ke dalam echo.Context API
    - BodyLimit
    - Logger
    - Gzip
    - Recover
    - BasicAuth
    - JWTAuth
    - Secure
    - CORS
    - Static

CORS merupakan singkatan dari Cross Origin Resource Sharing. Memungkinkan resource sharing dengan berbagai origin/domain. Biasanya digunakan pada web application yang memiliki akses API dengan berbagai domain. Beberapa konfigurasi CORS yang sering digunakan:
- Access-Control-Allow-Origin: Untuk menentukan domain/origin yang dapat mengirim request ke server
- Access-Control-Allow-Headers: Untuk menentukan request header yang dapat digunakan ketika mengirim request ke server
- Access-Control-Allow-Methods: Untuk menentukan HTTP Method yang dapat digunakan ketika mengirim request ke server
- Access-Control-Allow-Age: Untuk menentukan durasi maksimum preflight request yang dapat dilakukan caching

Contoh implementasi CORS dengan konfigurasi:
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:8080", http://localhost:8081"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept}
        AllowMethods: []string{http.MethodGet, http.MethodPost}
    }))


Rate Limiter membatasi jumlah request kepada server. Mencegah request yang terlalu banyak untuk memastikan performance dari server tetap terjaga. Dan mencegah security breaches seperti DDos dan Brute Force Attack. Contoh implementasi Rate Limiter:
- e.Use(middleware.RateLimiter(middleware.NewrateLimiterMemoryStore(20)))
- e.Use(middleware.RateLimiter(middleware.NewrateLimiterMemoryStoreWithConfig(config)))

Logger Middleware melakukan logging informasi HTTP Request, sebagai foorprint, dan submer data untuk analytics. Contoh implementasi Log Middleware:
    e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
    }))

Auth Middleware digunakan untuk mengidentifikasi user dan mengamankan data dan informasi. Beberapa authentication middleware adalah:
1. Basic Authentication: merupakan teknik request http authentication, metode ini memerlukan informasi username dan password yang dimasukkan pada request header. Format pada Basic Authentication adalah: 'Authorization: Basic ' + base64encode('username:password')
2. JWT Middleware: Format pada JWT Middleware adalah: 'Authorization: Bearer ' + Token