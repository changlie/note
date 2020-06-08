<details>
  <summary> URL字段列表 </summary>
  
```golang
type URL struct {
    Scheme     string
    Opaque     string    // encoded opaque data
    User       *Userinfo // username and password information
    Host       string    // host or host:port
    Path       string    // path (relative paths may omit leading slash)
    RawPath    string    // encoded path hint (see EscapedPath method); added in Go 1.5
    ForceQuery bool      // append a query ('?') even if RawQuery is empty; added in Go 1.7
    RawQuery   string    // encoded query values, without '?'
    Fragment   string    // fragment for references, without '#'
}
```

</details>

<details>
  <summary> 设置获取query参数 </summary>

```golang
import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	fmt.Println("query arg q:", q["q"]) // query arg q: [dotnet]
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u) // https://google.com/search?q=golang
}
```

</details>


<details>
  <summary> 获取url path的两个方式 </summary>
  
```golang
import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Parse + String preserve the original encoding.
	u, err := url.Parse("https://example.com/foo%2fbar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Path) // /foo/bar
	fmt.Println(u.RawPath) // /foo%2fbar
	fmt.Println(u.String()) // https://example.com/foo%2fbar
}
```

</details>

