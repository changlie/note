新建文件req.cnf，内容如下：
```
[req]
distinguished_name = req_distinguished_name
x509_extensions = v3_req
prompt = no
[req_distinguished_name]
C = US
ST = VA
L = SomeCity
O = MyCompany
OU = MyDivision
CN = test.com
[v3_req]
keyUsage = critical, digitalSignature, keyAgreement
extendedKeyUsage = serverAuth
subjectAltName = @alt_names
[alt_names]
DNS.1 = test.com
```
> 注： CN和DNS.1都需要设定为域名(测试域名为test.com)

生成私钥和自签名证书：
```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout test.com.key -out test.com.crt -config req.cnf -sha256
```

> 转： [自签名证书](https://github.com/3gstudent/3gstudent.github.io/blob/master/_posts/2018-4-18-CIA%20Hive%20Beacon%20Infrastructure%E5%A4%8D%E7%8E%B02%E2%80%94%E2%80%94%E4%BD%BF%E7%94%A8Apache%20mod_rewrite%E5%AE%9E%E7%8E%B0https%E6%B5%81%E9%87%8F%E5%88%86%E5%8F%91.md)


> 参考： https://www.cnblogs.com/benwu/articles/4891758.html
