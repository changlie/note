### 字符串占位符
```golang
res := os.Expand("hi! ${name}, age:${age}, come from $addr.  $unknow.", mapping)
fmt.Println(res)
func mapping(key string) string {
  fmt.Println("call by other func")
  m := map[string]string{
    "name":"jerry",
    "age":"33",
    "addr":"china",
  }
  if key != "" {
    val, ok := m[key]
    if ok {
      return val
    }
    return key
  }

  return key
}
```
输出
```
hi! jerry, age:33, come from china.  unknow.
```

### 环境变量相关（所有操作都不会持久化）
```golang
func Clearenv()   // 清除环境变量（只限这次启动的会话，相关于缓存式清空）
func Environ() []string // 获取所有环境变量
func LookupEnv(key string) (string, bool)  // 获取单个环境变量
func Getenv(key string) string // 获取单个环境变量
func Setenv(key, value string) error // 设置环境变量
func Unsetenv(key string) error // 删除单个环境变量
```



>  参考：[os包](https://cloud.tencent.com/developer/article/1342799)
