package main

import (
	"fmt"
	"time"
)

func multiple(x int){
	for i:=0; i<5; i++{
		time.Sleep(3 * time.Second)
		x += x
		fmt.Println(x)
	}
}

func main(){
	go multiple(1)
	multiple (100)
}