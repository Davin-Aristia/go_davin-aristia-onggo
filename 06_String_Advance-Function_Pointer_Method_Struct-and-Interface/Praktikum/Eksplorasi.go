package main

import "fmt"

type student struct {
	name string
	nameEncode string
	score int
}


type Chiper interface {
	Encode() string
	Decode() string
}


func (s *student) Encode() string {
  var nameEncode = ""
  
  	for i:=0; i<len(s.name); i++{
		nameEncode += string(122-(s.name[i]-97))
	}
  
  return nameEncode
}


func (s *student) Decode() string {
	var nameDecode = ""
	
	for i:=0; i<len(s.name); i++{
		nameDecode += string(122-(s.name[i]-97))
	}
	
	return nameDecode

}


func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a
	
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)
	
	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student's name " + a.name + "is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student's name " + a.name + "is : " + c.Decode())
	}
}