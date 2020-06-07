打开浏览器
```golang
import (
    "os/exec"
)

func main() {
    cmd := exec.Command("cmd", "/C", "start", "http://github.com/changlie")
    // linux
    // cmd := exec.Command("xdg-open", "http://as:8080/hello")
    cmd.Run()
}
```
