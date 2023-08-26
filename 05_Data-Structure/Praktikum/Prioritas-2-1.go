package main


import "fmt"


func PairSum(arr []int, target int) []int {
	var res []int
	for index, value := range arr{
		for i:=index+1; i<len(arr); i++{
			if value + arr[i] == target{
				res = []int{index, i}
				return res
			}		
		}
	}
	return res
}


func main() {

	// Test cases

	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]

}

