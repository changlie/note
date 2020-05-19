### 日期处理

```golang
// 时间戳
fmt.Println("时间戳: ", time.Now().Unix())


// 日期格式化, 记忆: 6,1,2  3,4,5
fmt.Println("日期格式化: ", time.Now().Format("2006/01/02 15:04:05"))


// 时间戳转Time对象(go中的日期对象)
t := time.Unix(1586597850, 0)
fmt.Println(t.Unix(), ": ", t.Format("2006/01/02 15:04:05"))


// 格式化string转Time对象: 1586597850
// 方法一:
t = time.Date(2020, time.April, 11, 17, 37, 30, 0, time.Local)
fmt.Println(t.Unix())
// 方法二:
t, _ = time.Parse("2006/01/02 15:04:05", "2020/04/11 17:37:30")
fmt.Println("string 转Time对象: ", t.Unix())

```

> 参考:

[golang -- 时间日期总结](https://studygolang.com/articles/669)
