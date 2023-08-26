package main

import (
	"fmt"
	"math"
)

func main(){
	//Bilangan Prima
	var number int
	fmt.Printf("Masukkan Angka: ")
	fmt.Scan(&number)
	fmt.Println(primeNumber(number))
}

func primeNumber(number int) bool {
	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(number)))); i++{
			if number % i == 0{
				return false
			}
		}
	return true
}