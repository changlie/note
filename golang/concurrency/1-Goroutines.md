### Goroutines
A goroutine is a lightweight thread managed by the Go runtime.  
goroutine是由go运行时管理的轻量级线程。  
```
go f(x, y, z)
```
starts a new goroutine running  
以上为启动一个协程
```
f(x, y, z)
```
The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.   
f, x, y z的求值在当前协程中进行， f函数的调用会在新的协程中执行。   
Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives. (See the next slide.)   
协程在相同的地址空间运行，因此访问共享内存必须是同步的。`sync`包为同步操作提供有用的api，但你不太需要用到它们，因为go中还有其它更便利的api。
e.g. 
```golang
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello=====")
}
```
