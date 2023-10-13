package main

import "fmt"

func main() {
	fmt.Println("Sum of 44,56,22,87,67: ", Addition(44, 56, 22, 87, 67))
	Addition(44, 56, 22, 87, 67)
}

func Addition(data ...int) int {
	sum := 0
	for _, num := range data {
		sum += num
	}
	return sum
}
