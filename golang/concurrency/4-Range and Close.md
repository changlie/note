A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression:        
发送者可以关闭一个通道以指示数据已全部发送完毕。接收者可以通过接收表达式赋值第二个参数以检测一个通道是否已关闭。
```golang
v, ok := <-ch
```
ok is false if there are no more values to receive and the channel is closed.   
ok值为false, 表示所有值已接收完毕，通道已关闭。
The loop `for i := range c` receives values from the channel repeatedly until it is closed.  
循环表达式`for i := range c`会一直从通道接收值，直到通道被关闭后，退出循环。
Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
注意：通道应该被发送者关闭，而不是应该是接收者。向已关闭的通道发生数据会触发运行时异常。   
Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.   
另一个需要注意的是：通道与文件是不同的，通常情况下，你是不需要关闭它们的。 只有接收者必须被告知数据已全部接收完毕时，关闭操作才是必要的。例如终止一个对通道的range循环    
e.g.
```golang
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	fmt.Println("push finish!!!")
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i, len(c), cap(c))
	}
}
```
