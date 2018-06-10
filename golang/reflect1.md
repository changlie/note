
## 记录一下 golang的反射

类型a 是一个 结构体
```go
// 类型信息获取
t := reflect.TypeOf(obj)      //要求obj是一个对象的值, 否则为空
fmt.Println("a's type is ",t.Name())
// 值操作
v := reflect.ValueOf(obj).Elem()    //要求obj是一个对象的引用, 否则会报错
k := v.Type()
// 遍历 obj 内的成员变量（成员须是公有的，否则尝试获取值时， 会报错）
for i := 0; i < v.NumField(); i++ {
  key := k.Field(i)
  val := v.Field(i)
  fmt.Println("line",i, "--> ", key.Name, val.Type(), val.Interface())
}
```
