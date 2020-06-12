### Interfaces
An interface type is defined as a set of method signatures.  
接口被定义为方法集的签名  
A value of interface type can hold any value that implements those methods.  
接口的值可以是任何实现了接口方法的类型的值  
Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).  
注意： 下例22行有个错误， `Vertex`并没有实现`Abser`接口， 因为`Abs`方法定义在 `*Vertex` （指针）类型上
```golang
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

### Interfaces are implemented implicitly 接口被隐式实现
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.   
类型实现一个接口只需要实现了接口的方法即可。无需显示声明，无需关键字`implements`    
Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.    
隐式接口使接口定义与其实现分离，这样在任何包实现接口的时，就不需声明当前类型实现的是什么接口了。   
```golang
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```

### Interface values 接口的值
Under the hood, interface values can be thought of as a tuple of a value and a concrete type:   
接口的值可以认为是具体类型与该类型的值的元组  
```
(value, type)
```
An interface value holds a value of a specific underlying concrete type.  
接口值包含指定类型的值  
Calling a method on an interface value executes the method of the same name on its underlying type.   
调用一个接口的方法时，实际调用的是底层类型的同名方法  
```golang
import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

### Interface values with nil underlying values 接口为nil时的底层值
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.  
如果接口内部的具体值为nil, 将使用nil receiver调用方法  
In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)   
这样情况在某些语言中是会触发空指针异常的，但是，在go语言中， 编写一些被nil receiver调用的，可以被正常处理的方法是常见的   
Note that an interface value that holds a nil concrete value is itself non-nil.   
注意，内部具体值为nil的接口本身不为nil  
```golang
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

### Nil interface values
A nil interface value holds neither value nor concrete type.    
一个nil接口既不包含具体值，也不包含具体类型   
Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.    
调用一个nil接口的方法会触发运行时异常，因为接口元组没有指向具体可调用的方法的类型  
```golang
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

### The empty interface
The interface type that specifies zero methods is known as the empty interface:   
空接口是指没有包含任何方法的接口   
```golang
interface{}
```
An empty interface may hold values of any type. (Every type implements at least zero methods.)   
空接口可以持有任何类型的值。（因为每个类型都至少实现了0个方法）   
Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.   
空接口可以用来处理未知类型值。例如，方法`fmt.Print`可以接收任意数量的interface{}类型的值
```golang
import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

### Type assertions  类型推断
A type assertion provides access to an interface value's underlying concrete value.  
类型推断使我们可以访问接口包含的具体值   
```golang
t := i.(T)
```
This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.     
上面表达式推断接口值i包含类型为T, 并把底层的T类型的值赋值给变量t  
If i does not hold a T, the statement will trigger a panic.  
如果i包含的不是T类型，表达式将会触发一个运行时异常。   

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.   
类型推断一个接口值是否包含指定类型，并返回两个值：具体值，推断的bool结果值
```golang
t, ok := i.(T)
```
If i holds a T, then t will be the underlying value and ok will be true.   
如果i持有类型T, t为i持有的具体值，ok为true  
If not, ok will be false and t will be the zero value of type T, and no panic occurs.  
如果i不持有类型T, t为类型T的默认值， ok为false, 且不会触发运行时异常  
Note the similarity between this syntax and that of reading from a map.  
注意这种类型推断与读取map的相似性
```golang
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
```

### Type switches
A type switch is a construct that permits several type assertions in series.  
类型switch允许我们对接口值同时进行多个类型推断   
A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.   
类型switch与普通的switch类似，但类型switch的case指定的类型而不是值，这些类型将与接口内的具体的类型比较。
```golang
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```
The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.   
类型switch声明与类型推断`i.(T)`语法相同, 但指定的类型T，被关键字`type`所替代  

This switch statement tests whether the interface value i holds a value of type T or S. In each of the T and S cases, the variable v will be of type T or S respectively and hold the value held by i. In the default case (where there is no match), the variable v is of the same interface type and value as i.  
switch语句测试接口值i持有类型T或S值。在caseT，S中，v将分别是T类型的值，S类型的值。如果i持有的既不是T类型的值，也不是S类型的值，程序将执行默认case分支，默认case中， 变量v等于变量i.

e.g.
```golang
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```


























