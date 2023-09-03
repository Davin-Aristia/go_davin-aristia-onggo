package main

import (
	"fmt"
	"sync"
	"strconv"
)

func textSplit(text string, split int) []string{
	var res []string
	length := len(text) / split
	start := 0

	for i:=0; i<split;i++{
		if i==split-1{
			res = append(res ,text[start:])
		}else{
			end := start+length
			res = append(res ,text[start:end])
			start = end
		}
	}
	return res
}

func letterCount(text string) map[rune]int{
	res := make(map[rune]int)

	for _,value := range text{
		res[value]++
	}
	return res
}

func calculateFrequency(text string) string{
	// kata yang dibagi menjadi n bagian, dalam kasus ini dibagi menjadi 4
	split := 4
	letterChannel := make(chan map[rune]int, split)
	var wg sync.WaitGroup

	chunks := textSplit(text, split)
	for _,value := range chunks{
		wg.Add(1)
		go func(text string){
			defer wg.Done()
			letterChannel <- letterCount(text)
		}(value)
	}

	go func() {
        wg.Wait()
        close(letterChannel)
    }()

	mapRes := make(map[rune]int)
	for value := range letterChannel{
		for key,val := range value{
			mapRes[key] += val
		}
	}

	var res string
	for key,value := range mapRes{
		res += strconv.QuoteRune(key) + ": "+strconv.Itoa(value) + "\n"
	}
	return res
}

func main(){
	fmt.Println(calculateFrequency("helloworld"))
	fmt.Println(calculateFrequency("ILoveGolang"))
}