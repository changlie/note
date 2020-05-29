> 参考： [Defer, Panic, and Recover](https://blog.go-zh.org/defer-panic-and-recover)

A ***defer statement*** pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.      
defer声明会将一个函数调用暂存至一个列表。列表中的函数调用会在函数返回后被执行。defer通常用于简化各种清理操作      
e.g.
```golang
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
```
The behavior of defer statements is straightforward and predictable. There are three simple rules:    
延迟语句的行为是直接且可预测的.它遵循三个规则：    
1. A deferred function's arguments are evaluated when the defer statement is evaluated.       
函数调用被保存至list时，函数调用的参数会被计算好
```golang
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
// "0"
```
2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.     
缓存列表中函数调用的执行顺序为： 后进先出
```golang
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}
// "3210"
```
3. Deferred functions may read and assign to the returning function's named return values     
延迟函数可以对命名返回值进行读取，赋值
```golang
func c() (i int) {
    defer func() { i++ }()
    return 1
}
// 2
```

***Panic*** is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.    
panic是一个能停止正常控制流程并触发运行时异常的内置函数。    
假设F函数调用了panic函数，F函数将停止执行并返回给调用者，但F函数内的所有deferred函数会正常执行。程序会继续执行，直到当前协程内所有函数执行返回，然后程序崩溃。

***Recover*** is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.     
recover是一个能让程序从运行时异常中恢复的内置函数。          
recover只能在deferred函数中生效。执行recover会捕获panic时产生的信息并恢复正常的程序执行。

Here's an example program that demonstrates the mechanics of panic and defer:
以下是演示panic与defer机制的示例
```golang
package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
// Calling g.
// Printing in g 0
// Printing in g 1
// Printing in g 2
// Printing in g 3
// Panicking!
// Defer in g 3
// Defer in g 2
// Defer in g 1
// Defer in g 0
// Recovered in f 4
// Returned normally from f.
```
