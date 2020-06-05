打开浏览器
```golang
import (
    "os/exec"
)

func main() {
    cmd := exec.Command("cmd", "/C", "start", "http://github.com/changlie")
    cmd.Run()
}
```
