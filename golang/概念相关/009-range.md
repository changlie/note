> 参考 [范围 range](https://www.runoob.com/go/go-range.html)

range 关键字用于 for 循环中迭代字符串(string)、数组(array)、切片(slice)、集合(map)或通道(channel)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。  
string
```golang
s := "hello, 世界"
for i := range s {
  fmt.Print(i, ", ")
}
fmt.Println("")
for i, c := range s {
  fmt.Printf("%v:'%v', ", i, string(c))
}
fmt.Println("")
```

array/slice
```golang
arr := []int{98, 56, 1, 23, 55}
for i := range arr {
  fmt.Print(i, ", ")
}
fmt.Println("")
for i, num := range arr {
  fmt.Printf("%v:%v, ", i, num)
}
fmt.Println("")
```

map
```golang
m := map[string]interface{}{
  "name":"tome",
  "age":33,
  "man":false,
}
for key := range m {
  fmt.Print(key, ", ")
}
fmt.Println("")
for key, val := range m {
  fmt.Printf("%v:%v, ", key, val)
}
fmt.Println("")
```

channel
```golang
import (
  "fmt"
  "math/rand"
)

func main() {
  ch := make(chan int, 10)
	go addNumToChannel(ch)
	for i := range ch {
		fmt.Printf("%v, ", i)
	}
}

func addNumToChannel(ch chan<- int) {
	len := 10 
	for i:=0; i<len; i++ {
		ch <- rand.Intn(len)
	}
	fmt.Println("add finish!")
  close(ch)
}
```

### Range
The range form of the for loop iterates over a slice or map.  
for循环的range格式可以迭代切片或map  
When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.  
对切片进行range操作，每个迭代会返回两个值。第一个是索引值，第二个是索引对应的元素值
```golang
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```
You can skip the index or value by assigning to _.  
你可以忽略索引值或者元素值通过赋值给下划线 `_`
```golang
for i, _ := range pow
for _, value := range pow
```
If you only want the index, you can omit the second variable.  
如果你只想要索引值，你可以忽略第二个值
```golang
for i := range pow
```
```golang
func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```




