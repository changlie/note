> 转: [Type, value and equality of interfaces](https://yourbasic.org/golang/interfaces-explained/)


Interface type  
Structural typing  
The empty interface  
Interface values  
Equality  

### Interface type
An interface type consists of a set of method signatures. A variable of interface type can hold any value that implements these methods.

In this example both Temp and *Point implement the MyStringer interface.

type MyStringer interface {
	String() string
}
type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
Actually, *Temp also implements MyStringer, since the method set of a pointer type *T is the set of all methods with receiver *T or T.

When you call a method on an interface value, the method of its underlying type is executed.

var x MyStringer

x = Temp(24)
fmt.Println(x.String()) // 24 °C

x = &Point{1, 2}
fmt.Println(x.String()) // (1,2)

### Structural typing
A type implements an interface by implementing its methods. No explicit declaration is required.

In fact, the Temp, *Temp and *Point types also implement the standard library fmt.Stringer interface. The String method in this interface is used to print values passed as an operand to functions such as fmt.Println.

var x MyStringer

x = Temp(24)
fmt.Println(x) // 24 °C

x = &Point{1, 2}
fmt.Println(x) // (1,2)

### The empty interface

The interface type that specifies no methods is known as the empty interface.

interface{}
An empty interface can hold values of any type since every type implements at least zero methods.

var x interface{}

x = 2.4
fmt.Println(x) // 2.4

x = &Point{1, 2}
fmt.Println(x) // (1,2)
The fmt.Println function is a chief example. It takes any number of arguments of any type.

func Println(a ...interface{}) (n int, err error)

### Interface values
An interface value consists of a concrete value and a dynamic type: [Value, Type]

In a call to fmt.Printf, you can use %v to print the concrete value and %T to print the dynamic type.

var x MyStringer
fmt.Printf("%v %T\n", x, x) // <nil> <nil>

x = Temp(24)
fmt.Printf("%v %T\n", x, x) // 24 °C main.Temp

x = &Point{1, 2}
fmt.Printf("%v %T\n", x, x) // (1,2) *main.Point

x = (*Point)(nil)
fmt.Printf("%v %T\n", x, x) // <nil> *main.Point
The zero value of an interface type is nil, which is represented as [nil, nil].

Calling a method on a nil interface is a run-time error. However, it’s quite common to write methods that can handle a receiver value [nil, Type], where Type isn’t nil.

You can use type assertions or type switches to access the dynamic type of an interface value. See Find the type of an object for more details.

### Equality
Two interface values are equal

if they have equal concrete values and identical dynamic types,
or if both are nil.
A value t of interface type T and a value x of non-interface type X are equal if

t’s concrete value is equal to x
and t’s dynamic type is identical to X.
var x MyStringer
fmt.Println(x == nil) // true

x = (*Point)(nil)
fmt.Println(x == nil) // false
In the second print statement, the concrete value of x equals nil, but its dynamic type is *Point, which is not nil.

Warning: See Nil is not nil for a real-world example where this definition of equality leads to puzzling results.
