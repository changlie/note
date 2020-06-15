We've seen how channels are great for communication among goroutines.    
我们已见识了在协程间通信中使用通道是一个多么棒的事件。     
But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?    
但如果我们不需要进行协程间通信呢？如果我们仅仅是希望同一时间只有一个协程能访问共享变量以避免冲突又怎么做呢？       
This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.   
这个概念叫互斥，为这个概念提供数据结构的俗称叫互斥体。        
Go's standard library provides mutual exclusion with sync.Mutex and its two methods:    
Go标准库通过`sync.Mutex`及它的两个方法提供互斥的功能     
```
Lock
Unlock
```
We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock as shown on the Inc method.     
如下例方法`Inc`所示： 我们可以定义一个互斥的代码块通过被`Lock`, `Unlock`这两个函数调用包裹起来。       
We can also use defer to ensure the mutex will be unlocked as in the Value method.     
如下例方法`Value`所示： 我们也可以使用defer声明确保`mutex`一定被解锁。       
e.g.
```golang
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

```
