package main


import (
	"fmt"
	"sort"
)


type pair struct {
	name string
	count int
}

func sortMap(data map[string]int) []string{
	keys := make([]string, 0, len(data))
  
	for key := range data {
		keys = append(keys, key)
	}
  
	sort.SliceStable(keys, func(i, j int) bool{
		return data[keys[i]] < data[keys[j]]
	})
	
	return keys
  }

func MostAppearItem(items []string) []pair {
	var res []pair
	var mapResult = map[string]int{}
	for _, value := range items{
		mapResult[value]++
	}
	for _,value := range sortMap(mapResult){
		res = append(res, pair{
			name : value,
			count: mapResult[value],
		})
	}
	return res
}


func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}

}