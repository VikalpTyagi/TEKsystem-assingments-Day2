package main

import (
	"fmt"
	"time"
)


func main() {
	ch := make(chan int, 5)
	go sender(ch)
	go reciever(ch)
	time.Sleep(4*time.Second)
	fmt.Println("end of main function")
}

func sender(ch chan int) {
	time.Sleep(2*time.Second)
	for i := 1; i < 6; i++ {
		ch <- i
	}
}
func reciever(ch chan int) {
	for chanl := range ch {
		fmt.Println("value recieved: ", chanl)
	}
}
