<details>
  <summary> 文件服务器 </summary>

```golang
import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}
```

</details>

<details>
  <summary> 为文件服务器添加一个前缀 </summary>

```golang
func main() {
  // 访问相应文件要在url加指定前缀才可以，否则报404
  // 以下例子：d:/目录下有个demo.html, url为/demo.html会报404， url为/tmpfiles/demo.html才可以访问成功
  http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("d:/"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

</details>
