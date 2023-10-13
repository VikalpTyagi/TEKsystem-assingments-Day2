package main

import "fmt"

type Employee struct {
	Name string
	Age  int
	City string
}

func main() {
	emp1 := Employee{
		Name: "Vikalp",
		Age:  24,
		City: "Kota",
	}
	emp2 := Employee{
		Name: "Vishal",
		Age:  27,
		City: "Patna",
	}

	fmt.Printf("Employee1 :%#v\n", emp1)
	fmt.Printf("Employee2: %#v\n", emp2)
}
