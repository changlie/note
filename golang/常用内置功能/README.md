打开浏览器
```golang
import "os/exec"

func main() {
    cmd := exec.Command("cmd", "/C", "start", "http://github.com/changlie")
    // linux
    // cmd := exec.Command("xdg-open", "http://as:8080/hello")
    cmd.Run()
}
```

判断某类型是否实现了某接口
```golang
var u User
// 方式一
ty := reflect.TypeOf(u)
if ty.Implements(reflect.TypeOf((*Json)(nil)).Elem()) {
    fmt.Println("type User implement interface Json")
}

// 方式二
var v interface{} = u
if val, ok := v.(Json); ok {
   fmt.Println("type User implement interface Json")
}

```
