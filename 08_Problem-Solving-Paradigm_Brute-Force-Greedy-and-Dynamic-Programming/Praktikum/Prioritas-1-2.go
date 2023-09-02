package main

import "fmt"

func Pascal(numRows int) [][]int{
	var res [][]int
	var temp []int

	for i:=1; i<=numRows; i++{
		value := 1
		temp = nil
		for j:=1; j<=i; j++{
			temp = append(temp, value)
			value = value*(i-j)/j
		}
		res = append(res,temp)
	}
	return res
}

func main() {
  	fmt.Println(Pascal(5))
  	fmt.Println(Pascal(6))
  	fmt.Println(Pascal(7))
}