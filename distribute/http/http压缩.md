> 参考：https://www.cnblogs.com/RainLa/p/8057367.html
>> https://zhuanlan.zhihu.com/p/90427059



http压缩通常是通过在reponse header指定Content-Encoding首部，告诉客户端response的压缩格式，这样客户端才能正确解压。

压缩的种类有哪几种？
Content-Encoding的值可以是下面这些：

```
// 常见的都是单个值
Content-Encoding: gzip   // 最常用
Content-Encoding: compress // 基本被废弃
Content-Encoding: deflate  
Content-Encoding: identity // 表示未经过压缩和修改
Content-Encoding: br  // brotli
​
// Multiple, in the order in which they were applied
Content-Encoding: gzip, identity
Content-Encoding: deflate, gzip
```

最常用的压缩格式是gzip，因为它有不错的压缩率和优秀的兼容性。br是近几年渐渐流行的一种压缩格式，全称叫"brotli"，是由Google推出的一种压缩率很高的压缩格式。但因为较新，所以浏览器兼容性并不是很好，支持的浏览器列表如下：

```
Google Chrome: Chrome 49+ 
Mozilla Firefox: Firefox 44+ 
Opera: Opera 36+
```

如今，国外国内很多大厂都在慢慢提高对brotli的使用率，它会是将来的主流。
