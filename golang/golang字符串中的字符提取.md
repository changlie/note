golang 的字符串默认是用utf-8编码保存的. 所以直接用内置函数len()是得不到真正的字符数量的.


### 字符数量获取
```golang
fmt.Println("char count:", len("1994你好")) // 错误方式
fmt.Println("char count:", utf8.RuneCountInString("1994你好")) // 正确方式
```

### 字符串遍历
```golang
const nihongo = "石碣語007"
for i, w := 0, 0; i < len(nihongo); i += w {
	runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
	fmt.Printf("%v: %c \n", i, runeValue)
	w = width
}
fmt.Println("==========================")

s := "123木头人"
for i, v := range s {
	fmt.Println(i, ": ", string(v))
}
```

### 字符串与byte 互转
```golang
func stringToBytes() {
	s := "天堂"
	data := []byte(s)
	fmt.Println("data: ", data)
	fmt.Println("=====================")
}

func bytesToString() {
	var data [10]byte
	data[1] = '1'
	data[2] = 'T'
	fmt.Println(string(data[:]))
	fmt.Println("====================")
}
```
