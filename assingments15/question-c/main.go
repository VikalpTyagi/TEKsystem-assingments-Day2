package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)
func main() {
	go sender(10)
	go sender(20)
	go sender(30)
	go reciver()
	go reciver()
	time.Sleep(1*time.Second)
	fmt.Println("End of main func")
	
}
func sender(num int){
ch<- num
}
func reciver(){
 res:= <- ch
 fmt.Println("recieved from channel",res)
}