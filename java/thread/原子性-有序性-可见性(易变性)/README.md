> 参考：https://blog.csdn.net/a60782885/article/details/77803757
>> https://zhuanlan.zhihu.com/p/36654566

>> http://ifeve.com/easy-happens-before/



对于Java并发编程，一般来说有以下的关注点：

1.线程安全性，正确性。

2.线程的活跃性(死锁，活锁)

3.性能



其中线程的安全性问题是首要解决的问题，线程不安全，运行出来的结果和预期不一致，那就连基本要求都没达到了。

保证线程的安全性问题，本质上就是保证线程同步，实际上就是线程之间的通信问题。

### 在操作系统中线程通信有以下几种方式：
1.信号量    2.信号    3.管道    4.共享内存    5.消息队列    6.socket

### java中线程通信主要使用共享内存的方式。

共享内存的通信方式首先要关注的就是**可见性**和**有序性**。而**原子性**操作一般都是必要的，所以主要关注这三个问题。
