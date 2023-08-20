package main

import "fmt"

func main(){
	//Odd Even
	var OddEven int
	fmt.Printf("Masukkan Nilai Ganjil/Genap: ")
	fmt.Scan(&OddEven)
	if OddEven%2 == 0{
		fmt.Println("Bilangan Genap")
	}else {
		fmt.Println("Bilangan Ganjil")
	}
}