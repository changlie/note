Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.   
通道是一个类型化管道，你可以通过管道使用通道运算符`<-`发送值以及接收值。  
```golang
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```
(The data flows in the direction of the arrow.)     
数据流向箭头指向的方向。   
Like maps and slices, channels must be created before use:   
和map与切片一样，通道在使用前必须先被创建。   
```golang
ch := make(chan int)
```
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.    
默认情况下，另一端准备好之前，数据的发送与接收操作会一直阻塞。 这特性允许协程在没有使用显式锁与条件变量的情况下进行同步。   
The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.      
下例代码对切片里面的数字进行求和，并这项工作分配到两个协程中。一旦两个协程完成了自己的工作，程序便会计算出最终的结果。   
e.g.
```golang
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

> 参考：

[Golang 漫谈之channel妙法](https://juejin.im/entry/5da165c9f265da5b8f107dbf)

[channel的使用场景](https://segmentfault.com/a/1190000017958702)
