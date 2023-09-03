package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "encoding/json"
    "io"
    "sync"
)

type Response struct{
    Title string `json:"title"`
    Price float64 `json:"price"`
    Category string `json:"category"`
}

func main() {
    responseChannel := make(chan Response)
    var responseObject []Response
    var wg sync.WaitGroup
    var apiWg sync.WaitGroup

    // HTTP request ke API
    response, err := http.Get("https://fakestoreapi.com/products")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }
    defer response.Body.Close()

    // cek response status
    if response.StatusCode != http.StatusOK {
        fmt.Println("API returned non-OK status code:", response.Status)
        return
    }

    // menggunakan goroutine untuk menyimpan data yang didapatkan ke variabel responseObject
    apiWg.Add(1)
    go func() {
        defer apiWg.Done()
        responseData, err := io.ReadAll(response.Body)
        if err != nil {
            log.Fatal(err)
        }
        json.Unmarshal(responseData, &responseObject)
    }()
    
    // menggunakan goroutine untuk mengirim data dari responseObject ke responseChannel
    wg.Add(1)
    go func() {
        defer wg.Done()
        apiWg.Wait()
        for _, value := range responseObject {
            responseChannel <- value
        }
        close(responseChannel)
    }()

    // menutup channel saat semua proses goroutine selesai dijalankan
    go func() {
        wg.Wait()
    }()

    fmt.Println("products data")
    for value := range responseChannel{
        fmt.Println("===")
        fmt.Println("title: ", value.Title)
        fmt.Println("price: ", value.Price)
        fmt.Println("category: ", value.Category)
        fmt.Println("===")
    }
}
