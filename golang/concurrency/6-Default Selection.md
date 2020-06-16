The default case in a select is run if no other case is ready.     
`select`语句存在默认case，而没有其他case可以被执行时，默认case会被执行。    
Use a default case to try a send or receive without blocking:    
在使用默认case的情况下，从通道接收值或发送值至通道是不会阻塞的   
```golang
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

e.g.
```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```
