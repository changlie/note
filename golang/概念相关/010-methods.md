### Methods
Go does not have classes. However, you can define methods on types.  
go没有类的概念，但你可以给类型定义方法  
A method is a function with a special receiver argument.  
方法是一个带有特殊receiver参数的函数  
The receiver appears in its own argument list between the func keyword and the method name.  
receiver位于func关键字与方法名称之间      
In this example, the Abs method has a receiver of type Vertex named v.    
下例中，方法`Abs`有个名为`v`，类型为`Vertex`的Receiver  
```golang
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

### Methods are functions
Remember: a method is just a function with a receiver argument.  
记住： 方法只是带receiver参数的函数  
Here's Abs written as a regular function with no change in functionality.   
下例`Abs`被编写为与上例功能一样的普通函数  
```golang
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```

You can declare a method on non-struct types, too.   
你也可以为非结构体类型定义方法   
In this example we see a numeric type MyFloat with an Abs method.   
下例为：一个带有`Abs`方法的数值类型`MyFloat`   
You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).      
你只可以为本包内定义的类型定义方法。为非本包内的定义的类型（包括像`int`这样的内置类型）定义方法是不被允许的   
```golang
import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```

### Pointer receivers
You can declare methods with pointer receivers.  
你可以定义带指针receiver的方法  
This means the receiver type has the literal syntax `*T` for some type `T`. (Also, T cannot itself be a pointer such as `*int`.)  
也就是说，对于类型`T`, receiver类型有语法`*T` (当然，`T`是不可以是 像`int`这样的内置类型)   

For example, the Scale method here is defined on *Vertex.  
下列中，方法`Scale`是定义在`*Vertex`类型的方法  

Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.  
带指针receiver的方法可以修改指针receiver指向的值（如下例Scale）. 由于方法经常修改他们的receiver，所以指针receiver比值receiver更常用。   

Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.   
移除第16行的函数`Scale`的`*`, 观察程序结果的变化  

With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.   
如果方法`Scale`带的是值receiver的话，它操作的是从原先的Vertex值复制出来的值。它想修改Vertex值的话，它必须带指针receiver 
```golang
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

```

### Pointers and functions
Here we see the Abs and Scale methods rewritten as functions.  
我们把上例的`Abs`, `Scale`方法改写为函数   
Again, try removing the `*` from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?  
再次移除`Scale`函数的`*`。 你能看见程序发生了什么变化吗？ 知道还需要修改什么地方才可以让程序编译通过  
(If you're not sure, continue to the next page.)  
如果还不太明白为什么，接着往下看  
```
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
```

### Methods and pointer indirection
Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:   
比较上两个程序，你或许注意到了：函数的指针形参只接收一个指针参数   
```golang
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```
while methods with pointer receivers take either a value or a pointer as the receiver when they are called:   
而带指针receiver的方法在被调用时， 可以接收一个值或者指针作为它的receiver
```golang
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```
For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.     
对于表达式`v.Scale(5)`, 尽管`v`不是指针是值，带指针receiver的方法`Scale`还是会执行。这是出于方便的目的，go允许`(&v).Scale(5)`简写成`v.Scale(5)`
```golang
import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
```

The equivalent thing happens in the reverse direction.    
同样的，带值receiver的方法被调用时接收一个指针作为receiver, 也是可以被简写的    
Functions that take a value argument must take a value of that specific type:    
函数的值形参接收的参数必须是一个值  
```golang
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```
while methods with value receivers take either a value or a pointer as the receiver when they are called:   
而带值receiver的方法被调用时，可以接收一个值或指针作为它的receiver:  
```golang
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```
In this case, the method call p.Abs() is interpreted as (*p).Abs().   
`(*p).Abs()`被简写成了`p.Abs()`
```golang
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
```

### Choosing a value or pointer receiver  
There are two reasons to use a pointer receiver.   
以下是选择使用指针receiver的两个理由：    
The first is so that the method can modify the value that its receiver points to.  
1. 为了让方法可以修改receiver指向的值   
The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.     
2. 避免每次方法调用都复制一次receiver指向的值。而且当receiver是一个大结构体时，这种方式是更高效的   
In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.    
下例中，方法`Scale`, `Abs`都是带`*Vertex`类型receiver的，尽管方法`Abs`并没有修改它的receiver   
In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)     
对于同一类型的方法， 要么都带值receiver, 要么都带指针receiver， 不要混合使用
```golang
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```




























