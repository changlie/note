Go supports first class functions, higher-order functions, user-defined function types, function literals, closures, and multiple return values.

This rich feature set supports a functional programming style in a strongly typed language.

### first class functions
头等函数（first-class function）是指在程序设计语言中，函数被当作头等公民。这意味着，函数可以作为别的函数的参数、函数的返回值，赋值给变量或存储在数据结构中。 有人主张应包括支持匿名函数（函数字面量，function literals）。在这样的语言中，函数的名字没有特殊含义，它们被当作具有函数类型的普通的变量对待。1960年代中期，克里斯托弗·斯特雷奇在“functions as first-class citizens”中提出这一概念

### User-defined function types
> In Go, functions can be passed around just like any other value. A function's type signature describes the types of its arguments and return values.

### Higher-order functions
> A function can use other functions as arguments and return values.

### Function literals and closures
> Anonymous functions can be declared in Go, as in this example. Function literals are closures: they inherit the scope of the function in which they are declared.

A function literal represents an anonymous function.
```golang
FunctionLit = "func" Signature FunctionBody .
```
```golang
func(a, b int, z float64) bool { return a*b < int(z) }
```
A function literal can be assigned to a variable or invoked directly.
```golang
f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)
```
**Function literals are closures**: they may refer to variables defined in a surrounding function. Those variables are then shared between the surrounding function and the function literal, and **they survive as long as they are accessible.**

### Multiple return values
> Go functions can return multiple values.
