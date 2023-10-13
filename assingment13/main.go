package main

import (
	"fmt"
	// "io/ioutil"
	"os"
)

func main() {
	file, err :=os.Create("a.txt")
	if err != nil {
			fmt.Println("Error found while creation: ", err)
			return
		}
	
	something,err:= file.WriteString("Hello this file is created by Vikalp")
	if err != nil {
		fmt.Println("Error found while writing: ", err)
		return
	}
	//This will tell us how many spaces we have use
	// fmt.Println(something)
	_= something

}

func 
