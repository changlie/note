> 参考 [Go和HTTPS|Tony Bai](https://tonybai.com/2015/04/30/go-and-https/)

### 服务端私钥与证书
1. 根证书可以签发新的证书
2. 校验证书时需要加载根证书
生成根证书私钥
```
openssl genrsa -out ca.key 2048
```
用根证书私钥生成根证书
```
openssl req -x509 -new -nodes -key ca.key -subj "/CN=tonybai.com" -days 5000 -out ca.crt
```
生成服务端私钥
```
openssl genrsa -out server.key 2048
```
生成服务端公钥
```
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
```
用根证书与服务端公钥生成服务端证书
```
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000
```
客户端私钥与证书生成
```bash
$openssl genrsa -out client.key 2048
Generating RSA private key, 2048 bit long modulus
………………..+++
………………..+++
e is 65537 (0×10001)
$openssl req -new -key client.key -subj "/CN=tonybai_cn" -out client.csr
$openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 5000
Signature ok
subject=/CN=tonybai_cn
Getting CA Private Key
```


### 实例
server 相关依赖
```golang
import (
	"fmt"
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)
```
client 相关依赖
```golang
import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "net/url"
	"crypto/tls"
	"crypto/x509"
	"time"
)
```

<details>
  <summary> client忽略证书校验 </summary>

server
```golang
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintf(w, "Hello, This is an example of https service at version1", )
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("action server:9999 start up...")
	http.ListenAndServeTLS(":9999", "d:/tls/server.crt", "d:/tls/server.key", nil)
}
```
client
```golang
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:9999/createUser")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s \n", body)
```

</details>

<details>
  <summary> client校验server的证书 </summary>

server
```golang
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintf(w, "Hello, This is an example of https service at version1", )
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("action server:9999 start up...")
	http.ListenAndServeTLS(":9999", "d:/tls/server.crt", "d:/tls/server.key", nil)
}
```
client
```golang
	pool := x509.NewCertPool()

	caCrtBytes, err := ioutil.ReadFile("d:/tls/ca.crt")
	if err == nil {
		pool.AppendCertsFromPEM(caCrtBytes)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}

	client := &http.Client{Transport: tr, Timeout: 60 * time.Second}
	addr := "https://localhost:9999/save"
	resp, err := client.Get(addr)
	if err != nil {
		fmt.Println("request failed", err)
		return 
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s, error: %v \n", body, err)
```

</details>

<details>
  <summary> 证书双向校验 </summary>

server
```golang
type myHandler struct {

}

func (this *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL, r.Method)
	fmt.Fprintf(w, "it's demo of https server.")
}

func main() {
	pool := x509.NewCertPool()
	caCrtBytes, _ := ioutil.ReadFile("d:/tls/ca.crt")
	pool.AppendCertsFromPEM(caCrtBytes)

	s := &http.Server{
		Addr: ":9999",
		Handler: &myHandler{},
		TLSConfig: &tls.Config{
			ClientCAs: pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	fmt.Println("server:9999 v22")
	err := s.ListenAndServeTLS("d:/tls/server.crt", "d:/tls/server.key")
	fmt.Println(err)
}
```
client
```golang
	pool := x509.NewCertPool()
	caCrtBytes, _ := ioutil.ReadFile("d:/tls/ca.crt")
	pool.AppendCertsFromPEM(caCrtBytes)

	cliCrt, err := tls.LoadX509KeyPair("d:/tls/client.crt", "d:/tls/client.key")
	if err != nil {
		fmt.Println("LoadX509KeyPair error", err)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: pool, 
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:9999/delete")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s ; err: %v \n", body, err)
```

</details>
