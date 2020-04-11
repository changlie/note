golang 一共有三种字符串拼接方式
```golang
// 方法一:
s1 := "hey,"
s1 += "tom"
fmt.Println("way1:", s1)

// 方法二:
var buf bytes.Buffer
buf.WriteString("hello ")
buf.WriteString("world")
fmt.Println("way2:", buf.String())

// 方法三:
var ss []string
ss = append(ss, "tom")
ss = append(ss, "cat")
fmt.Println("way3:", strings.Join(ss, ""))
```
三种拼接方式性能对比
```golang
  // 方法一:
	var s1 string
	startTime1 := time.Now()
	for i := 0; i<100000; i++ {
		s1 += "hello, world"
	}
	fmt.Println("way1 spend:", time.Now().Sub(startTime1).Seconds(), len(s1))

	// 方法二:
	var buf bytes.Buffer
	startTime2 := time.Now()
	for i := 0; i<100000; i++ {
		buf.WriteString("hello, world")
	}
	finalStr2 := buf.String()
	fmt.Println("way2 spend:", time.Now().Sub(startTime2).Seconds(), len(finalStr2))

	// 方法三:
	var ss []string
	startTime3 := time.Now()
	for i := 0; i<100000; i++ {
		ss = append(ss, "hello, world")
	}
	finalStr3 := strings.Join(ss, "")
	fmt.Println("way3 spend:", time.Now().Sub(startTime3).Seconds(), len(finalStr3))
```
控制台输出
```
way1 spend: 7.423729725 1200000
way2 spend: 0.001119575 1200000
way3 spend: 0.007166595 1200000
```
性能对比结论
```
buf > + > join
```

> 参考:

[字符串常用操作](https://studygolang.com/articles/11388)
