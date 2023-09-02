package main


import (
	"fmt"
)


func Rome(num int) string {
	dictRome := []struct {
        roman string
        number int
    }{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
	for _,value := range dictRome{
		if num > value.number{
			return value.roman + Rome(num - value.number)
		} else if num == value.number{
			return value.roman
		}
	}
	return "Missing Value"
}


func main() {
	fmt.Println(Rome(4))
	fmt.Println(Rome(9))
	fmt.Println(Rome(23))
	fmt.Println(Rome(2021))
	fmt.Println(Rome(1646))
}

