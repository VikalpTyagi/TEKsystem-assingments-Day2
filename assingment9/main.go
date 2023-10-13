package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Area()
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
func (r Rectangle) Area() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	circle := Circle{
		Radius: 2.5,
	}
	rec1 := Rectangle{
		Width:  24.5,
		Height: 56,
	}

	fmt.Println("Area of Circle: ", circle.Area())
	fmt.Println("Area of Rectangle: ", rec1.Area())
}
