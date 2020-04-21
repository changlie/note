> 参考

[Go http.server实战](https://studygolang.com/articles/20248)

```golang
package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var index = `
<html>
<head>
<title>my server</title>
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0">
</head>
<body style="margin:0; padding:0; width:100%;">
<form action="/" method="post" style="margin-bottom:233px;margin:0; padding:0; width:100%;">
title:<br/>
<input type="text" name="title"  style="width:100%;line-height: 60px;margin:0;" /><br/>
content:<br/>
<textarea name="content" rows="10" cols="55" ></textarea><br/>
<input type="submit" value="submit" style="width:200px;font-size:30px;line-height: 90px;height: 100px;background-color: rgb(244,244,244);"/>	
</form>
</body>
</html>
`

func main()  {
	//监听协议
	http.HandleFunc("/",HelloWorldHandler)
	http.HandleFunc("/user/login",UserLoginHandler)
	fmt.Println("server before start... at", time.Now().Format("2006/01/02 15:04:05"))
	//监听服务
	err := http.ListenAndServe(":8888",nil)

	fmt.Println("server start...")
	if err != nil {

		fmt.Println("服务器错误", err)
	}

}

func HelloWorldHandler(w http.ResponseWriter,r *http.Request)  {

	if strings.Contains(r.URL.String(), "favicon") {
		fmt.Fprintf(w,"xxx")
		return
	}

	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)

	if r.Method == "POST" {
		// 解析url传递的参数
		r.ParseForm()
		for k, v := range r.Form {
			fmt.Println("key:", k)
			// join() 方法用于把数组中的所有元素放入一个字符串。
			// 元素是通过指定的分隔符进行分隔的
			fmt.Println("val:", strings.Join(v, ""))
			fmt.Println("-----------------------------")
		}
	}

	fmt.Fprintf(w,index)
}

func UserLoginHandler(response http.ResponseWriter,request *http.Request)  {
	fmt.Println("Handler Hello")
	fmt.Fprintf(response,"Login Success")
}
```
