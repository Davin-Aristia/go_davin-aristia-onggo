package main


import "fmt"


func munculSekali(angka string) []int {
    var res []int
    for index, value := range angka{
        for index2, value2 := range angka{
            if value == value2 && index != index2{
                break
            }
            if index2 == len(angka)-1{
                res = append(res, int(value)-48)
            }
        }
    }
    return res
}


func main() {
	fmt.Println(munculSekali("1234123")) // [4]
	fmt.Println(munculSekali("76523752")) // [6 3]
	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]
}

