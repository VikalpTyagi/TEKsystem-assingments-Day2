package main

import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}
var ch = make(chan int)
func main() {
	wg.Add(2)
	go sender(10)
	go reciver()
	wg.Wait()
	fmt.Println("End of main func")
	
}
func sender(num int){
	defer wg.Done()
for i := 0; i <= num; i++ {
	ch <-i
}
close(ch)
}
func reciver(){
	defer wg.Done()
 for res := range ch{
	fmt.Println("recieved from channel",res)
 }
}