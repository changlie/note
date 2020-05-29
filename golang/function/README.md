Go supports first class functions, higher-order functions, user-defined function types, function literals, closures, and multiple return values.

This rich feature set supports a functional programming style in a strongly typed language.

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
