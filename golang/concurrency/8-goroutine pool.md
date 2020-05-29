```golang
package main

import (
  "fmt"
  "time"
)

var taskQueue chan int
func main() {
	taskQueue = make(chan int)
	go func() {
		for i := 0; i<100; i++ {
			time.Sleep(1 * time.Second)
			taskQueue <- i*9
		}
	}()

	go func() {
		for i := 1; i<6; i++ {
			go func(id int) {
				for {
					fmt.Printf("goroutine id: %v, val: %v \n", id, <-taskQueue)
				}
			}(i)
		}
	}()

	time.Sleep(9999*time.Second)
}
```
