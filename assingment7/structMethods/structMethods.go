package main

import "fmt"

type Rectange struct {
	width  int
	height int
}

func main() {
	rec1 := Rectange{
		width:  24,
		height: 58,
	}
	fmt.Println("Area of rectangle: ", rec1.cal())
}

func (r Rectange) cal() int {
	return 2 * (r.width + r.height)
}
