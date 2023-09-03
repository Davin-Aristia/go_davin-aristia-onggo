package main

import (
	"fmt"
	"sync"
	"math"
)

func multiple3(n int){
	c := make(chan int, n/3)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		defer wg.Done()
		for i:=3; i<=n; i+=3{
			c <- i
		}
	}()

	wg.Add(1)
	go func(){
		defer wg.Done()
		for i:=0; i<int(math.Floor(float64(n/3))); i++{
			fmt.Println(<-c)
		}
	}()
	wg.Wait()
}

func main(){
	multiple3(31)
}