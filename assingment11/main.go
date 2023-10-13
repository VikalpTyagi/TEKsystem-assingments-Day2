package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode int
}

type Person struct {
	Name string
	Address
}

func main() {
	per1 := Person{
		Name: "Vikalp Tyagi",
		Address: Address{
			Street:  "Mahaveer Nagar Third",
			City:    "Kota",
			ZipCode: 324005,
		},
	}
	fmt.Printf("Data of person: %v",per1)
}
