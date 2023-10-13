package person

import (
	"assing8/model"
	"fmt"
)

func PrintNameAge(person model.Person) {
	fmt.Printf("Name: %q\nAge: %d",person.Name,person.Age)
}
