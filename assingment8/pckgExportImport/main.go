package main

import (
	"assing8/model"
	"assing8/person"
)

func main() {
	per1 := model.Person{
		Name: "Vikalp Tyagi",
		Age: 24,
	}
	person.PrintNameAge(per1)
}