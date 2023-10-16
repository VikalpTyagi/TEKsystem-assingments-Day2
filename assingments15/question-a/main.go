package main

import (
	"fmt"
	"time"
)


func main() {
	ch := make(chan int)
	go sender(ch)
	go reciever(ch)
	time.Sleep(4*time.Second)
	fmt.Println("end of main function")
}

func sender(ch chan int) {
	for i := 1; i < 6; i++ {
		ch <- i
	}
}
func reciever(ch chan int) {
	for chanl := range ch {
		fmt.Println("value recieved: ", chanl)
	}
}
