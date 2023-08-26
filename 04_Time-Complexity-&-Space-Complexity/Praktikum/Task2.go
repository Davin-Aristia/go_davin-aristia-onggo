package main

import (
	"fmt"
)

func main(){
	//Pangkat
	var x, n int
	fmt.Printf("Masukkan Angka: ")
	fmt.Scan(&x)
	fmt.Printf("Masukkan Pangkat: ")
	fmt.Scan(&n)
	fmt.Println(pow(x, n))
}

func pow(x, n int) int {
	res := 1

	for n > 0{
		if n % 2 == 1{
			res = res * x
		}
		n = n / 2
		x = x * x
	}
	return res
}