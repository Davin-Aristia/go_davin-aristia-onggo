package main

import (
	"fmt"
	"time"
	// "sync"
)

func multiple(x int){
	var res int
	for i:=0; i<5; i++{
		time.Sleep(3 * time.Second)
		res += x
		fmt.Println(res)
	}
}

func main(){
	go multiple(1)
	go multiple(10)

	time.Sleep(16 * time.Second)
}