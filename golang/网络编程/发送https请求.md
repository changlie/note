
无代理时
```golang
package main

import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }

    client := &http.Client{Transport: tr}

    seedUrl := "https://www.douban.com/"
    resp, err := client.Get(seedUrl)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Errorf("get https://www.douban.com/ error")
        panic(err)
    }

    fmt.Printf("%s\n", body)

}
```

有代理时
```golang
package main

import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {
    proxy := func(_ *http.Request) (*url.URL, error) {
        return url.Parse("http://你的代理:80")
    }

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        Proxy:           proxy,
    }

    client := &http.Client{Transport: tr}

    seedUrl := "https://www.douban.com/"
    resp, err := client.Get(seedUrl)
    if err != nil {
        fmt.Errorf("get https://www.douban.com/ error")
        panic(err)
    }

    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Printf("%s\n", body)
}
```





> 参考：

[Go和HTTPS](https://tonybai.com/2015/04/30/go-and-https/)

[https原理以及golang基本实现](https://www.cnblogs.com/Goden/p/4639672.html)

[golang net/http访问https](https://segmentfault.com/a/1190000016445178)
