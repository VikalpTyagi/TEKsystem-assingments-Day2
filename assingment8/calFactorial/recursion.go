package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number to find Factorial of: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Error FOund: ", err)
		return
	}
	fmt.Println("Factorial of given number is: ", findFactorial(num))
}

func findFactorial(num int) int {
	if num == 1 {
		return num
	}
	return findFactorial(num-1) * num
}
