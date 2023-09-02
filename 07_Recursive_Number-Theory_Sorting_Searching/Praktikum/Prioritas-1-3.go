package main


import (
	"fmt"
	"math"
)


func primeX(number int) int {
	counter,res := 1,3
	for counter < number{
		for i := 2; i <= int(math.Ceil(math.Sqrt(float64(res)))); i++{
			if res % i == 0{
				break
			}
			if i == int(math.Ceil(math.Sqrt(float64(res)))){
				counter++
			}
		}
		res++
	}
	return res-1
}


func main() {

fmt.Println(primeX(1)) // 2

fmt.Println(primeX(5)) // 11

fmt.Println(primeX(8)) // 19

fmt.Println(primeX(9)) // 23

fmt.Println(primeX(10)) // 29

}