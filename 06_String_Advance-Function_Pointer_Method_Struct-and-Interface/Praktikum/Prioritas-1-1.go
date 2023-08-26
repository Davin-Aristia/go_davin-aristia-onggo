package main

import (
	"fmt"
	"strconv"
)

type Car struct{
    carType string;
    fuelin float64;
}

func (c Car) DistanceEstimation() string{
	distance := c.fuelin/1.5
	return "car type:"+" "+c.carType+" , "+"est. distance: "+strconv.FormatFloat(distance, 'f', -1, 64)
}

func main(){
	c := Car{
		carType: "SUV",
		fuelin: 10.5,
	}
	fmt.Println(c.DistanceEstimation())
}