The select statement lets a goroutine wait on multiple communication operations.     
`select`语句允许一个协程可以等待多个通信操作。   
A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.       
`select`语句会一直阻塞，直到其中的一个case可以被执行，然后执行相应的case。如果同时多个case可以被执行， 它会随机选择一个case执行。

e.g.
```golang
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

```
