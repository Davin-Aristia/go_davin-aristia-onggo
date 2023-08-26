package main


import "fmt"


func getMinMax(numbers ...*int) (min, max int) {
	min, max = *numbers[0], *numbers[0]
	for _, number := range numbers{
		if *number > max{
			max = *number
		}
		if *number < min{
			min = *number
		}
	}
	return min,max
}


func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	fmt.Println()

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println(min, "is the minimum number")
	fmt.Println(max, "is the maximum number")
}