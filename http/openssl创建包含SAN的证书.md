创建自定义配置文件
```
[req]
distinguished_name = req_distinguished_name
req_extensions = v3_req

[req_distinguished_name]
countryName = country
stateOrProvinceName = province
localityName = city
organizationName = company name
commonName =  domain name or ip
 
[v3_req]
subjectAltName = @alt_names

[alt_names]
DNS.1=chenlh.cn
DNS.2=yun.com
DNS.3=localhost
DNS.4=changlie.com
```
创建根证书
```bash
// 创建根私钥
$ openssl genrsa -out changlie.key 2048
Generating RSA private key, 2048 bit long modulus
............................+++
............................+++
e is 65537 (0x10001)

// 创建根证书
$ openssl req -new -key changlie.key -out changlie.csr -config ssl.cnf                                                  You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
country []:cn
province []:gd
city []:sz
company name []:dream.fly
domain name or ip []:chenlh.cn

$ openssl x509 -req -days 365 -in changlie.csr -signkey changlie.key -out changlie.crt
Signature ok
subject=/C=cn/ST=gd/L=sz/O=dream.fly/CN=chenlh.cn
Getting Private key
```

基于根证书创建子证书
```bash
// 生成子证书私密钥
$ openssl genrsa -out local.key 2048
Generating RSA private key, 2048 bit long modulus
......................................+++
..........+++
e is 65537 (0x10001)

// 基于根证书创建子证书
$ openssl req -new -key local.key -out localhost.csr -config ssl.cnf
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
country []:cn
province []:gd
city []:sz
company name []:dt
domain name or ip []:localhost

$ openssl x509 -req -days 500 -CA changlie.crt -CAkey changlie.key -in localhost.csr -out local.crt -CAcreateserial -ex
tensions v3_req -extfile  ssl.cnf
Signature ok
subject=/C=cn/ST=gd/L=sz/O=dt/CN=localhost
Getting CA Private Key
```




[使用openssl生成证书-cnblogs](https://www.cnblogs.com/hugetong/p/11579749.html)

[OpenSSL生成带有SubjectAltName的自签名证书-csdn](https://blog.csdn.net/u010983881/article/details/83619603)

[]()

[]()

[]()

[]()
