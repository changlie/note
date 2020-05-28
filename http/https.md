### 简介
SSL/TLS协议的基本过程是这样的：
```
（1） 客户端向服务器端索要并验证公钥。

（2） 双方协商生成"对话密钥"。

（3） 双方采用"对话密钥"进行加密通信。
```
上面过程的前两步，又称为"握手阶段"（handshake）。

握手阶段分成五步:
```
第一步，爱丽丝给出协议版本号、一个客户端生成的随机数（Client random），以及客户端支持的加密方法。

第二步，鲍勃确认双方使用的加密方法，并给出数字证书、以及一个服务器生成的随机数（Server random）。

第三步，爱丽丝确认数字证书有效，然后生成一个新的随机数（Premaster secret），并使用数字证书中的公钥，加密这个随机数，发给鲍勃。

第四步，鲍勃使用自己的私钥，获取爱丽丝发来的随机数（即Premaster secret）。

第五步，爱丽丝和鲍勃根据约定的加密方法，使用前面的三个随机数，生成"对话密钥"（session key），用来加密接下来的整个对话过程。
````

### 密钥与证书的生成
- 根证书可以签发新的证书
- 校验证书时需要加载根证书
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


> 参考：

[图解SSL/TLS协议-阮一峰](http://www.ruanyifeng.com/blog/2014/09/illustration-ssl.html)

[SSL/TLS协议远行机制的概述-阮一峰](http://www.ruanyifeng.com/blog/2014/02/ssl_tls.html)

[HTTPS中CA证书的签发及使用过程](https://www.cnblogs.com/xdyixia/p/11610102.html)

[CA 证书](https://blog.whezh.com/certification/)
