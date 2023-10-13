package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fileName:= "sample.txt"
	file, err :=os.Create(fileName)
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
	byteSlice,err:=os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error found while reading content: ", err)
		return
	}
	str:=string(byteSlice)
	fmt.Println("count is: ",countFile(strings.Fields(str)))

}

func countFile(content []string) int{
return len(content)
}