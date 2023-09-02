package main

import (
    "fmt"
	"io"
    "log"
    "net/http"
    "os"
    "encoding/json"
	"reflect"
)

func main() {
    response, err := http.Get("https://fakestoreapi.com/products")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    data,_ := json.Marshal(string(responseData))
    // fmt.Println(reflect.TypeOf(string(responseData)))
    // fmt.Println(responseData)
    // fmt.Println(string(responseData))
    fmt.Println(reflect.TypeOf(string(data)))

}
