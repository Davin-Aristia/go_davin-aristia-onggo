package main


import (
	"fmt"
	"math"
)


func absoluteMatrix(arr [][]int) int {
	var res int = 0
	for i:=0; i<len(arr); i++{
		res += arr[i][i]
		res -= arr[i][len(arr)-1-i]
	}
	return int(math.Abs(float64(res)))
}


func main() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	fmt.Println(absoluteMatrix(a))
}

