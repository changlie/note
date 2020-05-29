The default case in a select is run if no other case is ready.

Use a default case to try a send or receive without blocking:
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
