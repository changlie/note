func ReadFull(r Reader, buf []byte) (n int, err error)
```golang
import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read=\n")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(longBuf))

}
```

func WriteString(w Writer, s string) (n int, err error)
```golang
import (
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout, "Hello World")

}
```


io.Copy 从标准输入读信息写到一个文件中 ***(interesting)***

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
