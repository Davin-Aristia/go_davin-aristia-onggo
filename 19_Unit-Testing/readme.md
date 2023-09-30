Software Testing merupakan proses untuk menganalisis sebuah software item untuk mendeteksi perbedaan antara kondisi yang sudah ada dan kondisi yang dibutuhkan, selain itu juga untuk mengevaluasi fitur dari sebuah software item

Tujuan dari dilakukan testing:
- Mencegah regresi
- Tingkat kepercayaan dalam refactoring
- Improve Code Design
- Code Documentation
- Scaling the Team

Level pada Testing:
- UI: (End To End) uji interaksi antara keseluruhan melalui User Interface
- Integration: Menguji modul tertentu atau sub-system melalui API
- Unit: Menguji bagian terkecil yang dapat diuji (single logical operation) dari sebuah aplikasi melalui metode

Berdasarkan bahasa pemrograman, harus dipilih unit testing framework. Framework menyediakan alat dan struktur yang dibutuhkan untuk menjalankan testing secara efisien. Beberapa unit testing framework dapat dilihat pada: https://en.wikipedia.org/wiki/List_of_unit_testing_frameworks

Struktur Unit Testing:
- Sentralisasi test file di dalam test folder
- Menyimpan test file bersama-sama dengan production file

Runner merupakan alat untuk menjalankan test file:
- Menggunakan watch mode: jika ada perubahan dalam codebase, secara otomatis menjalankan run test dan membuat Test Driven Development lebih efisien
- Pilihlah runner yang paling cepat

Mocking: Sebuah test case harus bersifat independen, dan perlu membuat mock object (objek palsu yang mensimulasikan sifat dari objek asli)

Coverage: Programmer perlu memastikan apakah mereka sudah membuat test yang cukup. Coverage tool menganalisis kode aplikasi ketika test berjalan

Beberapa step simple untuk test:
- Membuat test file baru pada Go:
    1. Nama file diakhiri dengan library_test.go (Contoh: user_test.go)
    2. Lokasi file:
        a. Folder yang sama, dan package yang sama
        b. Folder yang sama, package yang berbeda
- Menulis fungsi test:
    1. Nama fungsi: Test & Huruf Kapital
    2. Harus memiliki fungsi func TestXxx(t *testing.T)
- Run Testing: go test ./lib/calculate -cover
- Run with Report Coverage: go test ./lib/calculate -coverprofile=cover.out && go tool cover -html=cover.out
- Menulis fungsi test untuk echo:
    1. Setup Echo (Method, Params, dsb)
    2. Memanggil fungsi rest api di package controller
    3. Membuat fungsi test
- Run with Report Coverage All Package:
    go test -v -coverpkg=./... -coverprofile=profile.cov ./...
    go tool cover -func profile.cov


Contoh fungsi test untuk echo:
package controllers

func InitEcho() *echo.Echo{
    config.InitDB()
    e := echo.New()

    return e
}

func TestGetUserControllers(t *testing.T) {
    var testCases = []struct{
        name                 string
        path                 string
        expectStatus         int
        expectBodyStartsWith string
    }{
        {
            name:                 "berhasil",
            path:                 "/users",
            expectBodyStartsWith: "{\"status\":\"success\",\"users\":[",
            expectStatus:         http.StatusOK,
        },
    }

    e := InitEcho()
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    for _, testCase := range testCases{
        c.SetPath(testCase.path)

        // Assertions
        if assert.NoError(t, GetUserControllers(c)){
            assert.Equal(t, http.StatusOK, rec.Code)
            body := rec.Body.String()
            // asset.Equal(t, userJSON, rec.Body.String())
            assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
        }
    }
}