package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeNumber struct{
	val int
	m sync.Mutex
}

func (i *SafeNumber) Get() int{
	i.m.Lock()
	defer i.m.Unlock()
	return i.val
}

func (i *SafeNumber) Set(val int){
	i.m.Lock()
	defer i.m.Unlock()
	res := val
	for i:=1; i<val; i++{
		res*=i
	}
	i.val = res
}

func factorial(num int) int{
	i:= &SafeNumber{}

	go func(){
		i.Set(num)
	}()
	time.Sleep(time.Second)
	return i.Get()
}

func main(){
	fmt.Println(factorial(5))
	fmt.Println(factorial(6))
	fmt.Println(factorial(7))
}