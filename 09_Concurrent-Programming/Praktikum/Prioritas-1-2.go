package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan int)

	go func(){
		for i:=3; i<=30; i+=3{
			c <- i
		}
	}()

	go func(){
		for i:=0; i<10; i++{
			fmt.Println(<-c)
		}
	}()
	time.Sleep(10 * time.Millisecond)
}