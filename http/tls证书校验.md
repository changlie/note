### 一、证书及相关概念简介
SSL／TLS：一种加密协议，提供了计算机网络通信的安全规范和方案。
OpenSSL：它是SSL／TLS的一个实现，而且包括了一批非常棒的工具软件。
X.509：是由ITU-T为了公开密钥基础建设（PKI）与授权管理基础建设（PMI）提出的产业标准，规范了公开密钥认证、证书吊销列表、授权证书、证书路径验证算法、证书内容及格式等（https的证书就是遵循这个标准）。X.509的证书基于ASN.1来描述，按照rfc5280规范，X509 V3版本的证书基本语法如下（只列举了Certificate和TBSCertificate，其它更详细的请参考rfc5280）。其中tbsCertificate的数据段被拿来做信息摘要，并且用上级证书的私钥加密后形成签名置入https证书中。
```
Certificate  ::=  SEQUENCE  {
        tbsCertificate       TBSCertificate,
        signatureAlgorithm   AlgorithmIdentifier,
        signatureValue       BIT STRING  }
TBSCertificate  ::=  SEQUENCE  {
        version         [0]  EXPLICIT Version DEFAULT v1,
        serialNumber         CertificateSerialNumber,
        signature            AlgorithmIdentifier,
        issuer               Name,
        validity             Validity,
        subject              Name,
        subjectPublicKeyInfo SubjectPublicKeyInfo,
        issuerUniqueID  [1]  IMPLICIT UniqueIdentifier OPTIONAL,
                             -- If present, version MUST be v2 or v3
        subjectUniqueID [2]  IMPLICIT UniqueIdentifier OPTIONAL,
                             -- If present, version MUST be v2 or v3
        extensions      [3]  EXPLICIT Extensions OPTIONAL
                             -- If present, version MUST be v3
        }
```
- ASN.1（Abstract Syntax Notation One），一种描述数字对象的方法和标准。ASN.1提供了多种数据编码方法。包括了BER、DER、PER和XER等。这些编码方法规定了将数字对象转换成应用程序能够处理、保存和网络传输的二进制编码形式的一组规则。
- DER编码（Distinguished Encoding Rules）：属于ASN.1下的BER（Basic Encode Rules）编码派生出来的编码规则，这种编码规则下，一个相同的ASN.1对象编码后能得到唯一的编码数据（BER编码不能保证这一点，即一个对象编码后可能产生多个不同的编码数据）
- PEM编码（Privacy Enhanced Mail）：是一种保密邮件的编码标准，在rfc1421规范中规定。X.509的证书在DER编码的基础上进行base64编码，然后添加一些头、尾标志就是PEM格式编码了，头尾的标志也是PEM的一部分，不要随意改动。比如baidu.com.cer经过下面命令转成PEM后：
```bash
//转换命令
openssl x509 -inform der -in <.cer file> -out <.pem output file>
```
//PEM格式的文件如下，头和尾告诉我们这是一个certificate：
```
-----BEGIN CERTIFICATE-----
MIIG0DCCBbigAwIBAgIMGNoar9s9QTCfF9MLMA0GCSqGSIb3DQEBCwUAMGYxCzAJ这里省略一堆类似的符号Qht29ZyGrAAdgJfW/9iFD+Kg5jsj0KyGxxzB3i4QHOjkftpQZrCYAvLWQuZu4vTX8daTPwJTcfR/R6MAweYUJUcuS98=
-----END CERTIFICATE-----
```
另外，感兴趣的话，你可以直接将baidu.com.cer直接base64编码，可以看到编码出来的内容除了没有PEM的头尾标志，其它内容是一致的。附带几个base64命令：
```bash
openssl base64 -d -in <infile> -out <outfile>
openssl base64 -in <infile> -out <outfile>
openssl enc -base64 <<< <sting to encode>
openssl enc -base64 -d <<< <string to decode>
```

把DER转化成给人看的信息
使用以下命令将der格式的证书信息提取出来
```bash
//-inform der告诉openssl，要转换的文件格式为der，如果是pem则无需指明
//-noout，不要同时生成该证书的pem内容，如果没有这个选项，会同时在输出文件的末尾附带上证书的pem，你可以去掉试一下
openssl x509 -in bd.cer -inform der -text -noout >> bd.cer.noout.txt
```
执行后，打开bd.cer.noout.txt，可以看到内容如下（建议先看RSA的加密算法RSA算法原理）
```
Public Key Algorithm: rsaEncryption，表明证书公钥是RSA公钥
Modulus (2048 bit):指出RSA的模(即RSA算法一开始选的两个质数的乘积)
Exponent: 65537 (0x10001)，指出RSA的指数
Signature Algorithm: sha256WithRSAEncryption，指出证书的签名，以及其签名算法，先用sha256做信息摘要，再对摘要内容做RSA加密。
```






[证书校验](https://www.zybuluo.com/blueGhost/note/807076)
