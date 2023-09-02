package main


import (
	"fmt"
	"strconv"
)


func Binary(num int) []int {
	var res []int
	for i:=0; i<=num; i++{
		temp,_:=strconv.Atoi(fmt.Sprintf("%b", i))
		res = append(res,temp)
	}
	return res
}


func main() {
	fmt.Println(Binary(2))
	fmt.Println(Binary(3))
	fmt.Println(Binary(4))
	fmt.Println(Binary(5))
	fmt.Println(Binary(6))
}

