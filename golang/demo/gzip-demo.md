```golang
package main

import (
    "bytes"
    "compress/gzip"
    "encoding/base64"
    "fmt"
    "io/ioutil"
)

func main()  {
    //src := []byte("碳水化合物工作")
    src, _ := ioutil.ReadFile("d:/qk-files")
    str := encode(src)
    fmt.Println(str)

    err := ioutil.WriteFile("d:/qk-gzip", []byte(str), 0666)
    fmt.Println(err)

    res := decode(str)
    fmt.Println(string(res))
}

func decode(str string) []byte {
    data, _ := base64.StdEncoding.DecodeString(str)
    // 解压
    rdata := bytes.NewReader(data)
    r, _ := gzip.NewReader(rdata)
    s, _ := ioutil.ReadAll(r)
    return s
}

func encode(src []byte) string {
    var b bytes.Buffer
    // 压缩
    gz := gzip.NewWriter(&b)
    if _, err := gz.Write(src); err != nil {
        panic(err)
    }
    if err := gz.Flush(); err != nil {
        panic(err)
    }
    if err := gz.Close(); err != nil {
        panic(err)
    }
    str := base64.StdEncoding.EncodeToString(b.Bytes())
    return str
}

```
