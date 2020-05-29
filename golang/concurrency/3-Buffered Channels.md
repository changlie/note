Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

ch := make(chan int, 100)
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Modify the example to overfill the buffer and see what happens.

e.g.
```golang
package main

import "fmt"
import "time"

func main() {
	ch := make(chan int, 2)
	go sendToChannel(ch)
	go getFromChannel(ch)
	time.Sleep(9999*time.Second)
}

func sendToChannel(ch chan int) {
	for i := 1; i<11; i++ {
		ch <- i
		fmt.Printf("send %v to channel \n", i)
	}
}


func getFromChannel(ch chan int) {
	for {
		fmt.Printf("<-=- get %v from channel \n", <-ch)
	}
}
```
