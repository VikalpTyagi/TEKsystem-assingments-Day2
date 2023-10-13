package main

import (
	"fmt"
	"moduleAss5/model"
	"moduleAss5/shape"
)

func main() {
	circle1 := model.Circle{
		Radius: 2.5,
	}
	fmt.Printf("Radius: %v\nArea: %v ",circle1.Radius,shape.AreaCal(circle1.Radius))
}