package main


import "fmt"


func MaxSequence(arr []int) int {
	var res, add int
	for index,_ := range arr{
		for i:=len(arr); i>=index+1;i--{
			add = 0
			tempSlice := arr[index:i]
			for _,value := range tempSlice{
				add+=value
			}
			if add > res{
				res = add
			}
		}
	}
	return res
}


func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12
}

