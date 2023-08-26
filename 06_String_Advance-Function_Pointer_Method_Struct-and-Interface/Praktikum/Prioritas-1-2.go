package main

import (
	"fmt"
	"strconv"
)

type Student struct{
    name []string;
    score []int;
}

func (s Student) Score() string{
	var sum, min, max int = 0, s.score[0], s.score[0]
	var nameMin, nameMax string = s.name[0], s.name[0]
	for index,value := range s.score{
		sum += value
		if value < min{
			min = value
			nameMin = s.name[index]
		}
		if value > max{
			max = value
			nameMax = s.name[index]
		}
	}
	fmt.Println(sum)
	average := float64(sum) / 5
	return "Average Score: " + strconv.FormatFloat(average, 'f', 1, 64) + "\nMin Score of Students: " + nameMin + " (" + strconv.Itoa(min) + ")\nMax Score of Students: " + nameMax + " (" + strconv.Itoa(max) + ")"
}

func main(){
	s := Student{ 
		name: []string{"Rizky","Kobar","Ismail","Umam","Yopan"},
		score: []int{80,75,70,75,60},
	}
	fmt.Println(s.Score())
}