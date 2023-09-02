package main


import "fmt"

var storage = map[int]int{}
func fibonacci(number int) int {
	value, exist := storage[number]
	if exist{
		return value
	}

	if number <= 1{
		storage[number] = number
	} else {
		storage[number] = fibonacci(number-1) + fibonacci(number-2)
	}
	return storage[number]
}


func main() {
	fmt.Println(fibonacci(0)) // 0
	fmt.Println(fibonacci(2)) // 1
	fmt.Println(fibonacci(9)) // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}