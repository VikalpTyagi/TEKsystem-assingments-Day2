package main

import "fmt"

type Cat struct {
	Name  string
	Sound string
}
type Dog struct {
	Name  string
	Sound string
}

type Animal interface {
	MakeSound()
}

func (c Cat) MakeSound() string {
	return c.Sound
}
func (d Dog) MakeSound() string {
	return d.Sound
}

func main() {
	cat := Cat{
		Name:  "Jack",
		Sound: "Meow",
	}
	dog := Dog{
		Name:  "Tyson",
		Sound: "Woof",
	}
	fmt.Println("Sound made by cat: ", cat.MakeSound())
	fmt.Println("Sound made by dog: ", dog.MakeSound())
}
