package main


import "fmt"


func caesar(offset int, input string) string {
	res := ""
	add := offset % 26
	for i:=0; i<len(input); i++{
		if input[i]+byte(add) > 122{
			res += string(input[i]+byte(add)-26)
		}else{
			res += string(input[i]+byte(add))
		}
	}
	return res
}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // def
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}