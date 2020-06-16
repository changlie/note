Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:    
通道是可以带缓冲区的。在`make`函数的第二参数提供缓冲区的长度用于初始化一个带缓冲区的通道。    
```golang
ch := make(chan int, 100)
```
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.         
通道的缓冲区已满时，对该通道进行发送数据的操作会发生阻塞。通道的缓冲区为空时，对该通道进行接收数据也会发生阻塞。   
Modify the example to overfill the buffer and see what happens.      
修改示例以溢出缓冲区，看看会发生什么。   
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
