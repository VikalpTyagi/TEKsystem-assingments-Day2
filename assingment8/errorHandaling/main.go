package main

import (
	"fmt"
	"os"
)

func main() {
	fileA := `C:\Users\ORR Training 17\Documents\day2&3\assingment8\errorHandaling\a.txt`
	fileB := `C:\Users\ORR Training 17\Documents\day2&3\assingment8\errorHandaling\b.txt`
	
	fmt.Println("Deleting 'a.txt' file\n",removeFile(fileA))
	fmt.Println("Deleting 'b.txt' file\n",removeFile(fileB))

}

func removeFile(fileName string) string {
	err := os.Remove(fileName)
	if err != nil {
		return err.Error()
	}
	return "Deleted sucessfully!!"
}
