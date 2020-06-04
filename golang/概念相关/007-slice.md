> learn more : [Slices: usage and internals](https://blog.golang.org/slices-intro)

### Slices 切片
A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).


An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.    
数组是有固定大小的。而slice则是大小不固定的，数组元素的灵活视图。实际上，切片比数组使用更广泛。    
The type []T is a slice with elements of type T.  
类型`[]T`是一个元素类型为T的切片  
A slice is formed by specifying two indices, a low and high bound, separated by a colon:  
切片两个索引（高位，低位）组成的，索引间用冒号分隔  
```
a[low : high]
```
This selects a half-open range which includes the first element, but excludes the last one.  
这是个半开的范围：包括第一个元素，不包括最后一个元素  
The following expression creates a slice which includes elements 1 through 3 of a:  
以下表达式创建一个包含a数组索引1到3的元素的切片
```
a[1:4]
```
```golang
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
```

### Slices are like references to arrays 切片可看作是数组的引用
A slice does not store any data, it just describes a section of an underlying array.   
切片不储存任何数据，它只是底层数组片段的一个视图  
Changing the elements of a slice modifies the corresponding elements of its underlying array.  
切片元素的改变会修改底层数组相应的元素  
Other slices that share the same underlying array will see those changes.  
共享同一个底层数组的切片可以看到这些变化  
```golang
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
```

### Slice literals 切片字面值
A slice literal is like an array literal without the length.  
切片的字面值就像没有长度的数组字面值一样  
This is an array literal:  
这是数组字面值
```
[3]bool{true, true, false}
```
And this creates the same array as above, then builds a slice that references it:  
创建和上面一样的数组，然后创建一个切片引用它
```
[]bool{true, true, false}
```
```golang
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
```

### Slice defaults 
When slicing, you may omit the high or low bounds to use their defaults instead.  
进行截取操作时，你可以忽略高位或低位索引，用它们的默认值替代  
The default is zero for the low bound and the length of the slice for the high bound.  
低位索引默认值为0， 高位索引默认值为切片的长度  
For the array   
对于数组
```
var a [10]int
```
these slice expressions are equivalent:  
下面的切片表达式是等价的
```
a[0:10]
a[:10]
a[0:]
a[:]
```
```golang
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
```

### Slice length and capacity 切片的长度与容量
A slice has both a length and a capacity.  
切片在长度与容量的概念  
The length of a slice is the number of elements it contains.  
切片的长度是它包含的元素的数量   
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.  
切片的容量是从切片第一个元素开始统计的底层数组的元素的数量  
The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).  
切片`s`的长度与容量可以通过表达式`len(s)`, `cap(s)`获得  
You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.  
你可以通过再次截取切片扩展其长度，证明它有足够的容量。 尝试改变下面例子的切片操作，使其扩展超过自身的容量，看看会发生什么。
```golang
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```


### Nil slices  
The zero value of a slice is nil.   
slice默认值为nil   
A nil slice has a length and capacity of 0 and has no underlying array.   
nil切片长度与容量都为0， 且没有底层数组  
```golang
import "fmt"
import "unsafe"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
	fmt.Println(unsafe.Sizeof(s))
}
```

### Creating a slice with make 用make函数创建切片
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.  
切片可以通过内建函数make创建；这是创建动态大小数组的方式  
The make function allocates a zeroed array and returns a slice that refers to that array:  
make函数分配一个数组，并返回一个引用该数组的切片  
```golang
a := make([]int, 5)  // len(a)=5
```
To specify a capacity, pass a third argument to make:  
通过make函数的第三个参数，可以指定切片的容量
```golang
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```
```golang
func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```

### Slices of slices
Slices can contain any type, including other slices.  
切片能包含任何类型的数据，包括其他切片
```golang
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
```

### Appending to a slice
It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.  
往切片追加新元素是常见的，go语言为此提供了内建函数append  
```
func append(s []T, vs ...T) []T
```
The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.  
append函数的第一个参数为T类型的切片，其余的参数为往切片追加的T类型的值  
The resulting value of append is a slice containing all the elements of the original slice plus the provided values.  
append函数返回值是一个包含原切片的所有元素的和新追加元素的切片  
If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.   
如果原数组太小装不下所有元素，一个更大的数组会被分配。返回的切片将指向新分配的数组













