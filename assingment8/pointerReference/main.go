package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func updateAge(stu *Student) {
	stu.Age++
}

func main() {
	stu1 := Student{
		Name: "Vikalp Tyagi",
		Age:  23,
	}
	fmt.Println("Students age: ", stu1.Age)
	updateAge(&stu1)
	fmt.Println("Updated students age: ", stu1.Age)

}
