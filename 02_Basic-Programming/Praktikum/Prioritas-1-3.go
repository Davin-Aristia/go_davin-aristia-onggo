package main

import "fmt"

func main(){
	//Grade
	var grade int
	fmt.Printf("Masukkan Nilai: ")
	fmt.Scan(&grade)
	switch true{
		case grade >= 80 && grade <= 100: fmt.Println("A")
		case grade >= 65 && grade <= 79: fmt.Println("B")
		case grade >= 50 && grade <= 64: fmt.Println("C")
		case grade >= 35 && grade <= 49: fmt.Println("D")
		case grade >= 0 && grade <= 34: fmt.Println("E")
		default: fmt.Println("Nilai Invalid")
	}
}