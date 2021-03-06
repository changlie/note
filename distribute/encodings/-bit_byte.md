> 参考： https://zhuanlan.zhihu.com/p/46040087

## 一.存储单位的bit 和 Byte
### 1.bit（比特）
bit也就是我们不一定听说过的比特，大名鼎鼎的比特币就是以此命名的。它的简写为小写字母 “b” 。

作为信息技术的最基本存储单元，因为比特实在太小了，所以大家生活中并不是经常听到。那么 bit 是什么呢？

电脑是以二进制存储以及发送接收数据的。二进制的一位，就叫做 1 bit。也就是说 bit 的含义就是二进制数中的一个数位，即 “0” 或者 "1"。

### 2.Byte（字节）
Byte 是字节的英文写法。它的简写为大写字母 “B"。

既然名字叫字节，那肯定跟字符有关系。是的。英文字符通常是一个字节，也就是 1B，中文字符通常是两个字节，也就是 2B。

字节 Byte 和比特 bit 的换算关系是 1 Byte = 8 bit 。

### 3. KB （千字节）
需要了解的是，1 KB 并不是一千字节，因为计算机只认识二进制，所以在这里的 KB，是 2 的 10 次方，也就是 1024 个字节。

另外很多表示存储单位的地方都把 B 写成 b，造成了大家认知的混乱。其实在存储单位计量中出现 b 的地方，它的意思仍然是 B，不要因为 bit 的缩写是 b 就被误导了，在存储计量中是不会用 比特，千比特 这种单位的。但是在网速计量中，b的真实意思就是指 比特 了，这个我们下面再说。

### 4.存储单位换算
存储单位换算关系如下：
```
1Byte = 1 bit
1KB = 1024B
1MB = 1024KB
1GB = 1024MB
1TB = 1024GB
```

## 二.网速怎么算？
网络线路的计量单位，也就是我们通常说的 2M 宽带，10 M 宽带的单位，是 比特每秒（bits per second）。比特每秒 的缩写为 bps，意思是每秒接收的平均比特数。更大的单位是 千比特每秒（Kbps）或 兆比特每秒（Mbps）。2M宽带，意味着每秒平均可以接受 2Mb 的数据，也就是二百万左右比特的数据，在这里，小写 b 的意思就是比特了。

而通常我们说的下载速度，也就是网速，是每秒下载的字节数。比如网速是 5 KB（这网速可是够慢的），意思就是每秒接收的数据是 五千字节。

那我们根据 一字节 等于 8 比特的 换算方法，就可以得出以下结论。

下载速度从理论上来说，应该是 带宽的 八分之一。

2M 宽带理论下载速度是 256 KB

10M 宽带理论下载速度是 1280 KB

实际上由于我们还需要接受一些下载需要的一些信息，如 IP 信息，HTTP 信息，再加上服务器传输速度，电脑配置等原因，网速会比理论慢一些。

由于很多人都会混淆 字节的大写 B 和比特的小写 b，造成各种混乱。所以在书写单位缩写时，一定要注意 字节 和 比特 的单位的大小写 。

总结一下，有以下几个要点需要注意：
```
存储单位和网速的单位，不管是 B 还是 b，代表的都是 字节 Byte。
带宽的单位，不管是 B 还是 b，代表的都是 比特 bit 。
```
只要记住这两点，就算别人写混淆了，我们也是可以明白单位的具体含义的。
