```golang

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	var b bytes.Buffer
  // 压缩
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte("YourDataHere")); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	fmt.Println(str)
	data, _ := base64.StdEncoding.DecodeString(str)
	fmt.Println(data)
  // 解压
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	fmt.Println(string(s))
}
```
