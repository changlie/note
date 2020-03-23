### java 网络编程

#### 相关概念
- TCP/IP是一个协议集
- socket是面向程序员的一套编程接口
- 两者关系：socket对传输层中的TCP/UDP协议进行了封装，对用户隐藏了内部TCP/UDP的具体实现，
它与TCP/IP的关系就像windows系统与win32编程接口一样。通过win32编程接口我们可以调用windows系统提供的功能，而通过socket 我们可以实现网络编程。


> 参考： https://segmentfault.com/a/1190000014476924
>> https://blog.csdn.net/qq_30270931/article/details/80483124

>> https://blog.csdn.net/honghailiang888/article/details/51121257/


引用：
```
TCP/IP只是一个协议栈，就像操作系统的运行机制一样，必须要具体实现，同时还要提供对外的操作接口。
这个就像操作系统会提供标准的编程接口，比如win32编程接口一样，
TCP/IP也要提供可供程序员做网络开发所用的接口，这就是Socket编程接口
```
