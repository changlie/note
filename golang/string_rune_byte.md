golang 的字符串默认是用utf-8编码保存的.
rune 是int32的别名, 用于单个字符储存, 相当于java中的char类型.


实例代码:
```golang
import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("byteTest out: ______________________")
	byteTest()
	fmt.Println("rangeTest out: ________________________")
	rangeTest()
	fmt.Println("runeTest out: _____________________________")
	runeTest()
}

var s = "5爱"

func byteTest()  {
	for i := 0; i<len(s); i++ {
		fmt.Println(s[i])
		fmt.Println(reflect.TypeOf(s[i]))
	}
}

func rangeTest()  {
	for _, v := range s {
		fmt.Println(v)
		fmt.Println(string(v))
		fmt.Println(v == '爱')
		fmt.Println(reflect.TypeOf(v))
		fmt.Println("-=> ")
	}
}

func runeTest()  {
	arr := []rune(s)
	for _, v := range arr {
		fmt.Println(v)
		fmt.Printf("%c \n", v)
		fmt.Println(reflect.TypeOf(v))
		fmt.Println("-*-")
	}
}
```

控制台输出:
```
byteTest out: ______________________
53
uint8
231
uint8
136
uint8
177
uint8
rangeTest out: ________________________
53
5
false
int32
-=>
29233
爱
true
int32
-=>
runeTest out: _____________________________
53
5
int32
-*-
29233
爱
int32
-*-
```

> 参考:

[重新认识字符串、字节、rune和字符](https://blog.csdn.net/kevin_tech/article/details/104284770)

[Golang字符串类型——byte、rune、string](https://blog.csdn.net/972301/article/details/89523243)
