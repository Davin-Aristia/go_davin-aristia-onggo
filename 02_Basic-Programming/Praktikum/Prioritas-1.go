package main

import "fmt"

func main(){
	//Luas Trapesium
	var alas,atap,tinggi int
	fmt.Printf("Masukkan Nilai Alas Trapesium: ")
	fmt.Scan(&alas)
	fmt.Printf("Masukkan Nilai Atap Trapesium: ")
	fmt.Scan(&atap)
	fmt.Printf("Masukkan Nilai Tinggi Trapesium: ")
	fmt.Scan(&tinggi)

	Luas := (alas+atap)/2 * tinggi
	fmt.Println("Luas Trapesium:", Luas)

	//Odd Even
	var OddEven int
	fmt.Printf("Masukkan Nilai Ganjil/Genap: ")
	fmt.Scan(&OddEven)
	if OddEven%2 == 0{
		fmt.Println("Bilangan Genap")
	}else {
		fmt.Println("Bilangan Ganjil")
	}

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

	//Fizz Buzz
	for i:=1; i<=100;i++{
		if i%15 == 0{
			fmt.Printf("FizzBuzz ")
		}else if i%3 == 0{
			fmt.Printf("Fizz ")
		}else if i%5 == 0{
			fmt.Printf("Buzz ")
		}else{
			fmt.Print(i," ")
		}
	}
}