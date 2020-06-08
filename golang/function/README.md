Go supports first class functions, higher-order functions, user-defined function types, function literals, closures, and multiple return values.  
Go支持头等函数，高阶函数，用户定义的函数类型，函数字面值，闭包，多个返回值  
This rich feature set supports a functional programming style in a strongly typed language.  
这些功能为强类型的函数式编程风格提供了支持  

### first class functions
头等函数，函数作为一等公民
> 头等函数（first-class function）是指在程序设计语言中，函数被当作头等公民。这意味着，函数可以作为别的函数的参数、函数的返回值，赋值给变量或存储在数据结构中。 有人主张应包括支持匿名函数（函数字面量，function literals）。在这样的语言中，函数的名字没有特殊含义，它们被当作具有函数类型的普通的变量对待。1960年代中期，克里斯托弗·斯特雷奇在“functions as first-class citizens”中提出这一概念

### User-defined function types
用户定义的函数类型
> In Go, functions can be passed around just like any other value. A function's type signature describes the types of its arguments and return values.  
函数可以像其它值一样被传递。 函数的类型签名描述了其参数和返回值的类型。

### Higher-order functions
高阶函数
> A function can use other functions as arguments and return values.  
函数可以作为其它函数的参数和返回值

### Function literals and closures
函数字面值和闭包（定义闭包时通常要用到函数字面值）
> Anonymous functions can be declared in Go, as in this example. Function literals are closures: they inherit the scope of the function in which they are declared.    
go支持匿名函数。函数字面值是闭包，其继承了定义它的函数的作用域

A function literal represents an anonymous function.  
函数字面值就是匿名函数
```golang
FunctionLit = "func" Signature FunctionBody .
```
```golang
func(a, b int, z float64) bool { return a*b < int(z) }
```
A function literal can be assigned to a variable or invoked directly.  
函数字面值可以被赋值给一个变量，或被直接调用
```golang
f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)
```
**Function literals are closures**: they may refer to variables defined in a surrounding function. Those variables are then shared between the surrounding function and the function literal, and **they survive as long as they are accessible.**  
函数就是闭包： 其可能引用函数内定义的变量。这些变量被定义它们的函数及函数字面值共享，只要它们能被访问，就会一直存活。

### Multiple return values
多个返回值
> Go functions can return multiple values.  
go函数能返回多个值
