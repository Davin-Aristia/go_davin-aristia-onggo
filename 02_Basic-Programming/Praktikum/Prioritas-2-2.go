package main

import "fmt"

func main(){
	//Bilangan Faktor
	var factor int
	fmt.Printf("Masukkan Nilai: ")
	fmt.Scan(&factor)
	for i:=1; i<=factor;i++{
		if factor%i == 0{
			fmt.Println(i)
		}
	}
}