
io.Copy 从标准输入读信息写到一个文件中 (*** interesting ***)

```golang
import (
  "io"
  "os"
  "syscall"
)
func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: mout  filepath....")
    return
  }
  
	filepath := os.Args[1]
	var out io.Writer
  var in io.Reader

  in = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
  out, _ = os.OpenFile(filepath, os.O_RDWR | os.O_CREATE, 0666)

  n, err := io.Copy(out, in)

  fmt.Printf("\n write: %d Byte,  err: %v \n", n, err)
}
```
