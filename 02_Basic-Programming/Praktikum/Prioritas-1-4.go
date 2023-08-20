package main

import "fmt"

func main(){
	//Fizz Buzz
	for i:=1; i<=100;i++{
		if i%15 == 0{
			fmt.Printf("FizzBuzz ")
		}else if i%3 == 0{
			fmt.Printf("Fizz ")
		}else if i%5 == 0{
			fmt.Printf("Buzz ")
		}else{
			fmt.Print(i," ")
		}
	}
}