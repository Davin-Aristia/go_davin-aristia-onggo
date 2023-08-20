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
}