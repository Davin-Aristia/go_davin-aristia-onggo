package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	//Palindrome
	fmt.Printf("masukkan kata: ")

  	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := scanner.Text()
	fmt.Println("captured:",words)
	
	if Palindrome(words){
		fmt.Println("Palindrome")
	}else{
		fmt.Println("Bukan Palindrome")
	}
}

func Palindrome(word string) bool {
	for i := 0; i<=len(word)/2; i++{
		if word[len(word)-i-1] != word[i]{
			return false
		}
	}
	return true
}