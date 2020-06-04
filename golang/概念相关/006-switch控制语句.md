### Switch
A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.  
switch是书写一连串if-else语句的简洁方式。它只远行条件表达式为true的case  
Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.  
go的switch语句与c,c++,java,等语言类似，但go的switch语句只运行被选择的case，而不是随后的case.  
其他语言每个case都需要带一个break语句，go语言会为每个case自动加上break，不需要手动添加  
go语言的switch语句case的值不必是常量或者整型
```golang
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

### Switch evaluation order 条件判断的顺序
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.  
switch的条件判断顺序从上到下的，当有case条件成立，停止执行。  
(For example,  
```golang
switch i {
case 0:
case f():
}
```
does not call f if i==0.)
i变量为0时，函数f不会被调用
```golang
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Thursday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("The day after tomorrow.")
	default:
		fmt.Println("Too far away.")
	}
}
```

### Switch with no condition 不带条件的switch
Switch without a condition is the same as switch true.  
不带条件的switch与switch true等同的  
This construct can be a clean way to write long if-then-else chains.  
这种构造是书写长if-then-else链的一种简洁方式。
```golang
import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

### fallthrough
使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
```golang
  i := 9
  switch {
  case i>1:
    fmt.Println("case 2>1")
    fallthrough
  case false:
    fmt.Println("case fase")
  default:
    fmt.Println("6、默认 case")
  }
```
