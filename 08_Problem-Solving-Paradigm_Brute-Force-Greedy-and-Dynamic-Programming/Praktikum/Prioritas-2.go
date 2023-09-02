package main


import (
	"fmt"
	"math"
)


func Frog(jumps []int) int {
	var dp []int

	for i:=0; i<len(jumps); i++{
		dp = append(dp,-1)
	}
	dp[0] = 0
	for i:=1; i<len(jumps);i++{
		jumpTwo := math.MaxInt64
		jumpOne := dp[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))
		if i > 1{
			jumpTwo = dp[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))
		}
		dp[i] = int(math.Min(float64(jumpTwo),float64(jumpOne)))
	}
	return dp[len(jumps)-1]
}


func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}