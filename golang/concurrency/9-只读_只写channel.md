定义只读，只写chan意义不太， 一般用于函数形参的描述：
```golang
import (
    "fmt"
    "time"
)

func main() {
    c := make(chan int)
    go send(c)
    go recv(c)
    time.Sleep(3 * time.Second)
}
//只能向chan里写数据
func send(c chan<- int) {
    for i := 0; i < 10; i++ {
        c <- i
    }
}
//只能取channel中的数据
func recv(c <-chan int) {
    for i := range c {
        fmt.Println(i)
    }
}
```
