Go functions may be closures.   
函数可以是闭包  
A closure is a function value that references variables from outside its body.   
闭包是一个引用了其外部变量的函数  
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.    
闭包函数可以访问及赋值给其引用的变量；在这种情况中，函数像是绑定在变量一样  
For example, the adder function returns a closure. Each closure is bound to its own sum variable.  
以下例子中， `adder`函数中返回了一个闭包函数。每个闭包函数被绑定于属于自己的`sum` 变量
```golang
import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```
